package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	var has hasContactRequest
	has.AddAuth("123abc", "h4jsu6pc5txybj04", strconv.FormatInt(time.Now().Unix(), 10), "4f69782e826841f794080cae87648e42")
	has.InitHasContactRequest("admin@vendor.pl", "piotrek.uryga+7574@onet.pl")

	b := new(bytes.Buffer)

	error := json.NewEncoder(b).Encode(has)
	if error != nil {
		log.Fatal(error)
	}

	req, error := http.NewRequest("POST", "https://www.salesmanago.pl/api/contact/hasContact", b)
	req.Header.Set("Accept", "application/json, application/json")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{Timeout: 30 * time.Second}
	response, error := client.Do(req)

	//response, error := http.Post("https://www.salesmanago.pl/api/contact/hasContact", "application/json", b )
	if error != nil {
		log.Fatalf("Cannot post")
	}
	defer response.Body.Close()
	bodyBytes, error := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	var s = new(hasContactResponse)
	err := json.Unmarshal(bodyBytes, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	fmt.Println(bodyString)

}
