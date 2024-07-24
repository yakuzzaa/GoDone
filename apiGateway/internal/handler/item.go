package handler

import (
	"net/http"
	"strconv"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/converter"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	item "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
)

// createItem creates an item for a to-do list.
// @Summary Create Item
// @Description Create item for to-do list
// @Tags item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List ID"
// @Param   request body  serializer.CreateItemRequest true "CreateItemRequest"
// @Success 200 {object} serializer.CreateItemResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/items/{list_id} [post]
func (h *ApiHandler) createItem(c *gin.Context) {
	startTime := time.Now()
	var req item.CreateRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("createItem: failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Warn("createItem: invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	resp, err := h.itemClient.CreateItem(c, &req)
	if err != nil {
		h.logger.Error("createItem: failed to create item", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("createItem: item created successfully", slog.Uint64("itemId", resp.Id), slog.Duration("duration", time.Since(startTime)))

	c.JSON(http.StatusOK, serializer.CreateItemResponse{
		Id: resp.Id,
	})
}

// getAllItems retrieves all items for a list.
// @Summary Get List Items
// @Description Get List Items
// @Tags item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List ID"
// @Success 200 {object} serializer.ListItemResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/items/{list_id} [get]
func (h *ApiHandler) getAllItems(c *gin.Context) {
	startTime := time.Now()
	var req item.ListRequest
	var err error

	idStr := c.Param("list_id")
	req.ListId, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("getAllItems: failed to parse list ID", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Warn("getAllItems: invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	resp, err := h.itemClient.ListItem(c, &req)
	if err != nil {
		h.logger.Error("getAllItems: failed to list items", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalGetItemResponse(resp)
	if err != nil {
		h.logger.Error("getAllItems: failed to marshal response", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("getAllItems: retrieved items successfully", slog.Duration("duration", time.Since(startTime)))

	c.JSON(http.StatusOK, respJson)
}

// getItemById retrieves an item by its ID.
// @Summary Get Item by id
// @Description Get List Items
// @Tags item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List ID"
// @Param item_id path string true "Item ID"
// @Success 200 {object} serializer.GetItemResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/items/{list_id}/{item_id} [get]
func (h *ApiHandler) getItemById(c *gin.Context) {
	startTime := time.Now()
	var req item.GetRequest
	var err error

	listIdStr := c.Param("list_id")
	req.ListId, err = strconv.ParseUint(listIdStr, 10, 64)
	itemIdStr := c.Param("item_id")
	req.Id, err = strconv.ParseUint(itemIdStr, 10, 64)

	if err != nil {
		h.logger.Error("getItemById: failed to parse IDs", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Warn("getItemById: invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	resp, err := h.itemClient.GetItem(c, &req)
	if err != nil {
		h.logger.Error("getItemById: failed to get item", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalDetailItemResponse(resp)
	if err != nil {
		h.logger.Error("getItemById: failed to marshal response", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("getItemById: retrieved item successfully", slog.Uint64("itemId", req.Id), slog.Duration("duration", time.Since(startTime)))

	c.JSON(http.StatusOK, respJson)
}

// updateItem updates an item by its ID.
// @Summary Update Item by id
// @Description Get List Items
// @Tags item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List ID"
// @Param item_id path string true "Item ID"
// @Param   request body  serializer.UpdateItemRequest true "UpdateItemRequest"
// @Success 200 {object} serializer.UpdateItemResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/items/{list_id}/{item_id} [put]
func (h *ApiHandler) updateItem(c *gin.Context) {
	startTime := time.Now()
	var req item.UpdateRequest
	var err error

	listIdStr := c.Param("list_id")
	req.ListId, err = strconv.ParseUint(listIdStr, 10, 64)
	itemIdStr := c.Param("item_id")
	req.Id, err = strconv.ParseUint(itemIdStr, 10, 64)

	if err != nil {
		h.logger.Error("updateItem: failed to parse IDs", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Warn("updateItem: invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("updateItem: failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	_, err = h.itemClient.UpdateItem(c, &req)
	if err != nil {
		h.logger.Error("updateItem: failed to update item", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("updateItem: item updated successfully", slog.Uint64("itemId", req.Id), slog.Duration("duration", time.Since(startTime)))

	c.JSON(http.StatusOK, serializer.UpdateItemResponse{
		Status: "Updated",
	})
}

// deleteItem deletes an item by its ID.
// @Summary Delete Item by id
// @Description Delete Item by id
// @Tags item
// @Accept  json
// @Produce  json
// @Param list_id path string true "List ID"
// @Param item_id path string true "Item ID"
// @Success 200 {object} serializer.UpdateItemResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /api/items/{list_id}/{item_id} [delete]
func (h *ApiHandler) deleteItem(c *gin.Context) {
	startTime := time.Now()
	var req item.DeleteRequest
	var err error

	listIdStr := c.Param("list_id")
	req.ListId, err = strconv.ParseUint(listIdStr, 10, 64)
	itemIdStr := c.Param("item_id")
	req.Id, err = strconv.ParseUint(itemIdStr, 10, 64)

	if err != nil {
		h.logger.Error("deleteItem: failed to parse IDs", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		h.logger.Warn("deleteItem: invalid token")
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
		return
	}
	req.UserId = userId.(uint64)

	_, err = h.itemClient.DeleteItem(c, &req)
	if err != nil {
		h.logger.Error("deleteItem: failed to delete item", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("deleteItem: item deleted successfully", slog.Uint64("itemId", req.Id), slog.Duration("duration", time.Since(startTime)))

	c.JSON(http.StatusOK, serializer.DeleteItemResponse{
		Status: "Deleted",
	})
}
