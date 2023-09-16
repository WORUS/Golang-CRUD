package handler

import (
	"fmt"
	"net/http"

	todo "github.com/WORUS/Golang-CRUD"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	fmt.Print(id) ////////////////
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//id, err := h.services.TodoList.Create(id, input)
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
