package userAPI

import (
	"net/http"

	"github.com/a5932016/gin_example/internal/user"
	"github.com/a5932016/gin_example/internal/user/userRequest"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var request userRequest.List
	if err := c.ShouldBindQuery(&request); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	response, err := user.List(request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
