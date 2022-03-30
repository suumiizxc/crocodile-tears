package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// models "github.com/suumiizxc/gin-bookstore/models/core/customer"
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
	fmt.Println("bodyData : ", bodyData)
	req, err := http.NewRequest("POST", "http://202.131.242.158:4020/nes.s.Web/NesFront", bytes.NewBuffer(bodyData))

	if err != nil {
		log.Printf("Request failed : %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "NESSESSION=22QivFUT3jGC681187SzyfmEzJn7DL")
	req.Header.Add("op", "10201170")
	req.Header.Add("company", "11")
	req.Header.Add("lang", "1")
	req.Header.Add("role", "53")

	client := &http.Client{Timeout: time.Second * 10}

	fmt.Println("req : ", req)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error reading response : ", err)
	}
	defer resp.Body.Close()
	fmt.Println("status code :", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Request failed : %s", err)
	}
	fmt.Println("body :", body)
	bodyString := string(body)
	log.Print(bodyString)

	var tmp []interface{}
	var data = []byte(bodyString)
	if err := json.Unmarshal(data, &tmp); err != nil {
		log.Fatal(err)
	}
}
