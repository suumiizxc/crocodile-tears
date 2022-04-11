package customer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	// models "github.com/suumiizxc/gin-bookstore/models/core/customer"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

/*[{"name":"add_type","name2":"addtype_","parentTypeId2":"91","addrCodeLen":21,"orderNo":123}]*/

type CreateAddressTypeInput struct {
	Name          string `json:"name"`
	Name2         string `json:"name2"`
	ParentTypeId  string `json:"parentTypeId"`
	ParentTypeId2 string `json:"parentTypeId2"`
	AddrCodeLen   uint   `json:"addrCodeLen"`
	OrderNo       uint   `json:"orderNo"`
}

func AddressTypeList(c *gin.Context) {
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
	response := helper_core.CH.Request(helper_core.ADDRESS_TYPE_LIST, bodyData)
	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}
	c.JSON(response.StatusCode, gin.H{"data": response.Data, "message": "successfully"})
}

func AddressTypeCreate(c *gin.Context) {
	var input CreateAddressTypeInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorDTO": errDTO.Error()})
		return
	}

	var bodyData = []byte(fmt.Sprintf(`[{"name":"%v","name2":"%v","ParentTypeId": "%v","parentTypeId2":"%v","addrCodeLen":%v,"orderNo":%v}]`, input.Name, input.Name2, input.ParentTypeId, input.ParentTypeId2, input.AddrCodeLen, input.OrderNo))
	response := helper_core.CH.Request(helper_core.ADDRESS_TYPE_INSERT, bodyData)
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
	c.JSON(http.StatusBadRequest, gin.H{"errorREPONSE": response.Err.Error()})
}

// func EditNation(c *gin.Context) {
// 	limit, err := strconv.ParseUint(c.Param("limit"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in limit"})
// 		return
// 	}
// 	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in page"})
// 		return
// 	}
// 	var bodyData = []byte(fmt.Sprintf(`[[], %v, %v]`, page*limit, limit))
// 	response := helper_core.CH.Request(helper_core.LIST_NATION, bodyData)
// 	if response.Err != nil {
// 		log.Printf("Request failed : %s", response.Err.Error())
// 	}
// 	c.JSON(response.StatusCode, gin.H{"data": response.Data})
// }

func AddressTypeDelete(c *gin.Context) {
	nationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[{}, %v]`, nationID))
	response := helper_core.CH.Request(helper_core.ADDRESS_TYPE_DELETE, bodyData)

	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}

	c.JSON(response.StatusCode, gin.H{"data": nil, "message": "successfully"})
}

func AddressTypeGet(c *gin.Context) {
	nationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[%v]`, nationID))
	response := helper_core.CH.Request(helper_core.ADDRESS_TYPE_SELECT, bodyData)

	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.Err.Error()})
		return
	}

	var jsondata map[string]interface{}
	json.Unmarshal([]byte(response.DataString), &jsondata)

	c.JSON(response.StatusCode, gin.H{"data": jsondata})
}
