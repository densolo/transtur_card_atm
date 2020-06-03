package server

import (
	"log"
	"encoding/json"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

type CardMessage struct {
	CardStateColor string `json:"card_state_color"`
	CardStateText string  `json:"card_state_text"`
}

type StateHandler interface {
	SendGreyState(text string)
	SendRedState(text string)
	SendBlueState(text string)
}

type ElectronStateHandler struct {
}

func (state ElectronStateHandler) SendRedState(text string) {
	SendCardState("danger", text)
}

func (state ElectronStateHandler) SendGreyState(text string) {
	SendCardState("secondary", text)
}

func (state ElectronStateHandler) SendBlueState(text string) {
	SendCardState("primary", text)
}


func SendCardState(color string, text string) {
	log.Printf("SendCardState: %s - %s", color, text)

	msg := CardMessage{
		CardStateColor: color,
		CardStateText: text,
	}
	
	msg_json, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		log.Printf("Failed to serialize json: %s", err.Error())
		return
	}

	log.Printf("SendCardState json: %s", string(msg_json))
	appWindow := GetGuiWindow()
	if appWindow != nil {
		bootstrap.SendMessage(appWindow, "update", string(msg_json))
	}
}


var (
	GlobalStateHandler = ElectronStateHandler{}
)
