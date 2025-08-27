package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
)

func (pr *productRepository) DeleteProduct(productId string) *rest_err.RestErr {
	log.Println("Iniciando DeleteProduct repository!")

	query := `
		DELETE FROM products WHERE id = ?
	`
	result, err := pr.db.Exec(query, productId)
	if err != nil {
		log.Println("Erro ao executar query de deleção de produto: ", err)
		return rest_err.NewInternalServerError("Erro ao deletar produto")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas: ", err)
		return rest_err.NewInternalServerError("Erro ao deletar produto")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Produto não encontrado")
	}

	log.Println("Produto deletado com sucesso!")
	return nil
}
