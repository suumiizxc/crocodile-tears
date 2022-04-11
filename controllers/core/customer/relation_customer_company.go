package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateRelationCustomerCompanyInput struct {
	Name    string `json:"name"`
	Name2   string `json:"name2"`
	OrderNo uint   `json:"orderNo"`
}

func RelationCustomerCompanyList(c *gin.Context) {
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in page " + err.Error()})
		return
	}
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[[], %v, %v]`, (page-1)*limit, limit))
	response := helper_core.CH.Request(helper_core.RELATION_CUSTOMER_AND_COMPANY_LIST, bodyData)
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

func RelationCustomerCompanyCreate(c *gin.Context) {
	var input CreateRelationCustomerCompanyInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in errDTO : " + errDTO.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[{name : "%v", name2 : "%v", orderNo : %v}]`, input.Name, input.Name2, input.OrderNo))
	response := helper_core.CH.Request(helper_core.RELATION_CUSTOMER_AND_COMPANY_CREATE, bodyData)

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

func RelationCustomerCompanyGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.RELATION_CUSTOMER_AND_COMPANY_GET, bodyData)
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

func RelationCustomerCompanyDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in id " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[%v]`, id))
	response := helper_core.CH.Request(helper_core.RELATION_CUSTOMER_AND_COMPANY_DELETE, bodyData)
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
