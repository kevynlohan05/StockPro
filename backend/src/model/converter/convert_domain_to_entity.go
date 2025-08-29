package converter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kevynlohan05/StockPro/src/model/entity"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func ConvertUserDomainToEntity(userDomain userModel.UserDomainInterface) *entity.UserEntity {
	entity := &entity.UserEntity{
		Name:     userDomain.GetName(),
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Sector:   userDomain.GetSector(),
		Role:     userDomain.GetRole(),
		IsActive: userDomain.GetIsActive(),
	}

	if userDomain.GetID() != "" {
		var id int
		if _, err := fmt.Sscanf(userDomain.GetID(), "%d", &id); err == nil {
			entity.ID = id
		}
	}

	return entity
}

func ConvertProductDomainToEntity(productDomain productModel.ProductDomainInterface) *entity.ProductEntity {
	var imagesJSON []byte
	if imgs := productDomain.GetImages(); len(imgs) > 0 {
		imagesJSON, _ = json.Marshal(imgs)
	} else {
		imagesJSON = nil
	}

	purchasePrice, _ := strconv.ParseFloat(productDomain.GetPurchasePrice(), 64)
	salePrice, _ := strconv.ParseFloat(productDomain.GetSalePrice(), 64)

	entity := &entity.ProductEntity{
		Name:          productDomain.GetName(),
		Description:   productDomain.GetDescription(),
		Mark:          productDomain.GetMark(),
		PurchasePrice: purchasePrice,
		SalePrice:     salePrice,
	}

	if imagesJSON != nil {
		entity.Images = string(imagesJSON)
	}

	if productDomain.GetID() != "" {
		var id int
		fmt.Sscanf(productDomain.GetID(), "%d", &id)
		entity.ID = id
	}

	return entity
}

func ConvertStockDomainToEntity(stockDomain stockModel.StockDomainInterface) *entity.StockEntity {
	entity := &entity.StockEntity{
		ProductID: stockDomain.GetProductID(),
		UserID:    stockDomain.GetUserID(),
		Type:      stockDomain.GetStockType(),
		Quantity:  stockDomain.GetQuantity(),
		Reason:    stockDomain.GetReason(),
	}

	if stockDomain.GetID() != "" {
		var id int
		if _, err := fmt.Sscanf(stockDomain.GetID(), "%d", &id); err == nil {
			entity.ID = id
		}
	}

	return entity
}
