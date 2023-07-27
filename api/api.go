package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageInterface, logger logger.LoggerI) {
	NewHandler := handler.NewHandler(cfg, storage, logger)

	r.POST("/branches", NewHandler.CreateBranch)
	r.GET("/branches/:id", NewHandler.GetByIdBranch)
	r.GET("/branches", NewHandler.GetListBranch)
	r.PUT("/branches", NewHandler.UpdateBranch)
	r.DELETE("/branches/:id", NewHandler.DeleteBranch)

	r.POST("/tarif", NewHandler.CreateTarif)
	r.GET("/tarif/:id", NewHandler.GetByIdTarif)
	r.GET("/tarif", NewHandler.GetListTarif)
	r.PUT("/tarif", NewHandler.UpdateTarif)
	r.DELETE("/tarif/:id", NewHandler.DeleteTarif)

	r.POST("/staff", NewHandler.CreateStaff)
	r.GET("/staff/:id", NewHandler.GetByIdStaff)
	r.GET("/staff", NewHandler.GetListStaff)
	r.PUT("/staff", NewHandler.UpdateStaff)
	r.DELETE("/staff/:id", NewHandler.DeleteStaff)

	r.POST("/sales", NewHandler.CreateSales)
	r.GET("/sales/:id", NewHandler.GetByIdSales)
	r.GET("/sales", NewHandler.GetListSales)
	r.PUT("/sales", NewHandler.UpdateSales)
	r.DELETE("/sales/:id", NewHandler.DeleteSales)

	r.POST("/transactions", NewHandler.CreateTransaction)
	r.GET("/transactions/:id", NewHandler.GetByIdTransaction)
	r.GET("/transactions", NewHandler.GetListTransaction)
	r.PUT("/transactions", NewHandler.UpdateTransaction)
	r.DELETE("/transactions/:id", NewHandler.DeleteTransaction)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
