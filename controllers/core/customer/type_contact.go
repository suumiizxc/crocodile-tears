package customer

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

func TypeOfContactList(c *gin.Context) {
	bodyData := []byte(`[]`)
	response := helper_core.CH.Request(helper_core.TYPE_OF_CONTACT_LIST, bodyData)
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
