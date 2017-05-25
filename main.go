package main

import (
	"fmt"
	"log"
)

func main() {
	// start Real Time API session
	ws, id := slackConnect()
	fmt.Println("Connected")
	fmt.Println(id)

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if m.Type == "message" {
			fmt.Printf("msg")
			pollOffice()
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
