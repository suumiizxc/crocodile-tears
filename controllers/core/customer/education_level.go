package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

func GetEducationLevels(c *gin.Context) {
	var bodyData = []byte(`[]`)
	response := helper_core.CH.Request(helper_core.LIST_EDUCATION_LEVEL, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data})
}
