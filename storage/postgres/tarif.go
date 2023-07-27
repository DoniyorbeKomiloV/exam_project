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

type TarifRepo struct {
	db *pgxpool.Pool
}

func (t TarifRepo) Create(ctx context.Context, req *models.CreateTarif) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO staff_tarif(id, name, type, updated_at) VALUES ($1, $2, $3, NOW())`

	_, err := t.db.Exec(ctx, query, id, req.Name, req.Type)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TarifRepo) Update(ctx context.Context, req *models.UpdateTarif) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE staff_tarif 
		SET name = :name,
		    type = :type,
		    updated_at = NOW() 
		WHERE id = :id`

	params = map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
		"type": req.Type,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := t.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (t TarifRepo) GetById(ctx context.Context, req *models.TarifPrimaryKey) (*models.Tarif, error) {
	var (
		id    sql.NullString
		name  sql.NullString
		typee sql.NullString
	)

	query := `SELECT id, name, type FROM staff_tarif WHERE id = $1`

	err := t.db.QueryRow(ctx, query, req.Id).Scan(&id, &name, &typee)

	if err != nil {
		return nil, err
	}

	return &models.Tarif{
		Id:   id.String,
		Name: name.String,
		Type: typee.String,
	}, nil
}

func (t TarifRepo) GetList(ctx context.Context, req *models.TarifGetListRequest) (*models.TarifGetListResponse, error) {
	var (
		resp   = &models.TarifGetListResponse{}
		query  string
		where  = " WHERE deleted = false "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		order  = " ORDER BY created_at DESC "
	)

	query = `SELECT COUNT(*) OVER(), id, name, type	FROM staff_tarif`

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

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id    sql.NullString
			name  sql.NullString
			typee sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&typee,
		)

		if err != nil {
			return nil, err
		}

		resp.Tarifs = append(resp.Tarifs, &models.Tarif{
			Id:   id.String,
			Name: name.String,
			Type: typee.String,
		})
	}

	return resp, nil
}

func (t TarifRepo) Delete(ctx context.Context, req *models.TarifPrimaryKey) error {
	_, err := t.db.Exec(ctx, "UPDATE staff_tarif SET deleted = true, deleted_at = NOW() WHERE id = $1", req.Id)
	return err
}

func NewTarifRepo(db *pgxpool.Pool) *TarifRepo {
	return &TarifRepo{
		db: db,
	}
}
