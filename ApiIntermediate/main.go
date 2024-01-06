package main

import (
	"apiintermediate/models"
	"apiintermediate/req"
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var logger = log.Default()

func main() {
	u := url.URL{Scheme: "ws", Host: ":8080", Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	defer c.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		log.Printf("Message received: %s", message)
		dataElectricity := models.ElectronicDevice{}
		if err := json.Unmarshal(message, &dataElectricity); err != nil {
			panic(err)
		}
		messageServer, errGrpc := req.CreateElectricityRpc(dataElectricity)
		if errGrpc != nil {
			panic(errGrpc)
		}
		if err := c.WriteMessage(websocket.TextMessage, []byte(*messageServer)); err != nil {
			panic(err)
		}
	}
}
