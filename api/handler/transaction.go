package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTransaction godoc
// @ID create_transaction
// @Router /transactions [POST]
// @Summary Create Transaction
// @Description Create Transaction
// @Tags Transaction
// @Accept json
// @Procedure json
// @Param user body models.CreateTransaction true "CreateTransactionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateTransaction(c *gin.Context) {
	var createTransaction *models.CreateTransaction
	err := c.ShouldBindJSON(&createTransaction)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	TransactionId, err := h.strg.StaffTransaction().Create(c.Request.Context(), createTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Transaction, err := h.strg.StaffTransaction().GetById(c.Request.Context(), &models.TransactionPrimaryKey{Id: TransactionId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return

	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "User created",
		"data":    Transaction,
	})
}

// UpdateTransaction godoc
// @ID update_transaction
// @Router /transactions [PUT]
// @Summary Update Transaction
// @Description Update Transaction
// @Tags Transaction
// @Accept json
// @Procedure json
// @Param user body models.Transaction true "UpdateTransactionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateTransaction(c *gin.Context) {
	var transaction models.UpdateTransaction
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.StaffTransaction().Update(c.Request.Context(), &transaction)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while update transaction",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"status":  "OK",
		"message": "Success",
		"data":    resp,
	})
}

// GetByIdTransaction godoc
// @ID get_by_id_transaction
// @Router /transactions/{id} [GET]
// @Summary Get By Id Transaction
// @Description Get by id Transaction
// @Tags Transaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdTransaction(c *gin.Context) {
	var id = c.Param("id")
	transaction, err := h.strg.StaffTransaction().GetById(c.Request.Context(), &models.TransactionPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Server internal Error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "Transaction found",
		"data":    transaction,
	})
}

// GetListTransaction godoc
// @ID get_list_transaction
// @Router /transactions [GET]
// @Summary Get List Transactiones
// @Description Get List Transactiones
// @Tags Transaction
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListTransaction(c *gin.Context) {
	resp, err := h.strg.StaffTransaction().GetList(c.Request.Context(), &models.TransactionGetListRequest{
		Offset: 0,
		Limit:  10,
		Search: c.Query("search"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while getListTransaction",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list transaction response",
		"data":    resp,
	})
}

// DeleteTransaction godoc
// @ID delete_transaction
// @Router /transactions/{id} [DELETE]
// @Summary Delete Transaction
// @Description Delete Transaction
// @Tags Transaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteTransaction(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.StaffTransaction().Delete(c.Request.Context(), &models.TransactionPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while delete transaction",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, map[string]interface{}{
		"status":  "OK",
		"message": "Success",
		"data":    nil,
	})
}
