package main

import (
	"encoding/json"
	"fmt"
	"log"

	"time"

	"golang.org/x/net/websocket"
)

const NOTI_PING = "PING"

type PingMessage struct {
	Action          string `json:"action"` // PING
	NotificationKey string `json:"notification"`
}

func main() {
	const connectionCount = 300
	fmt.Printf("making %d connections.\n", connectionCount)
	origin := "https://www.linkedin.com"
	url := "ws://localhost:8080/connect?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImlzcyI6IklETSJ9.eyJpc3MiOiJJRE0iLCJpYXQiOjE0ODY2NzU3NDMsImV4cCI6MTQ4Njc0MDU0MywiaWRtaWQiOjQ0MjczLCJpZG1HdWlkIjoiNDQyNzMiLCJjcm10eXBlIjpudWxsLCJzdGFzdWIiOm51bGwsIm9yZ2lkIjoyMDA3MTUxMywiY3JtdWlkIjpudWxsLCJhYmwiOmZhbHNlLCJ0eWtfa2V5X3BvbGljeV9pZCI6IjU2YjIyNWQ1ZWMwNjU5MWQ3NTAwMDAwNiIsImxpY2Vuc2VzIjpbXSwiY2FwYWJpbGl0aWVzIjpbIlBsYXlib29rcyIsIlZpc2lvbiIsIlBvd2VyU3RhbmRpbmdzIiwiSW5zaWdodHMiLCJMb2NhbFByZXNlbmNlIl19.Xw-JiRZx_kg0EiiNwDxU5nj05tZYOx5GTK-7ElQhwHNyx93OVsY5ZS_h5u0uSVcnYipVe4hj8r56MhO1N0l-7l7zTW776Tou0ZZAvmS0uZaeoxtPUY3lCmhj0sp9HB85q6ee9_x_MwmpOKd8oUyZkQoJMKcxix-tw5VLRFUyxTO2DRxLLyzdXYmS2l54XWQDncbICGA1JZRsPiPVy7BwiCnJGSTD_PzsT5z7dZJJZPyeTTtnF46v676osk9NC7KDWhukTXICB8XO9ZG1ig8PPb15uLx84v1z3gVZI6302_YZn3ZAQ3pdHO_L0x4l4Evrj0v9Kp6wYEyOBonni3F6ZA"

	var connects [connectionCount]*websocket.Conn
	for i := 0; i < connectionCount; i++ {
		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal(err)
		}
		connects[i] = ws

		go receive(ws, i)
		go sendPing(ws, i)
	}

	time.Sleep(120 * time.Second)

	for j := 0; j < connectionCount; j++ {
		err := connects[j].Close()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func receive(ws *websocket.Conn, recvId int) {
	var msg string
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d Received: %s.\n", recvId, msg)
	}

	closeErr := ws.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
}

func sendPing(ws *websocket.Conn, sendId int) {

	for {
		time.Sleep(10 * time.Second)

		var msg PingMessage
		msg.Action = NOTI_PING
		msg.NotificationKey = fmt.Sprintf("%d: %s", sendId, time.Now().String())
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}

		sendErr := websocket.Message.Send(ws, msgBytes)
		if sendErr != nil {
			log.Fatal(sendErr)
		}
	}
}
