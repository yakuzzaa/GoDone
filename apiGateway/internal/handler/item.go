package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/converter"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	item "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/item_v1"
)

// createItem create item for to-do list.
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
	var req item.CreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}
	req.UserId = userId.(uint64)

	resp, err := h.itemClient.CreateItem(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, serializer.CreateItemResponse{
		Id: resp.Id})
}

// getAllItems get list of items for list.
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
	var req item.ListRequest
	var err error
	idStr := c.Param("list_id")
	req.ListId, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}
	req.UserId = userId.(uint64)

	resp, err := h.itemClient.ListItem(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalGetItemResponse(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, respJson)
}

// getItemById get item by id
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
	var req item.GetRequest
	var err error
	idStr := c.Param("list_id")
	req, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Message: "Invalid Token"})
	}
	req.UserId = userId.(uint64)

	resp, err := h.listClient.GetDetail(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	respJson, err := converter.MarshalDetailListResponse(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, respJson)
}

// updateItem update item by id
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

}

// deleteItem update item by id
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

}
