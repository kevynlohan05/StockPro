package repository

import (
	"database/sql"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	saleModel "github.com/kevynlohan05/StockPro/src/model/sale"
)

func NewSaleRepository(db *sql.DB) SaleRepository {
	return &saleRepository{db}
}

type saleRepository struct {
	db *sql.DB
}

type SaleRepository interface {
	CreateSaleTransaction(sale saleModel.SaleDomainInterface) *rest_err.RestErr
}
