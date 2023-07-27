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

type StaffRepo struct {
	db *pgxpool.Pool
}

func (s StaffRepo) Create(ctx context.Context, req *models.CreateStaff) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO staff(id, branch_id, tarif_id, type, name, updated_at, deleted) VALUES ($1, $2, $3, $4, $5, NOW(), false)`

	_, err := s.db.Exec(ctx, query, id, req.BranchId, req.TarifId, req.Type, req.Name)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s StaffRepo) Update(ctx context.Context, req *models.UpdateStaff) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE staff 
		SET branch_id = :branch_id,
		    tarif_id = :tarif_id,
		    type = :type,
		    name = :name,
		    balance = :balance,
		    updated_at = NOW() WHERE id = :id`

	params = map[string]interface{}{
		"id":        req.Id,
		"branch_id": req.BranchId,
		"tarif_id":  req.TarifId,
		"type":      req.Type,
		"name":      req.Name,
		"balance":   req.Balance,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s StaffRepo) GetById(ctx context.Context, req *models.StaffPrimaryKey) (*models.Staff, error) {
	var (
		id       sql.NullString
		branchId sql.NullString
		tarifId  sql.NullString
		typee    sql.NullString
		name     sql.NullString
		balance  float64
	)

	query := `SELECT id, branch_id, tarif_id, type, name, balance FROM staff WHERE id = $1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(&id, &branchId, &tarifId, &typee, &name, &balance)

	if err != nil {
		return nil, err
	}

	return &models.Staff{
		Id:       id.String,
		BranchId: branchId.String,
		TarifId:  tarifId.String,
		Type:     typee.String,
		Name:     name.String,
		Balance:  balance,
	}, nil
}

func (s StaffRepo) GetList(ctx context.Context, req *models.StaffGetListRequest) (*models.StaffGetListResponse, error) {
	var (
		resp   = &models.StaffGetListResponse{}
		query  string
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		order  = " ORDER BY created_at DESC "
	)

	query = `
		SELECT COUNT(*) OVER(), id, branch_id, tarif_id, type, name, balance FROM staff`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + order + offset + limit

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			branchId sql.NullString
			tarifId  sql.NullString
			typee    sql.NullString
			name     sql.NullString
			balance  float64
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&branchId,
			&tarifId,
			&typee,
			&name,
			&balance,
		)

		if err != nil {
			return nil, err
		}

		resp.Staffs = append(resp.Staffs, &models.Staff{
			Id:       id.String,
			BranchId: branchId.String,
			TarifId:  tarifId.String,
			Type:     typee.String,
			Name:     name.String,
			Balance:  balance,
		})
	}

	return resp, nil
}

func (s StaffRepo) Delete(ctx context.Context, req *models.StaffPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE staff SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewStaffRepo(db *pgxpool.Pool) *StaffRepo {
	return &StaffRepo{
		db: db,
	}
}
