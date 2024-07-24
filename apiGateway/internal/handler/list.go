package handler

import (
	"net/http"
	"strconv"

	"log/slog"

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
		h.logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Error("invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	h.logger.Info("createList request received", slog.Uint64("userId", req.UserId))

	resp, err := h.listClient.CreateList(c, &req)
	if err != nil {
		h.logger.Error("failed to create list", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("createList request successful", slog.Uint64("listId", resp.Id))

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
		h.logger.Error("invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	h.logger.Info("getAllLists request received", slog.Uint64("userId", req.UserId))

	resp, err := h.listClient.List(c, &req)
	if err != nil {
		h.logger.Error("failed to get lists", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalListGetResponse(resp)
	if err != nil {
		h.logger.Error("failed to marshal response", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("getAllLists request successful")

	c.JSON(http.StatusOK, respJson)
}

// getListById get user to-do list by id.
// @Summary Get List by id
// @Description Get user to-do list by id
// @Tags list
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Success 200 {object} serializer.CreateListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [get]
func (h *ApiHandler) getListById(c *gin.Context) {
	var req list.DetailRequest
	var err error
	idStr := c.Param("id")
	req.Id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("failed to parse ID", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Error("invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	h.logger.Info("getListById request received", slog.Uint64("userId", req.UserId), slog.Uint64("listId", req.Id))

	resp, err := h.listClient.GetDetail(c, &req)
	if err != nil {
		h.logger.Error("failed to get list detail", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalDetailListResponse(resp)
	if err != nil {
		h.logger.Error("failed to marshal response", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("getListById request successful")

	c.JSON(http.StatusOK, respJson)
}

// updateList update user to-do list.
// @Summary Update List
// @Description Update user to-do list
// @Tags list
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Param   request body  serializer.UpdateListRequest true "UpdateListRequest"
// @Success 200 {object} serializer.UpdateListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [put]
func (h *ApiHandler) updateList(c *gin.Context) {
	var req list.UpdateRequest
	var err error
	idStr := c.Param("id")
	req.Id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("failed to parse ID", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Error("invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("updateList request received", slog.Uint64("userId", req.UserId), slog.Uint64("listId", req.Id))

	_, err = h.listClient.UpdateList(c, &req)
	if err != nil {
		h.logger.Error("failed to update list", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("updateList request successful", slog.Uint64("listId", req.Id))

	c.JSON(http.StatusOK, serializer.UpdateListResponse{
		Status: "Updated",
	})
}

// deleteList delete user to-do list.
// @Summary Delete List
// @Description Delete user to-do list
// @Tags list
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Success 200 {object} serializer.DeleteListResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/lists/{id} [delete]
func (h *ApiHandler) deleteList(c *gin.Context) {
	var req list.DeleteRequest
	var err error
	idStr := c.Param("id")
	req.Id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("failed to parse ID", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Error("invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	h.logger.Info("deleteList request received", slog.Uint64("userId", req.UserId), slog.Uint64("listId", req.Id))

	_, err = h.listClient.DeleteList(c, &req)
	if err != nil {
		h.logger.Error("failed to delete list", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("deleteList request successful", slog.Uint64("listId", req.Id))

	c.JSON(http.StatusOK, serializer.DeleteListResponse{
		Status: "Deleted",
	})
}
