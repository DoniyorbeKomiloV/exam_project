package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateStaff godoc
// @ID create_staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param user body models.CreateStaff true "CreateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateStaff(c *gin.Context) {
	var createStaff *models.CreateStaff
	err := c.ShouldBindJSON(&createStaff)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	StaffId, err := h.strg.Staff().Create(c.Request.Context(), createStaff)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Staff, err := h.strg.Staff().GetById(c.Request.Context(), &models.StaffPrimaryKey{Id: StaffId})
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
		"message": "Staff created",
		"data":    Staff,
	})
}

// UpdateStaff godoc
// @ID update_staff
// @Router /staff [PUT]
// @Summary Update Staff
// @Description Update Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param user body models.Staff true "UpdateStaffRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateStaff(c *gin.Context) {
	var staff models.UpdateStaff
	err := c.ShouldBindJSON(&staff)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Staff().Update(c.Request.Context(), &staff)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while update staff",
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

// GetByIdStaff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [GET]
// @Summary Get By Id Staff
// @Description Get by id Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdStaff(c *gin.Context) {
	var id = c.Param("id")
	staff, err := h.strg.Staff().GetById(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
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
		"message": "Staff found",
		"data":    staff,
	})
}

// GetListStaff godoc
// @ID get_list_staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Procedure jsonStaff
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListStaff(c *gin.Context) {
	resp, err := h.strg.Staff().GetList(c.Request.Context(), &models.StaffGetListRequest{
		Offset: 0,
		Limit:  10,
		Search: c.Query("search"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while getListStaff",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list staff response",
		"data":    resp,
	})
}

// DeleteStaff godoc
// @ID delete_staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteStaff(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Staff().Delete(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while delete staff",
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
