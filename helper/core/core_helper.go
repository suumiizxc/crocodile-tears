package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/suumiizxc/gin-bookstore/helper/redis"
)

type core_helper struct {
	polaris_company string
	polaris_lang    string
	polaris_role    string
	polaris_url     string
}

func (corehelper *core_helper) Init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Failed env : %v", err))
	}
	corehelper.polaris_company = os.Getenv("POLARIS_COMPANY")
	corehelper.polaris_lang = os.Getenv("POLARIS_LANG")
	corehelper.polaris_role = os.Getenv("POLARIS_ROLE")
	corehelper.polaris_url = os.Getenv("POLARIS_URL")
}

// func (corehelper core_helper) setNewName() {
// 	if err := godotenv.Load(".env"); err != nil {
// 		panic(fmt.Sprintf("Failed env : %v", err))
// 	}
// 	corehelper.polaris_company = os.Getenv("POLARIS_COMPANY")
// 	corehelper.polaris_lang = os.Getenv("POLARIS_LANG")
// 	corehelper.polaris_role = os.Getenv("POLARIS_ROLE")
// 	corehelper.polaris_url = os.Getenv("POLARIS_URL")
// }

type Response struct {
	Data       []interface{}
	DataString string
	StatusCode int
	Err        error
}

func (_ch core_helper) Request(opcode string, field []byte) *Response {
	val, err := redis.RS.Get("POLARIS_COOKIE_TOKEN").Result()
	if err != nil {

		return &Response{Data: nil, DataString: "", StatusCode: 401, Err: fmt.Errorf("error : %v", err.Error())}
	}
	req, err := http.NewRequest("POST", _ch.polaris_url, bytes.NewBuffer(field))

	if err != nil {
		log.Printf("Request failed : %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "NESSESSION="+val)
	req.Header.Add("op", opcode)
	req.Header.Add("company", _ch.polaris_company)
	req.Header.Add("lang", _ch.polaris_lang)
	req.Header.Add("role", _ch.polaris_role)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)

	if err != nil {
		return &Response{Data: nil, DataString: "", StatusCode: resp.StatusCode, Err: fmt.Errorf("error : %v", err.Error())}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Request failed : %s", err)
	}
	bodyString := string(body)

	var tmp []interface{}
	var data = []byte(bodyString)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return &Response{Data: nil, DataString: bodyString, StatusCode: resp.StatusCode, Err: fmt.Errorf("error : %v", err.Error())}
	}
	return &Response{Data: tmp, DataString: bodyString, StatusCode: 200, Err: nil}
}

var CH = new(core_helper)
