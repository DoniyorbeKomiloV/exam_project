package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateSales godoc
// @ID create_sales
// @Router /sales [POST]
// @Summary Create Sales
// @Description Create Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param user body models.CreateSales true "CreateSalesRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateSales(c *gin.Context) {
	var createSales *models.CreateSales
	err := c.ShouldBindJSON(&createSales)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	SalesId, err := h.strg.Sales().Create(c.Request.Context(), createSales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Sales, err := h.strg.Sales().GetById(c.Request.Context(), &models.SalesPrimaryKey{Id: SalesId})
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
		"data":    Sales,
	})
}

// UpdateSales godoc
// @ID update_sales
// @Router /sales [PUT]
// @Summary Update Sales
// @Description Update Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param user body models.Sales true "UpdateSalesRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateSales(c *gin.Context) {
	var sales models.UpdateSales
	err := c.ShouldBindJSON(&sales)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Sales().Update(c.Request.Context(), &sales)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while update sales",
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

// GetByIdSales godoc
// @ID get_by_id_sales
// @Router /sales/{id} [GET]
// @Summary Get By Id Sales
// @Description Get by id Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdSales(c *gin.Context) {
	var id = c.Param("id")
	sales, err := h.strg.Sales().GetById(c.Request.Context(), &models.SalesPrimaryKey{Id: id})
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
		"message": "Sales found",
		"data":    sales,
	})
}

// GetListSales godoc
// @ID get_list_sales
// @Router /sales [GET]
// @Summary Get List Saleses
// @Description Get List Saleses
// @Tags Sales
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListSales(c *gin.Context) {
	resp, err := h.strg.Sales().GetList(c.Request.Context(), &models.SalesGetListRequest{
		Offset: 0,
		Limit:  10,
		Search: c.Query("search"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while getListSales",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list sales response",
		"data":    resp,
	})
}

// DeleteSales godoc
// @ID delete_sales
// @Router /sales/{id} [DELETE]
// @Summary Delete Sales
// @Description Delete Sales
// @Tags Sales
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteSales(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Sales().Delete(c.Request.Context(), &models.SalesPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while delete sales",
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
