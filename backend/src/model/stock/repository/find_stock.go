package repository

import (
	"log"
	"time"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	"github.com/kevynlohan05/StockPro/src/model/entity"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
)

func (sr *stockRepository) FindMovementsStock() ([]stockModel.StockDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindMovementsStock repository!")

	query := `
	SELECT id, product_id, user_id, type, quantity, reason, created_at
	FROM movements_stock
	`
	rows, err := sr.db.Query(query)
	if err != nil {
		log.Println("Erro ao executar query: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar Movements Stock")
	}
	defer rows.Close()

	var movements_stock []stockModel.StockDomainInterface

	for rows.Next() {
		var entity entity.StockEntity
		var createdAt time.Time

		err := rows.Scan(
			&entity.ID,
			&entity.ProductID,
			&entity.UserID,
			&entity.Type,
			&entity.Quantity,
			&entity.Reason,
			&createdAt,
		)

		if err != nil {
			log.Println("Erro ao escanear resultado da query: ", err)
			return nil, rest_err.NewInternalServerError("Erro ao buscar movements stock")
		}

		entity.CreatedAt = int(createdAt.Unix())

		movements_stock = append(movements_stock, converter.ConvertStockEntityToDomain(entity))
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro após iterar sobre linhas: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar movements_stock")
	}

	return movements_stock, nil
}

func (sr *stockRepository) FindStockQuantity() ([]stockModel.StockQuantityDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindStockQuantity respository!")

	query := `
		SELECT product_id, quantity
		FROM stock
	`

	rows, err := sr.db.Query(query)
	if err != nil {
		log.Println("Erro ao executar query: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar stock")
	}

	defer rows.Close()

	var stock []stockModel.StockQuantityDomainInterface

	for rows.Next() {
		var entity entity.StockQuantityEntity

		err := rows.Scan(
			&entity.ProductID,
			&entity.Quantity,
		)

		if err != nil {
			log.Println("Erro ao escanear resultado da query: ", err)
			return nil, rest_err.NewInternalServerError("Erro ao buscar stock")
		}

		stock = append(stock, converter.ConvertStockQuantityEntityToDomain(entity))
	}

	if err = rows.Err(); err != nil {
		log.Println("Erro após iterar sobre linhas: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar stock")
	}

	return stock, nil
}
