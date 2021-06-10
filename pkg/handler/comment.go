package handler

import (
	"github.com/gin-gonic/gin"
	common "github.com/p12s/wildberries-http-api"
	"net/http"
	"strconv"
)

func (h *Handler) createComment(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		return
	}

	var input common.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Comment.Create(idUser, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCommentsResponse struct {
	Data []common.Comment `json:"data"`
}

func (h *Handler) getAllComments(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		return
	}
	comments, err := h.services.Comment.GetAll(idUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})
}

func (h *Handler) getCommentById(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	comments, err := h.services.Comment.GetById(idUser, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *Handler) updateComment(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input common.UpdateCommentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Comment.Update(idUser, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteComment(c *gin.Context) {
	idUser, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Comment.Delete(idUser, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
