package handler

import (
	todo "To-do-list"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var inp todo.User
	if err := c.BindJSON(&inp); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(inp)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
