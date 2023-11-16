package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"os"
)

var Out = make(chan any, 9999)
var socket net.Conn
var config Config

func configLoad() {
	configFile, _ := os.ReadFile("config.json")
	err := json.Unmarshal(configFile, &config)

	if err != nil {
		log.Fatalln("Error while read config file", err)
	}
}

func connect() {
	var err error
	if socket, err = net.Dial("tcp", fmt.Sprintf("%s:%d", config.Servers.Host.Server.Address, config.Servers.Host.Server.Port)); err != nil {
		log.Println("Connection to HOST server error: ", err)
		// @TODO run reconnection procedure
	}

	for {
		msg := <-Out
		msgString, _ := json.Marshal(msg)
		_, err = fmt.Fprintln(socket, string(msgString))

		if err != nil {
			// @TODO Reconnect
			return
		}
	}
}

func main() {
	configLoad()
	//connect()
	wsc()
}

func wsc() {
	ws, _, err := websocket.DefaultDialer.Dial("wss://chat-1.goodgame.ru/chat2/", nil)
	if err != nil {
		log.Fatal("GoodGame websocket server connection error", err)
	}

	defer ws.Close()

	for {
		var message GoodgameMessage
		err := ws.ReadJSON(&message)

		if err != nil {
			log.Println("WS read error:", err)
			return
		}

		switch message.Type {
		case "welcome":
			request := GoodgameJoinRequest{
				Type: "join",
				Data: GoodgameJoinRequestData{
					ChannelId: "9126",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			ws.WriteJSON(&request)

			request = GoodgameJoinRequest{
				Type: "join",
				Data: GoodgameJoinRequestData{
					ChannelId: "53029",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			ws.WriteJSON(&request)

		case "success_join":
			fmt.Println("SUCCESS JOIN")

		case "channel_counters":
			fmt.Println("COUNTERS")

		case "message":
			log.Printf("[%s]: %s", message.Data.UserName, message.Data.Text)
		}
	}

}
