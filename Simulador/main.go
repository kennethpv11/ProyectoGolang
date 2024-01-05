package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simulador/models"
	"simulador/utils"
	"time"

	"github.com/gorilla/websocket"
)

var logger = log.Default()
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func producerData(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		values, err := utils.GenerateDataElectricity()

		if err != nil {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(err.Error())); err != nil {
				panic(err)
			}
			return
		}
		dataLog := models.ElectronicDevice{}
		err = json.Unmarshal(values, &dataLog)
		if err != nil {
			panic(err)
		}
		logger.Print("electricity info:", dataLog)
		if err := conn.WriteMessage(websocket.TextMessage, values); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}

}

func main() {

	logger.Print("Listening and serving HTTP on :8080 ")

	http.HandleFunc("/", producerData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
