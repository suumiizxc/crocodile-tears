package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

type CreateLanguageInput struct {
	LangCode string `json:"langCode"`
	Iso2     string `json:"iso2"`
	Name     string `json:"name"`
	Name2    string `json:"name2"`
	OrderNo  uint   `json:"orderNo"`
}

func LanguageList(c *gin.Context) {
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in page param : " + err.Error()})
		return
	}
	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit param : " + err.Error()})
		return
	}
	bodyData := []byte(fmt.Sprintf(`[[], %v, %v]`, (page-1)*limit, limit))
	response := helper_core.CH.Request(helper_core.LANGUAGE_LIST, bodyData)
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

func LanguageGet(c *gin.Context) {
	param := c.Param("param")
	bodyData := []byte(fmt.Sprintf(`["%v"]`, param))
	response := helper_core.CH.Request(helper_core.LANGUAGE_GET, bodyData)
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

func LanguageDelete(c *gin.Context) {
	param := c.Param("param")
	bodyData := []byte(fmt.Sprintf(`["%v"]`, param))
	response := helper_core.CH.Request(helper_core.LANGUAGE_DELETE, bodyData)
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

func LanguageCreate(c *gin.Context) {
	var input CreateLanguageInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in request " + errDTO.Error()})
		return
	}
	bodyData := []byte(
		fmt.Sprintf(
			`[{langCode : "%v", iso2 : "%v", name : "%v", name2 : "%v", orderNo : %v}]`,
			input.LangCode, input.Iso2, input.Name, input.Name2, input.OrderNo,
		))
	response := helper_core.CH.Request(helper_core.LANGUAGE_CREATE, bodyData)
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
