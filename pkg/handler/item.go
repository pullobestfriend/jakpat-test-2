package handler

import (
	"jakpat-test-2/entity"
	"jakpat-test-2/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddItem(c *gin.Context) {
	var item entity.Items

	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	id, err := h.usecases.AddItem(*user, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetItemByIdAndStatus(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	item, err := h.usecases.GetItemByIdAndStatus(itemId, status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) UpdateItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input entity.Items
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	err = h.usecases.UpdateItemById(*user, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}

func (h *Handler) GetItemsBySellerIdAndStatus(c *gin.Context) {
	sellerId, err := strconv.Atoi(c.Param("seller"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid seller id param")
		return
	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	item, err := h.usecases.GetItemsBySellerIdAndStatus(*user, sellerId, status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) CreateOrder(c *gin.Context) {
	var order entity.Oders

	if err := c.BindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	id, err := h.usecases.CreateOrder(*user, order.ItemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetOrderById(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	item, err := h.usecases.GetOrderById(*user, orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) UpdateOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	user := c.MustGet(usecase.CtxUserKey).(*entity.Users)

	err = h.usecases.UpdateOrderStatusByIdAndStatus(*user, id, status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}
