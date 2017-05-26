package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// start Real Time API session
	slackSock, botID := slackConnect()
	fmt.Println("Connected")
	fmt.Println(botID)

	for {
		// read each incoming message
		message, err := getMessage(slackSock)
		if err != nil {
			log.Printf("error retrieving message: %s", err)
		}

		if message.Type == "message" && strings.HasPrefix(message.Text, "<@"+botID+">") {
			//clean up text by removing the prefix
			message.Text = strings.Trim(message.Text, "<@"+botID+"> ")

			//command block
			switch message.Text {
			case "office":
				unregCount, regList := pollOffice()
				message.Text = fmt.Sprintf("Unregistered: %d, Registered: %s", unregCount, regList[0])
				postMessage(slackSock, message)
			case "spy":
				message.Text = "I'm already doing that."
				postMessage(slackSock, message)
			default:
				message.Text = "what are you doin'?"
				postMessage(slackSock, message)
			}

		}

	}

}
