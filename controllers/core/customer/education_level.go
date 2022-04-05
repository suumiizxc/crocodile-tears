package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateEducationLevelInput struct {
	Name    string `json:"name"`
	Name2   string `json:"name2"`
	OrderNo uint   `json:"orderNo"`
}

func EducationLevelList(c *gin.Context) {
	var bodyData = []byte(`[]`)
	response := helper_core.CH.Request(helper_core.EDUCATION_LEVEL_LIST, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data})
}

func EducationLevelCreate(c *gin.Context) {
	var input CreateEducationLevelInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in request" + errDTO.Error()})
	}
	var bodyData = []byte(fmt.Sprintf(`[{"name" : "%v", "name2" : "%v", "orderNo" : %v}]`, input.Name, input.Name2, input.OrderNo))
	response := helper_core.CH.Request(helper_core.EDUCATION_LEVEL_CREATE, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
}

func EducationLevelGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EDUCATION_LEVEL_GET, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			var jsondata map[string]interface{}
			json.Unmarshal([]byte(response.DataString), &jsondata)
			c.JSON(http.StatusOK, gin.H{"data": jsondata, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data, "message": "successfully"})
}

func EducationLevelDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EDUCATION_LEVEL_DELETE, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": nil, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nil, "message": "successfully"})
}
