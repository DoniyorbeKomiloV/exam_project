package storage

import (
	"app/api/models"
	"context"
)

type StorageInterface interface {
	Close()
	Branch() BranchRepoInterface
	StaffTarif() StaffTarifRepoInterface
	Staff() StaffRepoInterface
	Sales() SalesRepoInterface
	StaffTransaction() TransactionRepoInterface
}

type BranchRepoInterface interface {
	Create(ctx context.Context, req *models.CreateBranch) (string, error)
	Update(ctx context.Context, req *models.UpdateBranch) (int64, error)
	GetById(ctx context.Context, req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(ctx context.Context, req *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Delete(ctx context.Context, req *models.BranchPrimaryKey) error
}

type StaffTarifRepoInterface interface {
	Create(ctx context.Context, req *models.CreateTarif) (string, error)
	Update(ctx context.Context, req *models.UpdateTarif) (int64, error)
	GetById(ctx context.Context, req *models.TarifPrimaryKey) (*models.Tarif, error)
	GetList(ctx context.Context, req *models.TarifGetListRequest) (*models.TarifGetListResponse, error)
	Delete(ctx context.Context, req *models.TarifPrimaryKey) error
}

type StaffRepoInterface interface {
	Create(ctx context.Context, req *models.CreateStaff) (string, error)
	Update(ctx context.Context, req *models.UpdateStaff) (int64, error)
	GetById(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(ctx context.Context, req *models.StaffGetListRequest) (*models.StaffGetListResponse, error)
	Delete(ctx context.Context, req *models.StaffPrimaryKey) error
}

type SalesRepoInterface interface {
	Create(ctx context.Context, req *models.CreateSales) (string, error)
	Update(ctx context.Context, req *models.UpdateSales) (int64, error)
	GetById(ctx context.Context, req *models.SalesPrimaryKey) (*models.Sales, error)
	GetList(ctx context.Context, req *models.SalesGetListRequest) (*models.SalesGetListResponse, error)
	Delete(ctx context.Context, req *models.SalesPrimaryKey) error
}

type TransactionRepoInterface interface {
	Create(ctx context.Context, req *models.CreateTransaction) (string, error)
	Update(ctx context.Context, req *models.UpdateTransaction) (int64, error)
	GetById(ctx context.Context, req *models.TransactionPrimaryKey) (*models.Transaction, error)
	GetList(ctx context.Context, req *models.TransactionGetListRequest) (*models.TransactionGetListResponse, error)
	Delete(ctx context.Context, req *models.TransactionPrimaryKey) error
}
