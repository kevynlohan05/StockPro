package converter

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kevynlohan05/StockPro/src/model/entity"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
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
		entity.PurchasePrice,
		entity.SalePrice,
		images,
	)

	productDomain.SetID(fmt.Sprintf("%d", entity.ID))

	return productDomain
}
