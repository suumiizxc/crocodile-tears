package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/config"
	models "github.com/suumiizxc/car-marketplace/models/client"
)

type CreatePermissionInput struct {
	ClientID uint   `json:"client_id" binding:"required"`
	Key      string `json:"key" binding:"required"`
	Value    string `json:"value" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

func FindPermissionByCID(c *gin.Context) {
	var permissions []models.Permission
	if err := config.DB.Where("client_id = ?", c.Param("client_id")).Find(&permissions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": permissions})
}

func CreatePermission(c *gin.Context) {
	var input CreatePermissionInput
	permission := models.Permission{}
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if errDTO := smapping.FillStruct(&permission, smapping.MapFields(&input)); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDTO.Error()})
		return
	}
	if err := config.DB.Save(&permission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": permission})
}

func ValidatePermission(clientID uint, key string) (models.Permission, error) {
	var permission models.Permission
	if err := config.DB.Where("client_id = ?", clientID).Where("key = ?", key).First(&permission).Error; err != nil {
		return models.Permission{}, fmt.Errorf("error permission : %v", err.Error())
	}
	if permission.Status != "active" {
		return permission, fmt.Errorf("error permission : %v", "permission not active")
	}
	return permission, nil
}

func ValidatePermissionAction(client_id uint, module_name string, sub_module_name string, action string) (models.Permission, error) {
	key := module_name + "-" + sub_module_name + "-" + action
	var permission models.Permission
	if err := config.DB.Where("client_id = ?", client_id).Where("key = ?", key).First(&permission).Error; err != nil {
		return models.Permission{}, fmt.Errorf("error permission : %v", err.Error())
	}
	if permission.Status != "active" {
		return permission, fmt.Errorf("error permission : %v", "permission not active")
	}
	return permission, nil
}
