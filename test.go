package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8000/user/update/2"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(`{"usernmae:"PATCH 1","first_name":"PATCH 2","last_name":"PATCH 3"}`)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}