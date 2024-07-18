package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/converter"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	list "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/list_v1"
)

// createList create user to-do list.
// @Summary Create List
// @Description Create user to-do list
// @Tags list
// @Accept  json
// @Produce  json
// @Param   request body  serializer.CreateListRequest true "CreateListRequest"
// @Success 200 {object} serializer.CreateListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists [post]
func (h *ApiHandler) createList(c *gin.Context) {
	var req list.CreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}
	req.UserId = userId.(uint64)

	resp, err := h.listClient.CreateList(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.CreateListResponse{
		ListId: resp.Id})
}

// getAllLists get user to-do lists.
// @Summary Get Lists
// @Description Get user to-do lists
// @Tags list
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.ListsResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists [get]
func (h *ApiHandler) getAllLists(c *gin.Context) {
	var req list.ListRequest
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}
	req.UserId = userId.(uint64)

	resp, err := h.listClient.List(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalGetResponse(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, respJson)
}

// getListById get user to-do list by id.
// @Summary Get List by id
// @Description Get user to-do list by id
// @Tags list
// @Accept  json
// @Produce  json
// @Param   request body  serializer.ListByIdRequest true "ListByIdRequest"
// @Success 200 {object} serializer.CreateListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [get]
func (h *ApiHandler) getListById(c *gin.Context) {
	var req list.DetailResponse
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}

	resp, err := h.listClient.List(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalGetResponse(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, respJson)
}

// updateList update user to-do list.
// @Summary Update List
// @Description Update user to-do list
// @Tags list
// @Accept  json
// @Produce  json
// @Param   request body  serializer.UpdateListRequest true "UpdateListRequest"
// @Success 200 {object} serializer.UpdateListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [put]
func (h *ApiHandler) updateList(c *gin.Context) {

}

// deleteList delete user to-do list.
// @Summary Delete List
// @Description Delete user to-do list
// @Tags list
// @Accept  json
// @Produce  json
// @Param   request body  serializer.DeleteListRequest true "DeleteListRequest"
// @Success 200 {object} serializer.DeleteListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [delete]
func (h *ApiHandler) deleteList(c *gin.Context) {

}
