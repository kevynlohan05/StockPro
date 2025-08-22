package repository

import (
	"database/sql"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func NewUserRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db,
	}
}

type productRepository struct {
	db *sql.DB
}

type ProductRepository interface {
	CreateProduct(productDomain productModel.ProductDomainInterface) (productModel.ProductDomainInterface, *rest_err.RestErr)
	UpdateProduct(productId string, productDomain productModel.ProductDomainInterface) *rest_err.RestErr
	FindProductById(productId string) (productModel.ProductDomainInterface, *rest_err.RestErr)
	DeleteProduct(productId string) *rest_err.RestErr
}
