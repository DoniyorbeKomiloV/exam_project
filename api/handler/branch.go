package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBranch godoc
// @ID create_branch
// @Router /branches [POST]
// @Summary Create Branch
// @Description Create Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param user body models.CreateBranch true "CreateBranchRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateBranch(c *gin.Context) {
	var createBranch *models.CreateBranch
	err := c.ShouldBindJSON(&createBranch)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	BranchId, err := h.strg.Branch().Create(c.Request.Context(), createBranch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Branch, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPrimaryKey{Id: BranchId})
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
		"data":    Branch,
	})
}

// UpdateBranch godoc
// @ID update_branch
// @Router /branches [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param user body models.Branch true "UpdateBranchRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.UpdateBranch
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Branch().Update(c.Request.Context(), &branch)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while update branch",
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

// GetByIdBranch godoc
// @ID get_by_id_branch
// @Router /branches/{id} [GET]
// @Summary Get By Id Branch
// @Description Get by id Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdBranch(c *gin.Context) {
	var id = c.Param("id")
	branch, err := h.strg.Branch().GetById(c.Request.Context(), &models.BranchPrimaryKey{Id: id})
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
		"message": "Branch found",
		"data":    branch,
	})
}

// GetListBranch godoc
// @ID get_list_branch
// @Router /branches [GET]
// @Summary Get List Branches
// @Description Get List Branches
// @Tags Branch
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListBranch(c *gin.Context) {
	resp, err := h.strg.Branch().GetList(c.Request.Context(), &models.BranchGetListRequest{
		Offset:        0,
		Limit:         10,
		SearchName:    c.Query("search_name"),
		SearchAddress: c.Query("search_address"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while getListBranch",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list branch response",
		"data":    resp,
	})
}

// DeleteBranch godoc
// @ID delete_branch
// @Router /branches/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteBranch(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Branch().Delete(c.Request.Context(), &models.BranchPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while delete branch",
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
