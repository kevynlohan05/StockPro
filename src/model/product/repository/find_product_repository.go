package repository

import (
	"database/sql"
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	"github.com/kevynlohan05/StockPro/src/model/entity"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pr *productRepository) FindProductById(productId string) (productModel.ProductDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindProductById repository")

	query := `
		SELECT id, name, description, mark, purchase_price, sale_price, image 
		FROM products WHERE id = ?
	`

	row := pr.db.QueryRow(query, productId)

	var entity entity.ProductEntity

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Description,
		&entity.Mark,
		&entity.PurchasePrice,
		&entity.SalePrice,
		&entity.Images,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Produto não encontrado")
		}
		log.Println("Erro ao escanear resultado da query: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar produto")
	}

	return converter.ConvertProductEntityToDomain(entity), nil
}
