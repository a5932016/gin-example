package routes

import (
	"github.com/a5932016/gin_example/gin/api/userAPI"
	"github.com/gin-gonic/gin"
)

func APIReute(r *gin.Engine) {
	apiV1Group := r.Group("/api/v1/")

	userGroup := apiV1Group.Group("/user")
	{
		userGroup.GET("/list", userAPI.List)
	}
}
