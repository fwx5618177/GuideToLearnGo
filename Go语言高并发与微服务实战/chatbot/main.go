package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type requestBody struct {
	Key string `json:"key"`
	Info string `json:"info"`
	UserId string `json:"userid"`
}

type responseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string `json:"list"`
	Url string `json:"url"`
}

func process(inputChan <-chan string, userid string) {
	for {
		input := <- inputChan

		fmt.Println("input:", input)

		if input == "EOF" {
			break
		}
	
		reqData := &requestBody {
			Key: "1123",
			Info: input,
			UserId: userid,
		}
	
		byteData, _ := json.Marshal(&reqData)
	
		req, err := http.NewRequest("POST", 
			"http://www.tuling123.com/openapi/api",
			bytes.NewReader(byteData),
		)
	
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	
		client := http.Client{}
		resp, err := client.Do(req)
	
		if err != nil {
			fmt.Println("Network Error!", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
	
			var respData responseBody
	
			json.Unmarshal(body, &respData)
			fmt.Println("AI: " + respData.Text)
		}
	
		if resp != nil {
			resp.Body.Close()
		}
	}
}

func main() {
	var input string

	channel := make(chan string, 1)
	defer close(channel)

	go process(channel, string(rand.Int63()))

	for {
		fmt.Scanf("%s", &input)

		channel <- input

		if input == "EOF" {
			fmt.Println("Bye!")
			break
		}
	}
}

