package converter

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kevynlohan05/StockPro/src/model/entity"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
	stockModel "github.com/kevynlohan05/StockPro/src/model/stock"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func ConvertUserEntityToDomain(entity entity.UserEntity) userModel.UserDomainInterface {
	userDomain := userModel.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
		entity.Sector,
		entity.Role,
		entity.IsActive,
	)

	userDomain.SetID(fmt.Sprintf("%d", entity.ID))

	return userDomain
}

func ConvertProductEntityToDomain(entity entity.ProductEntity) productModel.ProductDomainInterface {
	var images []string
	err := json.Unmarshal([]byte(entity.Images), &images)
	if err != nil {
		log.Printf("Error unmarshaling images: %v\n", err)
		images = []string{} // fallback to empty
	}

	productDomain := productModel.NewProductDomainService(
		entity.Name,
		entity.Description,
		entity.Mark,
		fmt.Sprintf("%.2f", entity.PurchasePrice),
		fmt.Sprintf("%.2f", entity.SalePrice),
		images,
	)

	productDomain.SetID(fmt.Sprintf("%d", entity.ID))

	return productDomain
}

func ConvertStockEntityToDomain(entity entity.StockEntity) stockModel.StockDomainInterface {
	stockDomain := stockModel.NewStockEntity(
		entity.ProductID,
		entity.UserID,
		entity.Type,
		entity.Quantity,
		entity.Reason,
		entity.CreatedAt,
	)

	stockDomain.SetID(fmt.Sprintf("%d", entity.ID))

	return stockDomain
}

func ConvertStockQuantityEntityToDomain(entity entity.StockQuantityEntity) stockModel.StockQuantityDomainInterface {
	stockDomain := stockModel.NewStockQuantityEntity(
		entity.ProductID,
		entity.Quantity,
	)
	return stockDomain
}
