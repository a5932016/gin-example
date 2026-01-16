package pingAPI

import (
	"github.com/a5932016/gin_example/internal/ping"
	"github.com/a5932016/gin_example/internal/ping/pingRequest"
	"github.com/a5932016/gin_example/pkg/layout"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	websocket, request, err := layout.WebSocketRequestLayout[pingRequest.Start](c, true)
	if err != nil {
		return
	}

	layout.WebSocketLayout(c, ping.Start, 0, request, websocket)
}
