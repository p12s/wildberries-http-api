package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/p12s/wildberries-http-api"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input common.User
	fmt.Println(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)

	id, err := h.services.CreateUser(input)
	fmt.Println("id", id)
	fmt.Println("err", err)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}


	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
}
