package repository

import (
	"database/sql"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func NewStockRepository(db *sql.DB) StockRepository {
	return &stockRepository{db}
}

type stockRepository struct {
	db *sql.DB
}

type StockRepository interface {
	CreateStockMovement(stockDomain stockModel.StockDomainInterface) *rest_err.RestErr
	CheckStock(productID, quantity int) (bool, *rest_err.RestErr)
	FindMovementsStock() ([]stockModel.StockDomainInterface, *rest_err.RestErr)
	FindStockQuantity() ([]stockModel.StockQuantityDomainInterface, *rest_err.RestErr)
}
