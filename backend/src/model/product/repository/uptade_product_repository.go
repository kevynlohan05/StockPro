package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
)

func (pr *productRepository) UpdateProduct(productId string, productDomain productModel.ProductDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando UpdateProduct repository!")

	var currentProduct struct {
		Name          string
		Description   string
		Mark          string
		PurchasePrice float64
		SalePrice     float64
		Image         string
	}

	queryFind := `
		SELECT name, description, mark, purchase_price, sale_price, image 
		FROM products WHERE id = ?
	`
	err := pr.db.QueryRow(queryFind, productId).Scan(
		&currentProduct.Name,
		&currentProduct.Description,
		&currentProduct.Mark,
		&currentProduct.PurchasePrice,
		&currentProduct.SalePrice,
		&currentProduct.Image,
	)
	if err != nil {
		log.Println("Erro ao buscar produto existente: ", err)
		return rest_err.NewInternalServerError("Erro ao buscar produto existente")
	}

	log.Println(&currentProduct.Image)

	value := converter.ConvertProductDomainToEntity(productDomain)

	if value.Name == "" {
		value.Name = currentProduct.Name
	}
	if value.Description == "" {
		value.Description = currentProduct.Description
	}
	if value.Mark == "" {
		value.Mark = currentProduct.Mark
	}
	if value.PurchasePrice == 0 {
		value.PurchasePrice = currentProduct.PurchasePrice
	}
	if value.SalePrice == 0 {
		value.SalePrice = currentProduct.SalePrice
	}
	if value.Images == "" {
		value.Images = currentProduct.Image
	}

	queryUpdate := `
		UPDATE products SET
			name = ?,
			description = ?,
			mark = ?,
			purchase_price = ?,
			sale_price = ?,
			image = ?
		WHERE id = ?
	`
	result, err := pr.db.Exec(queryUpdate,
		value.Name,
		value.Description,
		value.Mark,
		value.PurchasePrice,
		value.SalePrice,
		value.Images,
		productId)

	if err != nil {
		log.Println("Erro ao executar query de atualização de produto: ", err)
		return rest_err.NewInternalServerError("Erro ao atualizar produto")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas: ", err)
		return rest_err.NewInternalServerError("Erro ao atualizar produto")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Produto não encontrado")
	}

	log.Println("Produto atualizado com sucesso!")
	return nil
}
