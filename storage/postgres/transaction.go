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
	var (
		cashier   string
		assistant string
		tarifId   string
		balance   float64
		typ       string
		cash      float64
		card      float64
		tarifId2  string
		balance2  float64
		typ2      string
		cash2     float64
		card2     float64
	)
	defer t.db.Close()
	query := `SELECT cashier_id, shop_assistant_id FROM sales WHERE id = $1`
	err := t.db.QueryRow(ctx, query, req.StaffId).Scan(&cashier, &assistant)
	if assistant == "" {
		query2 := `SELECT tarif_id, balance FROM staff WHERE id = $1`
		err = t.db.QueryRow(ctx, query2, cashier).Scan(&tarifId, &balance)
		if err != nil {
			return "", err
		}
		query3 := `SELECT type, cash, card FROM staff_tarif WHERE id = $1`
		err = t.db.QueryRow(ctx, query3, cashier).Scan(&typ, &cash, &card)
		if err != nil {
			return "", err
		}
		if typ == "fixed" {
			if req.Type == "cash" {
				balance += cash
			} else {
				balance += card
			}
		} else {
			if req.Type == "cash" {
				balance += req.Amount * cash
			} else {
				balance += req.Amount * card
			}
		}

		query4 := `UPDATE staff SET balance = :balance, updated_at = NOW() WHERE id = :cashier`
		_, err := t.db.Exec(ctx, query4)
		if err != nil {
			return "", err
		}
	} else {
		query2 := `SELECT tarif_id, balance FROM staff WHERE id = $1`
		err = t.db.QueryRow(ctx, query2, cashier).Scan(&tarifId, &balance)
		if err != nil {
			return "", err
		}
		err = t.db.QueryRow(ctx, query2, assistant).Scan(&tarifId2, &balance2)
		if err != nil {
			return "", err
		}
		query3 := `SELECT type, cash, card FROM staff_tarif WHERE id = $1`
		err = t.db.QueryRow(ctx, query3, cashier).Scan(&typ, &cash, &card)
		err = t.db.QueryRow(ctx, query3, assistant).Scan(&typ2, &cash2, &card2)
		if err != nil {
			return "", err
		}
		if typ == "fixed" {
			if req.Type == "cash" {
				balance += cash
			} else {
				balance += card
			}
		} else {
			if req.Type == "cash" {
				balance += req.Amount * cash
			} else {
				balance += req.Amount * card
			}
		}
		if typ2 == "fixed" {
			if req.Type == "cash" {
				balance2 += cash2
			} else {
				balance2 += card2
			}
		} else {
			if req.Type == "cash" {
				balance2 += req.Amount * cash2
			} else {
				balance2 += req.Amount * card2
			}
		}

		query4 := `UPDATE staff SET balance = $1, updated_at = NOW() WHERE id = $2`
		_, err := t.db.Exec(ctx, query4, balance, cashier)
		if err != nil {
			return "", err
		}
		_, err = t.db.Exec(ctx, query4, balance2, assistant)
		if err != nil {
			return "", err
		}

	}
	queryInsert := `INSERT INTO staff_transaction(id, sales_id, type, source_type, text, amount, staff_id, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	_, err = t.db.Exec(ctx, queryInsert, id, req.SalesId, req.Type, req.SourceType, req.Text, balance, req.StaffId)
	if err != nil {
		return "", err
	}
	_, err = t.db.Exec(ctx, queryInsert, id, req.SalesId, req.Type, req.SourceType, req.Text, balance2, req.StaffId)
	if err != nil {
		return "", err
	}
	query = `SELECT id, branch_id, tarif_id, type, name, balance, created_at, updated_at, deleted_at, deleted FROM staff WHERE id = $1`

	err = t.db.QueryRow(ctx, query, req.StaffId).Scan()

	return id, nil

}

func (t TransactionRepo) Update(ctx context.Context, req *models.UpdateTransaction) (int64, error) {
	var params map[string]interface{}
	defer t.db.Close()
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
	defer t.db.Close()
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
	defer t.db.Close()
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
	defer t.db.Close()
	_, err := t.db.Exec(ctx, "UPDATE staff_transaction SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}
