package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/config"
	"github.com/suumiizxc/car-marketplace/helper/redis"
	models "github.com/suumiizxc/car-marketplace/models/client"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	module_name     = "client"
	sub_module_name = "client"
)

type CreateClientInput struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Registration string `json:"registration" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
}

type UpdateClientInput struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	Registration string `json:"registration"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

type LoginClientPhoneInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginClientEmailInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FindClients(c *gin.Context) {
	var clients []models.Client
	config.DB.Find(&clients)
	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func FindClient(c *gin.Context) {
	var client models.Client
	if err := config.DB.Where("id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func CreateClient(c *gin.Context) {
	var input CreateClientInput
	client := models.Client{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if errDTO := smapping.FillStruct(&client, smapping.MapFields(&input)); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDTO.Error()})
		return
	}
	client.Password = _hashAndSalt([]byte(client.Password))
	if err := config.DB.Transaction(func(tx *gorm.DB) error {

		prev := models.Client{}
		if prev_err := tx.Where("email = ?", client.Email).First(&prev).Error; prev_err == nil {
			return fmt.Errorf("%v", "Already user created this email")
		}
		if err := tx.Create(&client).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println("pisda", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.Token = uuid.New().String()
	// json, err := json.Marshal(client)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// 	return
	// }
	// if err := redis.RS.Set(client.Token, json, 0).Err(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error-redis": err})
	// 	return
	// }
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(9000) + 1000
	var activition models.ClientActivation
	activition.ClientID = client.ID
	activition.ExpireDate = time.Now().Add(time.Minute * 5)
	activition.OTP = fmt.Sprintf("%v", randNumber)
	resp, _ := http.Get(fmt.Sprintf("https://internal.mik.mn/kXJyEbb7O3co?number=%v&msg=%v&source=INTERNAL", client.Phone, randNumber))

	_, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	activition.Status = "pending"
	if err := config.DB.Create(&activition).Error; err != nil {
		fmt.Println("activition token error : ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}

func LoginPhone(c *gin.Context) {
	var input LoginClientPhoneInput
	client := models.Client{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Where("phone = ?", input.Phone).First(&client)
	if !_comparePassword(client.Password, []byte(input.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed credential"})
		return
	}
	client.Token = uuid.New().String()
	jsonClient, err := json.Marshal(client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := redis.RS.Set(client.Token, jsonClient, 0).Err(); err != nil {
		c.JSON(http.StatusBadRequest, "Failed in redis")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func ActivatePhone(c *gin.Context) {
	client := models.Client{}
	otp := models.ClientActivation{}
	otp.OTP = c.Param("otp")
	clientID, err := strconv.ParseUint(c.Param("client_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "client_id not defined"})
		return
	}
	client.ID = uint(clientID)
	if err := config.DB.Where("otp = ?", otp.OTP).Where("client_id = ?", client.ID).Last(&otp).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Failed in otp", "message": err.Error()})
		return
	}
	if err := config.DB.Where("id = ?", client.ID).Find(&client).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Failed in client", "message": err.Error()})
		return
	}
	if client.ID > 0 {
		if client.ID == otp.ClientID {
			if otp.ExpireDate.Sub(time.Now()) >= 0 {
				client.IsActive = 1
				client.Role = 1
				fmt.Println("id : ", client.ID, " otp : ", otp.ID)
				config.DB.Model(&models.Client{}).Where("id = ?", client.ID).Update("role", 1).Update("is_active", 1)
				config.DB.Model(&models.ClientActivation{}).Where("id = ?", otp.ID).Update("status", "success")
				c.JSON(http.StatusOK, gin.H{"message": "successfully activated"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Activition token expired"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in activition token"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in activition token"})
		return
	}

}

func LoginEmail(c *gin.Context) {
	var input LoginClientEmailInput
	client := models.Client{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Where("email = ?", input.Email).First(&client)
	if !_comparePassword(client.Password, []byte(input.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed credential"})
		return
	}
	client.Token = uuid.New().String()
	jsonClient, err := json.Marshal(client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := redis.RS.Set(client.Token, jsonClient, 0).Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in redis"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client, "status": http.StatusText(http.StatusOK)})
}

func ProfileClient(c *gin.Context) {
	// clientToken, _, err := _validateClient(c.GetHeader("access_token"), module_name, sub_module_name, "profile")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	tokenid, _ := strconv.ParseUint(c.GetHeader("access_token"), 10, 64)
	var client models.Client
	if err := config.DB.Where("id = ?", tokenid).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func _validateClient(access_token string, module_name string, sub_module_name string, action string) (models.Client, models.Permission, error) {
	val, err := redis.RS.Get(access_token).Result()
	if err != nil {
		return models.Client{}, models.Permission{}, err
	}
	var clientToken models.Client
	json.Unmarshal([]byte(val), &clientToken)
	permission, err := ValidatePermissionAction(clientToken.ID, module_name, sub_module_name, action)
	if err != nil {
		return models.Client{}, models.Permission{}, err
	}
	fmt.Println("permission : ", permission)
	return clientToken, permission, nil
}

func _hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("failed to hash a password")
	}
	return string(hash)
}

func _comparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
