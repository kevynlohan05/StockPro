package service

import (
	"log"
	"strconv"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	buyModel "github.com/kevynlohan05/StockPro/src/model/buy"
)

func (bs *buyDomainService) CreateBuyService(buy buyModel.BuyDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando CreateBuyService")

	for _, item := range buy.GetItems() {
		ok, err := bs.repositoryStock.CheckStock(item.GetProductID(), item.GetQuantity())
		if err != nil {
			return err
		}
		if !ok {
			return rest_err.NewBadRequestError("Estoque insuficiente para produto ID " + strconv.Itoa(item.GetProductID()))
		}
	}

	if err := bs.repositoryBuy.CreateBuyTransaction(buy); err != nil {
		return err
	}

	return nil
}
