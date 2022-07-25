package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get("pi.openweathermap.org/data/3.0/onecall/timemachine/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
