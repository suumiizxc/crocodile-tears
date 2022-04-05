package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type EducationDegreeInput struct {
	Name    string `json:"name"`
	Name2   string `json:"name2"`
	OrderNo string `json:"orderNo"`
}

func EducationDegreeList(c *gin.Context) {
	var bodyData = []byte(`[]`)
	response := helper_core.CH.Request(helper_core.EDUCATION_DEGREE_LIST, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.Data})
}

func EducationDegreeCreate(c *gin.Context) {
	var input EducationDegreeInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDTO.Error()})
		return
	}
	bodyString := fmt.Sprintf(`[{"name" : "%v", "name2": "%v", "orderNo": %v}]`, input.Name, input.Name2, input.OrderNo)
	var bodyData = []byte(bodyString)
	response := helper_core.CH.Request(helper_core.EDUCATION_DEGREE_CREATE, bodyData)
	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, response.Err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response.DataString, "message": "successfully"})
}

func EducationDegreeDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf("[%v]", id))
	response := helper_core.CH.Request(helper_core.EDUCATION_DEGREE_DELETE, bodyData)
	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully", "data": nil})
}

func EducationDegreeGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.EDUCATION_DEGREE_GET, bodyData)
	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.Err.Error()})
		return
	}
	var jsondata map[string]interface{}
	json.Unmarshal([]byte(response.DataString), &jsondata)

	c.JSON(http.StatusOK, gin.H{"data": jsondata})
}
