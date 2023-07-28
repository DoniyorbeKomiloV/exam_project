package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTarif godoc
// @ID create_tarif
// @Router /tarif [POST]
// @Summary Create Tarif
// @Description Create Tarif
// @Tags Tarif
// @Accept json
// @Procedure json
// @Param user body models.CreateTarif true "CreateTarifRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateTarif(c *gin.Context) {
	var createTarif *models.CreateTarif
	err := c.ShouldBindJSON(&createTarif)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	TarifId, err := h.strg.StaffTarif().Create(c.Request.Context(), createTarif)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Tarif, err := h.strg.StaffTarif().GetById(c.Request.Context(), &models.TarifPrimaryKey{Id: TarifId})
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
		"data":    Tarif,
	})
}

// UpdateTarif godoc
// @ID update_tarif
// @Router /tarif [PUT]
// @Summary Update Tarif
// @Description Update Tarif
// @Tags Tarif
// @Accept json
// @Procedure json
// @Param user body models.Tarif true "UpdateTarifRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateTarif(c *gin.Context) {
	var tarif models.UpdateTarif
	err := c.ShouldBindJSON(&tarif)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.StaffTarif().Update(c.Request.Context(), &tarif)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while update tarif",
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

// GetByIdTarif godoc
// @ID get_by_id_tarif
// @Router /tarif/{id} [GET]
// @Summary Get By Id Tarif
// @Description Get by id Tarif
// @Tags Tarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdTarif(c *gin.Context) {
	var id = c.Param("id")
	tarif, err := h.strg.StaffTarif().GetById(c.Request.Context(), &models.TarifPrimaryKey{Id: id})
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
		"message": "Tarif found",
		"data":    tarif,
	})
}

// GetListTarif godoc
// @ID get_list_tarif
// @Router /tarif [GET]
// @Summary Get List Tarif
// @Description Get List Tarif
// @Tags Tarif
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListTarif(c *gin.Context) {
	resp, err := h.strg.StaffTarif().GetList(c.Request.Context(), &models.TarifGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while getListTarif",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list tarif response",
		"data":    resp,
	})
}

// DeleteTarif godoc
// @ID delete_tarif
// @Router /tarif/{id} [DELETE]
// @Summary Delete Tarif
// @Description Delete Tarif
// @Tags Tarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteTarif(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.StaffTarif().Delete(c.Request.Context(), &models.TarifPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while delete tarif",
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
