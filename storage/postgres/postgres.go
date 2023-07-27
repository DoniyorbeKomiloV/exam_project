package postgres

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db          *pgxpool.Pool
	branch      *BranchRepo
	tarif       *TarifRepo
	staff       *StaffRepo
	sales       *SalesRepo
	transaction *TransactionRepo
}

func (s *store) Branch() storage.BranchRepoInterface {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *store) StaffTarif() storage.StaffTarifRepoInterface {
	if s.tarif == nil {
		s.tarif = NewTarifRepo(s.db)
	}
	return s.tarif
}

func (s *store) Staff() storage.StaffRepoInterface {
	if s.staff == nil {
		s.staff = NewStaffRepo(s.db)
	}
	return s.staff
}

func (s *store) Sales() storage.SalesRepoInterface {
	if s.sales == nil {
		s.sales = NewSalesRepo(s.db)
	}
	return s.sales
}

func (s *store) StaffTransaction() storage.TransactionRepoInterface {
	if s.transaction == nil {
		s.transaction = NewTransactionRepo(s.db)
	}
	return s.transaction
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageInterface, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}
