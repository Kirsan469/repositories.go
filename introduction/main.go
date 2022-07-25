package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

	url := "https://api-football-v1.p.rapidapi.com/v3/timezone"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "07ee356bf3mshb3f33ef73f6db25p15586djsn9878a111fec8")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
