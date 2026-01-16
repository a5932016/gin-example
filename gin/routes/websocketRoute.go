package routes

import (
	"github.com/a5932016/gin_example/gin/api/pingAPI"
	"github.com/gin-gonic/gin"
)

func WebSocketReute(r *gin.Engine) {
	websocketV1Group := r.Group("/websocket/v1/")

	websocketV1Group.GET("/start", pingAPI.Start)
}
