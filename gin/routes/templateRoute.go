package routes

import (
	"github.com/a5932016/gin_example/gin/api/template"
	"github.com/gin-gonic/gin"
)

func TemplateRoute(r *gin.Engine) {
	templateGroup := r.Group("/template")

	templateGroup.GET("/websocket", template.WebSocket)
}
