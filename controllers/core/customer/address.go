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

/* [
{
	"addrCode":"",
	"addrCode2":"",
	"typeId":"87",
	"name":"New MONGOL",
	"name2":"New MONGOL",
	"nameShort":"NM",
	"nameShort2":"NM",
	"zipCode":121,
	"orderNo":0
 }
]
*/

type CreateAddressInput struct {
	AddrCode   string `json:"addrCode"`
	AddrCode2  string `json:"addrCode2"`
	TypeId     string `json:"typeId"`
	Name       string `json:"name"`
	Name2      string `json:"name2"`
	NameShort  string `json:"nameShort"`
	NameShort2 string `json:"nameShort2"`
	ZipCode    string `json:"zipCode"`
	OrderNo    uint   `json:"orderNo"`
}

func AddressList(c *gin.Context) {
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
	response := helper_core.CH.Request(helper_core.ADDRESS_LIST, bodyData)
	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}
	c.JSON(response.StatusCode, gin.H{"data": response.Data, "message": "successfully"})
}

func AddressCreate(c *gin.Context) {
	var input CreateAddressInput
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorDTO": errDTO.Error()})
		return
	}

	var bodyData = []byte(fmt.Sprintf(`[{"addrCode":"%v", "addrCode2":"%v", "typeId":"%v", "name":"%v", "name2":"%v", "nameShort":"%v", "nameShort2":"%v", "zipCode":%v, "orderNo":%v}]`, input.AddrCode, input.AddrCode2, input.TypeId, input.Name, input.Name2, input.NameShort, input.NameShort2, input.ZipCode, input.OrderNo))
	response := helper_core.CH.Request(helper_core.ADDRESS_INSERT, bodyData)
	if response.Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorREPONSE": response.Err.Error()})
		return
	}
	c.JSON(response.StatusCode, gin.H{"data": nil, "message": "successfully"})
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

func AddressDelete(c *gin.Context) {
	nationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[{}, %v]`, nationID))
	response := helper_core.CH.Request(helper_core.ADDRESS_DELETE, bodyData)

	if response.Err != nil {
		log.Printf("Request failed : %s", response.Err.Error())
	}

	c.JSON(response.StatusCode, gin.H{"data": nil, "message": "successfully"})
}

func AddressGet(c *gin.Context) {
	nationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed in parse params"})
		return
	}
	var bodyData = []byte(fmt.Sprintf(`[%v]`, nationID))
	response := helper_core.CH.Request(helper_core.ADDRESS_SELECT, bodyData)

	if response.StatusCode != 200 {
		c.JSON(http.StatusNotImplemented, gin.H{"error": response.Err.Error()})
		return
	}

	var jsondata map[string]interface{}
	json.Unmarshal([]byte(response.DataString), &jsondata)

	c.JSON(response.StatusCode, gin.H{"data": jsondata})
}
