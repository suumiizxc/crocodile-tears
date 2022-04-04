package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
	"github.com/suumiizxc/gin-bookstore/helper/redis"
)

func GetEducationDegrees(c *gin.Context) {
	val, err := redis.RS.Get("POLARIS_COOKIE_TOKEN").Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in redis"})
		return
	}
	var bodyData = []byte(`[]`)
	response := helper_core.CH.Request(helper_core.LIST_EDUCATION_DEGREE, val, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data})
}
