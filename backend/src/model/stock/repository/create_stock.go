package repository

import (
	"database/sql"
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func (r *stockRepository) CreateStockMovement(stockDomain stockModel.StockDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando CreateStockMovement repository")

	value := converter.ConvertStockDomainToEntity(stockDomain)
	log.Println(value)
	var queryInsert string
	if value.Type == "entrada" {
		queryInsert = `INSERT INTO stock (quantity, product_id) VALUES (?, ?)
		         ON DUPLICATE KEY UPDATE quantity = quantity + VALUES(quantity)`
	} else if value.Type == "saida" {
		queryInsert = `UPDATE stock SET quantity = quantity - ? WHERE product_id = ?`
	}

	_, err := r.db.Exec(queryInsert,
		value.Quantity,
		value.ProductID,
	)
	if err != nil {
		log.Println("Erro ao inserir movimentação de estoque no banco de dados: ", err)
		return rest_err.NewInternalServerError("Erro ao inserir movimentação de estoque no banco de dados")
	}

	queryInsertMovement := `INSERT INTO movements_stock (product_id, user_id, type, quantity, reason) VALUES (?, ?, ?, ?, ?)`
	_, err = r.db.Exec(
		queryInsertMovement,
		value.ProductID,
		value.UserID,
		value.Type,
		value.Quantity,
		value.Reason,
	)
	if err != nil {
		log.Println("Erro ao registrar movimentação de estoque no banco de dados: ", err)
		return rest_err.NewInternalServerError("Erro ao registrar movimentação de estoque no banco de dados")
	}

	log.Println("Movimentação de estoque criada com sucesso!")
	return nil
}

func (r *stockRepository) CheckStock(productID, quantity int) (bool, *rest_err.RestErr) {
	var current int
	err := r.db.QueryRow("SELECT quantity FROM stock WHERE product_id = ?", productID).Scan(&current)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, rest_err.NewBadRequestError("Produto sem estoque")
		}
		return false, rest_err.NewInternalServerError("Erro ao verificar estoque")
	}
	return current >= quantity, nil
}
