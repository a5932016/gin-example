package services

import (
	"github.com/a5932016/gin_example/gin/routes"
	"github.com/a5932016/gin_example/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RunWebService() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	routes.TemplateRoute(r)

	routes.APIReute(r)

	routes.WebSocketReute(r)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for key, value := range validators.ValidatorsMap {
			v.RegisterValidation(key, value)
		}
	}

	return r
}
