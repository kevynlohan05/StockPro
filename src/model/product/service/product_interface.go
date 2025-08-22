package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	repositoryProduct "github.com/kevynlohan05/StockPro/src/model/product/repository"
)

func NewProductDomainService(productRepository repositoryProduct.ProductRepository) ProductDomainService {
	return &productDomainService{productRepository}
}

type productDomainService struct {
	productRepository repositoryProduct.ProductRepository
}

type ProductDomainService interface {
	CreateProductServices(productModel.ProductDomainInterface) (productModel.ProductDomainInterface, *rest_err.RestErr)
	UpdateProduct(productId string, productDomain productModel.ProductDomainInterface) *rest_err.RestErr
	FindProductByIdServices(id string) (productModel.ProductDomainInterface, *rest_err.RestErr)
	DeleteProductServices(productId string) *rest_err.RestErr
}
