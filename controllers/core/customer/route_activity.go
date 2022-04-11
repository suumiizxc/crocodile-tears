package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateRouteOfActivityInput struct {
	Name             string `json:"name"`
	Name2            string `json:"name2"`
	ExtCode          string `json:"extCode"`
	OrderNo          uint   `json:"orderNo"`
	ParentIndustryID uint   `json:"parentIndustryId"`
}

func RouteOfActivityList(c *gin.Context) {

	value, err := strconv.ParseUint(c.Param("value"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in value : " + err.Error()})
		return
	}
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit : " + err.Error()})
		return
	}
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit : " + err.Error()})
		return
	}
	var bodyData []byte
	fmt.Println("input  : ", value)
	if value == 0 {
		bodyData = []byte(fmt.Sprintf(
			`[[], %v, %v]`,
			(page-1)*limit, limit,
		))

	} else {

		bodyData = []byte(fmt.Sprintf(
			`[[{_iField: "PARENT_INDUSTRY_ID", _iOperation: "=", _iType: 1, _iValue: "%v"}], %v, %v]`,
			value, (page-1)*limit, limit,
		))
	}

	response := helper_core.CH.Request(helper_core.ROUTE_OF_ACTIVITY_LIST, bodyData)
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

func RouteOfActivityGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.ROUTE_OF_ACTIVITY_GET, bodyData)
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
	c.JSON(http.StatusOK, gin.H{"data": response.Data})
}

func RouteOfActivityCreate(c *gin.Context) {
	var input CreateRouteOfActivityInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in dto : " + errDTO.Error()})
		return
	}
	var bodyData []byte
	if input.ParentIndustryID == 0 {
		bodyData = []byte(fmt.Sprintf(`[{}, {name : "%v", name2 :"%v", orderNo : %v, extCode:"%v"}]`, input.Name, input.Name2, input.OrderNo, input.ExtCode))
	} else {
		bodyData = []byte(fmt.Sprintf(`[{}, {name : "%v", name2 :"%v", orderNo : %v, extCode:"%v", parentIndustryId : %v}]`, input.Name, input.Name2, input.OrderNo, input.ExtCode, input.ParentIndustryID))
	}
	response := helper_core.CH.Request(helper_core.ROUTE_OF_ACTIVITY_CREATE, bodyData)
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

func RouteOfActivityDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[{},%v]`, id))
	response := helper_core.CH.Request(helper_core.ROUTE_OF_ACTIVITY_DELETE, bodyData)
	fmt.Println("response : ", response)
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
