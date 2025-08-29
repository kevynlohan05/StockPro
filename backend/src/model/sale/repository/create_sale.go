package repository

import (
	"log"
	"strconv"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	saleModel "github.com/kevynlohan05/StockPro/src/model/sale"
)

func (r *saleRepository) CreateSaleTransaction(sale saleModel.SaleDomainInterface) *rest_err.RestErr {
	tx, err := r.db.Begin()
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao iniciar transação")
	}
	defer tx.Rollback()

	res, err := tx.Exec(`INSERT INTO sales (user_id, total_value, payment_method) VALUES (?, ?, ?)`,
		sale.GetUserID(), sale.GetTotal(), sale.GetPaymentMethod())
	if err != nil {
		log.Println("Erro inserir sales:", err)
		return rest_err.NewInternalServerError("Erro ao inserir venda")
	}

	saleID, _ := res.LastInsertId()
	sale.SetID(strconv.FormatInt(saleID, 10))

	for _, item := range sale.GetItems() {
		_, err := tx.Exec(`INSERT INTO sales_items (sale_id, product_id, quantity, price) VALUES (?, ?, ?, ?)`,
			saleID, item.GetProductID(), item.GetQuantity(), item.GetPrice())
		if err != nil {
			log.Println("Erro inserir sales_items:", err)
			return rest_err.NewInternalServerError("Erro ao inserir item da venda")
		}

		_, err = tx.Exec(`UPDATE stock SET quantity = quantity - ? WHERE product_id = ?`, item.GetQuantity(), item.GetProductID())
		if err != nil {
			log.Println("Erro atualizar stock:", err)
			return rest_err.NewInternalServerError("Erro ao atualizar estoque")
		}

		_, err = tx.Exec(`INSERT INTO movements_stock (product_id, user_id, type, quantity, reason) VALUES (?,?, 'saida', ?, ?)`,

			item.GetProductID(), sale.GetUserID(), item.GetQuantity(), "Venda ID "+strconv.FormatInt(saleID, 10))
		if err != nil {
			log.Println("Erro movements_stock:", err)
			return rest_err.NewInternalServerError("Erro ao registrar movimentação de estoque")
		}
	}

	_, err = tx.Exec(`INSERT INTO cash (sale_id, user_id, type, value, description) VALUES (?, ?, 'entrada', ?, ?)`,
		saleID, sale.GetUserID(), sale.GetTotal(), "Venda ID "+strconv.FormatInt(saleID, 10))
	if err != nil {
		log.Println("Erro inserir cash:", err)
		return rest_err.NewInternalServerError("Erro ao registrar entrada no caixa")
	}

	_, err = tx.Exec(`
		INSERT INTO cash_balance (payment, value) 
		VALUES (?, ?) 
		ON DUPLICATE KEY UPDATE value = value + VALUES(value)
	`, sale.GetPaymentMethod(), sale.GetTotal())
	if err != nil {
		log.Println("Erro atualizar cash_balance:", err)
		return rest_err.NewInternalServerError("Erro ao atualizar saldo do caixa")
	}

	if err := tx.Commit(); err != nil {
		return rest_err.NewInternalServerError("Erro ao salvar transação")
	}

	return nil
}
