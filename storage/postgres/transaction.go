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

type TransactionRepo struct {
	db *pgxpool.Pool
}

func (t TransactionRepo) Create(ctx context.Context, req *models.CreateTransaction) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO staff_transaction(id, sales_id, type, source_type, text, amount, staff_id, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	_, err := t.db.Exec(ctx, query, id, req.SalesId, req.Type, req.SourceType, req.Text, req.Amount, req.StaffId)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TransactionRepo) Update(ctx context.Context, req *models.UpdateTransaction) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE staff_transaction 
		SET id = :id,
			sales_id = :sales_id,
			type = :type,
			source_type = :source_type,
			text = :text,
			amount = :amount,
			staff_id = :staff_id,
		    updated_at = NOW() 
		WHERE id = :id`

	params = map[string]interface{}{
		"id":          req.Id,
		"sales_id":    req.SalesId,
		"type":        req.Type,
		"source_type": req.SourceType,
		"text":        req.Text,
		"amount":      req.Amount,
		"staff_id":    req.StaffId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := t.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (t TransactionRepo) GetById(ctx context.Context, req *models.TransactionPrimaryKey) (*models.Transaction, error) {
	var (
		id         sql.NullString
		salesId    sql.NullString
		typee      sql.NullString
		sourceType sql.NullString
		text       sql.NullString
		amount     float64
		staffId    sql.NullString
		createdAt  sql.NullString
	)

	query := `SELECT sales_id, type, source_type, text, amount, staff_id, created_at, updated_at FROM staff_transaction	WHERE id = $1 ORDER BY created_at DESC`

	err := t.db.QueryRow(ctx, query, req.Id).Scan(&id, &salesId, &typee, &sourceType, &text, &amount, &staffId, &createdAt)

	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		Id:         id.String,
		SalesId:    salesId.String,
		Type:       typee.String,
		SourceType: sourceType.String,
		Text:       text.String,
		Amount:     amount,
		StaffId:    staffId.String,
	}, nil
}

func (t TransactionRepo) GetList(ctx context.Context, req *models.TransactionGetListRequest) (*models.TransactionGetListResponse, error) {
	var (
		resp   = &models.TransactionGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(), id, sales_id, type, source_type, text, amount, staff_id, created_at, updated_at
		FROM staff_transaction 
		WHERE deleted = false 
		ORDER BY created_at DESC 
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND source_type ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id         sql.NullString
			salesId    sql.NullString
			typee      sql.NullString
			sourceType sql.NullString
			text       sql.NullString
			amount     float64
			staffId    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&salesId,
			&typee,
			&sourceType,
			&text,
			&amount,
			&staffId,
		)

		if err != nil {
			return nil, err
		}

		resp.Transactions = append(resp.Transactions, &models.Transaction{
			Id:         id.String,
			SalesId:    salesId.String,
			Type:       typee.String,
			SourceType: sourceType.String,
			Text:       text.String,
			Amount:     amount,
			StaffId:    staffId.String,
		})
	}

	return resp, nil
}

func (t TransactionRepo) Delete(ctx context.Context, req *models.TransactionPrimaryKey) error {
	_, err := t.db.Exec(ctx, "UPDATE staff_transaction SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}
