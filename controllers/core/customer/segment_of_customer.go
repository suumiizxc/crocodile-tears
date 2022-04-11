package customer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateSOCInput struct {
	Name    string `json:"name"`
	Name2   string `json:"name2"`
	OrderNo uint   `json:"orderNo"`
}

func SOCList(c *gin.Context) {
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit"})
		return
	}
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in page"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[[], %v, %v]`, page*limit, limit))
	response := helper_core.CH.Request(helper_core.SEGMENT_OF_CUSTOMER_LIST, bodyData)
	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}
	c.JSON(response.StatusCode, gin.H{"data": response.Data, "message": "successfully"})
}

func SOCCreate(c *gin.Context) {
	var input CreateSOCInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorDTO": errDTO.Error()})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[{}, {"name":"%v", "name2": "%v", "orderNo": %v}]`, input.Name, input.Name2, input.OrderNo))
	response := helper_core.CH.Request(helper_core.SEGMENT_OF_CUSTOMER_INSERT, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorREPONSE": response.Err.Error()})
		return
	}
	c.JSON(response.StatusCode, gin.H{"data": nil, "message": "successfully"})
}

func SOCDelete(c *gin.Context) {
	socID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[{}, %v]`, socID))
	response := helper_core.CH.Request(helper_core.SEGMENT_OF_CUSTOMER_DELETE, bodyData)

	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}

	c.JSON(response.StatusCode, gin.H{"data": nil, "message": "successfully"})
}

func SOCGet(c *gin.Context) {
	socID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[%v]`, socID))
	response := helper_core.CH.Request(helper_core.SEGMENT_OF_CUSTOMER_SELECT, bodyData)
	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.Err.Error()})
		return
	}
	var jsondata map[string]interface{}
	json.Unmarshal([]byte(response.DataString), &jsondata)
	c.JSON(response.StatusCode, gin.H{"data": jsondata})
}
