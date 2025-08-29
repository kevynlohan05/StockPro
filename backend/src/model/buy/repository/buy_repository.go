package repository

import (
	"database/sql"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	buyModel "github.com/kevynlohan05/StockPro/src/model/buy"
)

func NewBuyRepository(db *sql.DB) BuyRepository {
	return &buyRepository{db}
}

type buyRepository struct {
	db *sql.DB
}

type BuyRepository interface {
	CreateBuyTransaction(buy buyModel.BuyDomainInterface) *rest_err.RestErr
}
