package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type officeResponse struct {
	Others     int      `json:"others"`
	Registered []string `json:"registered"`
}

func pollOffice() (int, []string) {
	res, err := http.Get(officeURL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer res.Body.Close()

	response := new(officeResponse)
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Fatalf("failed to decode JSON: %v", err)
	}

	return response.Others, response.Registered
}
