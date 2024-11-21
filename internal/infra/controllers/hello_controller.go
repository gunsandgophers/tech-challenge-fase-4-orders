package controllers

import (
	"net/http"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

// Index godoc
//	@Summary		Show the index payload
//	@Description	get index payload
//	@Tags			index
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Router			/ [get]
func (h *HelloController) Index(c httpserver.HTTPContext) {
	c.JSON(http.StatusOK, httpserver.Payload{"msg": "Hello World! :)"})
}
