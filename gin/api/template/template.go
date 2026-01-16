package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebSocket(c *gin.Context) {
	c.HTML(http.StatusOK, "websocket.html", nil)
}
