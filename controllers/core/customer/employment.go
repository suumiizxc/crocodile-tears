package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateEmploymentInput struct {
	Name        string `json:"name"`
	Name2       string `json:"name2"`
	Description string `json:"description"`
	OrderNo     uint   `json:"orderNo"`
}

func EmploymentList(c *gin.Context) {
	bodyData := []byte(`[]`)
	response := helper_core.CH.Request(helper_core.EMPLOYMENT_LIST, bodyData)
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

func EmploymentCreate(c *gin.Context) {
	var input CreateEmploymentInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in request body " + errDTO.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[{"name" : "%v", "name2" : "%v", "description":"%v", "orderNo" : %v}]`, input.Name, input.Name2, input.Description, input.OrderNo))
	response := helper_core.CH.Request(helper_core.EMPLOYMENT_CREATE, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			var jsondata map[string]interface{}
			json.Unmarshal([]byte(response.DataString), &jsondata)
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
}

func EmploymentDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EMPLOYMENT_DELETE, bodyData)
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

func EmploymentGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EMPLOYMENT_GET, bodyData)
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
