package repository

import (
	"database/sql"
)

func NewBuyRepository(db *sql.DB) BuyRepository {
	return &buyRepository{db}
}

type buyRepository struct {
	db *sql.DB
}

type BuyRepository interface {
}
