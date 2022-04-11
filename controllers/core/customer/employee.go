package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateEmployeeInput struct {
	Name    string `json:"name"`
	Name2   string `json:"name2"`
	OrderNo uint   `json:"orderNo"`
}

func EmployeeList(c *gin.Context) {
	bodyData := []byte(`[]`)
	response := helper_core.CH.Request(helper_core.EMPLOYEE_LIST, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data, "message": "successfully"})
}

func EmployeeCreate(c *gin.Context) {
	var input CreateEmployeeInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in dto " + errDTO.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[{name : "%v", name2 :"%v", orderNo : %v}]`, input.Name, input.Name2, input.OrderNo))
	response := helper_core.CH.Request(helper_core.EMPLOYEE_CREATE, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data, "message": "successfully"})
}

func EmployeeGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EMPLOYEE_GET, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data, "message": "successfully"})
}

func EmployeeDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EMPLOYEE_DELETE, bodyData)
	if response.Err != nil {
		if response.StatusCode == 200 {
			c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
			return
		}
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Failed in request " + response.DataString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data, "message": "successfully"})
}
