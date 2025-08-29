package repository

import (
	"log"
	"strconv"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	buyModel "github.com/kevynlohan05/StockPro/src/model/buy"
)

func (r *buyRepository) CreateBuyTransaction(buy buyModel.BuyDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando CreateBuyTransaction")

	tx, err := r.db.Begin()
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao iniciar transação")
	}
	defer tx.Rollback()

	res, err := tx.Exec(`INSERT INTO buys (user_id, total_value, payment_method) VALUES (?, ?, ?)`,
		buy.GetUserID(), buy.GetTotal(), buy.GetPaymentMethod())
	if err != nil {
		log.Println("Erro inserir buys:", err)
		return rest_err.NewInternalServerError("Erro ao inserir compra")
	}

	buyID, _ := res.LastInsertId()
	buy.SetID(strconv.FormatInt(buyID, 10))

	for _, item := range buy.GetItems() {
		_, err := tx.Exec(`INSERT INTO buys_items (buy_id, product_id, quantity, price) VALUES (?, ?, ?, ?)`,
			buyID, item.GetProductID(), item.GetQuantity(), item.GetPrice())
		if err != nil {
			log.Println("Erro inserir buys_items:", err)
			return rest_err.NewInternalServerError("Erro ao inserir item da compra")
		}

		_, err = tx.Exec(`UPDATE stock SET quantity = quantity + ? WHERE product_id = ?`, item.GetQuantity(), item.GetProductID())
		if err != nil {
			log.Println("Erro atualizar stock:", err)
			return rest_err.NewInternalServerError("Erro ao atualizar estoque")
		}

		_, err = tx.Exec(`INSERT INTO movements_stock (product_id, user_id, type, quantity, reason) VALUES (?,?, 'entrada', ?, ?)`,
			item.GetProductID(), buy.GetUserID(), item.GetQuantity(), "Compra ID "+strconv.FormatInt(buyID, 10))
		if err != nil {
			log.Println("Erro movements_stock:", err)
			return rest_err.NewInternalServerError("Erro ao registrar movimentação de estoque")
		}
	}

	_, err = tx.Exec(`INSERT INTO cash (transaction_id, user_id, type, value, description) VALUES (?, ?, 'saida', ?, ?)`,
		buyID, buy.GetUserID(), buy.GetTotal(), "Buy")
	if err != nil {
		log.Println("Erro inserir cash:", err)
		return rest_err.NewInternalServerError("Erro ao registrar saída no caixa")
	}

	_, err = tx.Exec(`
		INSERT INTO cash_balance (payment, value) 
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE value = value - VALUES(value)
		`, buy.GetPaymentMethod(), buy.GetTotal())
	if err != nil {
		log.Println("Erro atualizar cash_balance:", err)
		return rest_err.NewInternalServerError("Erro ao atualizar saldo do caixa")
	}

	if err := tx.Commit(); err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar transação")
	}

	return nil
}
