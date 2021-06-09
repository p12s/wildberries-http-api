package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createComment(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllComments(c *gin.Context) {

}

func (h *Handler) getCommentById(c *gin.Context) {

}

func (h *Handler) updateComment(c *gin.Context) {

}

func (h *Handler) deleteComment(c *gin.Context) {

}
