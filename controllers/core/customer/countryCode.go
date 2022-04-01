package customer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// models "github.com/suumiizxc/gin-bookstore/models/core/customer"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
)

func GetCountryCodes(c *gin.Context) {
	// countryCodes := []models.CountryCode{}

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
	response, err := helper_core.CH.Request("10201170", "VV0BQPbHHeJ6IlxM8MQTwBWoYDqnrc", bodyData)
	if err != nil {
		log.Printf("Request failed : %s", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": response})

	// 	req.Header.Add("Content-Type", "application/json")
	// 	req.Header.Add("Cookie", "NESSESSION=hCQzwGBMtY9YblZwtMnKfyO1UG49bN")
	// 	req.Header.Add("op", "10201170")
	// 	req.Header.Add("company", helper_core.PC)
	// 	req.Header.Add("lang", "1")
	// 	req.Header.Add("role", "53")

	// 	client := &http.Client{Timeout: time.Second * 3}

	// 	fmt.Println("req : ", req)
	// 	resp, err := client.Do(req)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	// 	fmt.Println("status code :", resp)
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Printf("Request failed : %s", err)
	// 	}
	// 	fmt.Println("body :", body)
	// 	bodyString := string(body)
	// 	log.Print(bodyString)

	// 	var tmp []interface{}
	// 	var data = []byte(bodyString)
	// 	if err := json.Unmarshal(data, &tmp); err != nil {
	// 		log.Fatal(err)
	// 	}
}
