package handler

import (
	"github.com/p12s/wildberries-http-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"hello": 321,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	// в задаче не требуется делать регистр.-авторизац.-аутентиф., но заготовка пусть будет
	var input User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
