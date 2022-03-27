package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Userid string `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func CreateCustomer(c *gin.Context) {
	// var input CreateCustomerInput
	// customer := models.Customer{}

	poster := []Post{{Userid: "1", Title: "foo", Body: "bar"}}
	json_data, err := json.Marshal(poster)
	if err != nil {
		fmt.Println(err.Error())
	}

	// resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts",
	// 	params)
	fmt.Println("request body : ", bytes.NewBuffer(json_data))
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(json_data))

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	post := Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Printf("Post added with ID %d", post.ID)
}
