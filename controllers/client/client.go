package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/gin-bookstore/config"
	models "github.com/suumiizxc/gin-bookstore/models/client"
	"gorm.io/gorm"
)

type CreateClientInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UpdateClientInput struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if errDTO := smapping.FillStruct(&client, smapping.MapFields(&input)); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDTO})
		return
	}
	if err := config.DB.Transaction(func(tx *gorm.DB) error {

		prev := models.Client{}
		if prev_err := tx.Where("email = ?", client.Email).First(&prev).Error; prev_err == nil {
			return fmt.Errorf("error : %v", "Already user created this email")
		}
		if err := tx.Create(&client).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}
