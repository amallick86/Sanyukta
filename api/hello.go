package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

// hello world api
// @Summary print hello world
// @Tags Hello
// @ID Hello
// @Produce json
// @Success 200 {object} HelloResponse
// @Router /hello [get]
func (server *Server) Hello(c *gin.Context) {
	msg := HelloResponse{
		Message: "Hello World",
	}

	c.JSON(http.StatusOK, msg)
}
