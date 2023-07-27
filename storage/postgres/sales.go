package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SalesRepo struct {
	db *pgxpool.Pool
}

func (s SalesRepo) Create(ctx context.Context, req *models.CreateSales) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO sales(id, branch_id, shop_assistant_id, cashier_id, price, payment_type, status, client_name, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())`

	_, err := s.db.Exec(ctx, query, id, req.BranchId, req.ShopAssistantId, req.CashierId, req.Price, req.PaymentType, req.Status, req.ClientName)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s SalesRepo) Update(ctx context.Context, req *models.UpdateSales) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE sales 
		SET id = :id, 
		    branch_id = :branch_id,
			shop_assistant_id = :shop_assistant_id,
			cashier_id = :cashier_id,
			price = :price,
			payment_type = :payment_type,
			status = :status,
			client_name = :client_name,
			updated_at = NOW()
		WHERE id = :id`

	params = map[string]interface{}{
		"id":                req.Id,
		"branch_id":         req.BranchId,
		"shop_assistant_id": req.ShopAssistantId,
		"cashier_id":        req.CashierId,
		"price":             req.Price,
		"payment_type":      req.PaymentType,
		"status":            req.Status,
		"client_name":       req.ClientName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s SalesRepo) GetById(ctx context.Context, req *models.SalesPrimaryKey) (*models.Sales, error) {
	var (
		id              sql.NullString
		branchId        sql.NullString
		shopAssistantId sql.NullString
		cashierId       sql.NullString
		price           float64
		paymentType     sql.NullString
		status          sql.NullString
		clientName      sql.NullString
	)

	query := `SELECT id, branch_id, shop_assistant_id, cashier_id, price, payment_type, status, client_name FROM sales WHERE id = $1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&branchId,
		&shopAssistantId,
		&cashierId,
		&price,
		&paymentType,
		&status,
		&clientName)

	if err != nil {
		return nil, err
	}

	return &models.Sales{
		Id:              id.String,
		BranchId:        branchId.String,
		ShopAssistantId: shopAssistantId.String,
		CashierId:       cashierId.String,
		Price:           price,
		PaymentType:     paymentType.String,
		Status:          status.String,
		ClientName:      clientName.String,
	}, nil
}

func (s SalesRepo) GetList(ctx context.Context, req *models.SalesGetListRequest) (*models.SalesGetListResponse, error) {
	var (
		resp   = &models.SalesGetListResponse{}
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), id, branch_id, shop_assistant_id, cashier_id, price, payment_type, status, client_name FROM sales`
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + order + offset + limit

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id              sql.NullString
			branchId        sql.NullString
			shopAssistantId sql.NullString
			cashierId       sql.NullString
			price           float64
			paymentType     sql.NullString
			status          sql.NullString
			clientName      sql.NullString
		)
		err := rows.Scan(
			&id,
			&branchId,
			&shopAssistantId,
			&cashierId,
			&price,
			&paymentType,
			&status,
			&clientName,
		)
		if err != nil {
			return nil, err
		}
		resp.Sales = append(
			resp.Sales,
			&models.Sales{
				Id:              id.String,
				BranchId:        branchId.String,
				ShopAssistantId: shopAssistantId.String,
				CashierId:       cashierId.String,
				Price:           price,
				PaymentType:     paymentType.String,
				Status:          status.String,
				ClientName:      clientName.String,
			})
	}

	return resp, nil
}

func (s SalesRepo) Delete(ctx context.Context, req *models.SalesPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE sales SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewSalesRepo(db *pgxpool.Pool) *SalesRepo {
	return &SalesRepo{
		db: db,
	}
}
