package layout

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type KeepConnectModel struct {
	Output   chan []byte
	Ctx      context.Context
	Cancel   context.CancelFunc
	Err      chan error
	Duration time.Duration
}

func NewKeepConnectModel(duration time.Duration) *KeepConnectModel {
	ctx, cancel := context.WithCancel(context.Background())
	return &KeepConnectModel{
		Output:   make(chan []byte, 1),
		Ctx:      ctx,
		Cancel:   cancel,
		Err:      make(chan error, 1),
		Duration: duration,
	}
}

func WebSocketRequestLayout[T any](c *gin.Context, isInput bool) (*websocket.Conn, T, error) {
	var request T

	// start create webSocket
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	webSocket, err := upGrader.Upgrade(c.Writer, c.Request, http.Header{})
	if err != nil {
		return webSocket, request, err
	}
	// end create webSocket

	// start read iperfResponse.ClientStart
	if isInput {
		_, msg, err := webSocket.ReadMessage()
		if err != nil {
			return webSocket, request, err
		}

		if err := json.Unmarshal(msg, &request); err != nil {
			return webSocket, request, err
		}
	}
	// end read iperfResponse.ClientStart

	return webSocket, request, nil
}

func WebSocketLayout[T any](c *gin.Context, f func(T, *KeepConnectModel), duration time.Duration, request T, webSocket *websocket.Conn) {
	defer webSocket.Close()

	var input = NewKeepConnectModel(duration)
	defer input.Cancel()

	go f(request, input)

	go func() {
		defer input.Cancel()

		for {
			select {
			case <-input.Ctx.Done():
				return
			default:
				_, _, err := webSocket.ReadMessage()
				if err != nil {
					fmt.Printf("error: %v\n", err)
					return
				}

				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					return
				}
			}
		}
	}()

	for {
		select {
		case <-input.Ctx.Done():
			return
		case val := <-input.Output:
			err := webSocket.WriteMessage(websocket.TextMessage, val)
			if err != nil {
				fmt.Printf("error: %v\n", val)
				return
			}
		case val := <-input.Err:
			fmt.Printf("websocket error: %v\n", val)
			return
		}
	}
}
