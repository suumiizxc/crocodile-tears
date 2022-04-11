package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateTypeOfOrganizationInput struct {
	Name     string `json:"name"`
	Name2    string `json:"name2"`
	CustType string `json:"custType"`
	OrderNo  uint   `json:"orderNo"`
	IsMain   uint   `json:"isMain"`
}

func TypeOfOrganizationList(c *gin.Context) {
	bodyData := []byte(`[]`)
	response := helper_core.CH.Request(helper_core.TYPE_OF_ORGANIZATION_LIST, bodyData)
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

func TypeOfOrganizationCreate(c *gin.Context) {
	var input CreateTypeOfOrganizationInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in dto " + errDTO.Error()})
		return
	}
	fmt.Println("input : ", input)
	bodyData := []byte(fmt.Sprintf(`[{"name" : "%v", "name2" : "%v", "custType":"%v", "orderNo" : %v, "isMain":%v}]`, input.Name, input.Name2, input.CustType, input.OrderNo, input.IsMain))
	response := helper_core.CH.Request(helper_core.TYPE_OF_ORGANIZATION_CREATE, bodyData)
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

func TypeOfOrganizationDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.TYPE_OF_ORGANIZATION_DELETE, bodyData)
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

func TypeOfOrganizationGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.TYPE_OF_ORGANIZATION_GET, bodyData)
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
