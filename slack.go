package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type responseRTMStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	URL   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	ID string `json:"id"`
}

func slackStart(token string) (wsURL, id string, err error) {
	URL := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", botToken)
	resp, err := http.Get(URL)
	if err != nil {
		log.Printf("error establishing connection %s:", err)
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	respObj := new(responseRTMStart)
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return
	}

	if !respObj.Ok {
		err = fmt.Errorf("Slack error: %s", respObj.Error)
		return
	}

	wsURL = respObj.URL
	id = respObj.Self.ID
	return
}

type Message struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func getMessage(ws *websocket.Conn) (m Message, err error) {
	err = websocket.JSON.Receive(ws, &m)
	return
}

func postMessage(ws *websocket.Conn, m Message) error {
	return websocket.JSON.Send(ws, m)
}

// Begins Real-Time API session and returns the websocket and the ID of the bot
func slackConnect() (*websocket.Conn, string) {
	websocketURL, ID, err := slackStart(botToken)
	if err != nil {
		log.Fatal(err)
	}

	webSocket, err := websocket.Dial(websocketURL, "", "https://api.slack.com/")
	if err != nil {
		log.Fatal(err)
	}

	return webSocket, ID
}
