package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pr *productRepository) CreateProduct(productDomain productModel.ProductDomainInterface) (productModel.ProductDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando CreateProduct repository!")

	value := converter.ConvertProductDomainToEntity(productDomain)

	if value == nil {
		log.Println("Error: conversion to entity failed")
		return nil, rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	query := `
		INSERT INTO products 
		(name, description, mark, purchase_price, sale_price, image)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := pr.db.Exec(
		query,
		value.Name,
		value.Description,
		value.Mark,
		value.PurchasePrice,
		value.SalePrice,
		value.Images,
	)

	if err != nil {
		log.Println("Erro ao inserir usuário no banco de dados: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao inserir usuário no banco de dados")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Erro ao obter ID do usuário inserido: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao obter ID do usuário inserido")
	}

	value.ID = int(id)

	log.Println("Produto inserido com sucesso no banco de dados.")
	return converter.ConvertProductEntityToDomain(*value), nil
}
