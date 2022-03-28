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

type Result struct {
	Status int
	Result string
	Reason string
}

func CreateCustomer(c *gin.Context) {

	poster := []Post{{Userid: "1", Title: "foo", Body: "bar"}}

	var tmp []interface{}
	var data = []byte(`[ 404, "error", "Not Found" ]`)
	if err := json.Unmarshal(data, &tmp); err != nil {
		log.Fatal(err)
	}
	// Not ugly! Not fragile!
	// fmt.Println("Status code : ", int(tmp[0].(float64)))

	fmt.Println("Status code : ", tmp)

	json_data, err := json.Marshal(poster)
	if err != nil {
		fmt.Println(err.Error())
	}
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
