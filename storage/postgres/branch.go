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

type BranchRepo struct {
	db *pgxpool.Pool
}

func (b BranchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO branch(id, name, address, phone_number, updated_at, deleted) VALUES ($1, $2, $3, $4, NOW(), false)`

	_, err := b.db.Exec(ctx, query, id, req.Name, req.Address, req.PhoneNumber)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (b BranchRepo) Update(ctx context.Context, req *models.UpdateBranch) (int64, error) {
	var params map[string]interface{}
	query := `UPDATE branch SET name = :name, address = :address, phone_number = :phone_number, updated_at = NOW() WHERE id = :id`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"address":      req.Address,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := b.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (b BranchRepo) GetById(ctx context.Context, req *models.BranchPrimaryKey) (*models.Branch, error) {
	var (
		id          sql.NullString
		name        sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
	)

	query := `SELECT id, name, address, phone_number FROM branch WHERE id = $1`

	err := b.db.QueryRow(ctx, query, req.Id).Scan(&id, &name, &address, &phoneNumber)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:          id.String,
		Name:        name.String,
		Address:     address.String,
		PhoneNumber: phoneNumber.String,
	}, nil
}

func (b BranchRepo) GetList(ctx context.Context, req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {
	var (
		resp   = &models.BranchGetListResponse{}
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		order  = " ORDER BY created_at DESC "
	)
	query :=
		`SELECT COUNT(*) OVER(), id, name, address, phone_number FROM branch`
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.SearchName != "" {
		where += ` AND name ILIKE '%' || '` + req.SearchName + `' || '%'`
	}

	if req.SearchAddress != "" {
		where += ` AND address ILIKE '%' || '` + req.SearchAddress + `' || '%'`
	}

	query += where + order + offset + limit
	fmt.Println(query)
	rows, err := b.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			name        sql.NullString
			address     sql.NullString
			phoneNumber sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&phoneNumber,
		)
		if err != nil {
			return nil, err
		}
		resp.Branches = append(resp.Branches, &models.Branch{Id: id.String, Name: name.String, Address: address.String, PhoneNumber: phoneNumber.String})
	}

	return resp, nil
}

func (b BranchRepo) Delete(ctx context.Context, req *models.BranchPrimaryKey) error {
	_, err := b.db.Exec(ctx, "UPDATE branch SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}
