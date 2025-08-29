package view

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/controller/model/response"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func ConvertUserDomainToResponse(userDomain userModel.UserDomainInterface) response.UserResponse {

	log.Println("Convertendo userDomain para userResponse")

	return response.UserResponse{
		ID:     userDomain.GetID(),
		Name:   userDomain.GetName(),
		Email:  userDomain.GetEmail(),
		Sector: userDomain.GetSector(),
		Role:   userDomain.GetRole(),
		Active: userDomain.GetIsActive(),
	}
}

func ConvertProductDomainToResponse(productDomain productModel.ProductDomainInterface) response.ProductResponse {

	log.Println("Convertendo userDomain para userResponse")

	return response.ProductResponse{
		ID:            productDomain.GetID(),
		Name:          productDomain.GetName(),
		Description:   productDomain.GetDescription(),
		Mark:          productDomain.GetMark(),
		PurchasePrice: productDomain.GetPurchasePrice(),
		SalePrice:     productDomain.GetSalePrice(),
		Images:        productDomain.GetImages(),
	}
}

func ConvertMovementStockDomainToResponse(movementDomain stockModel.StockDomainInterface) response.StockMovementResponse {

	log.Println("Convertendo movementDomain para movementResponse")

	return response.StockMovementResponse{
		ID:        movementDomain.GetID(),
		ProductID: movementDomain.GetProductID(),
		UserID:    movementDomain.GetUserID(),
		Type:      movementDomain.GetStockType(),
		Quantity:  movementDomain.GetQuantity(),
		Reason:    movementDomain.GetReason(),
		CreatedAt: movementDomain.GetCreatedAt(),
	}
}
