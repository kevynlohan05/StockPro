package converter

import (
	"fmt"

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
	productDomain := productModel.NewProductDomainService(
		entity.Name,
		entity.Description,
		entity.Mark,
		entity.PurchasePrice,
		entity.SalePrice,
		entity.Image,
	)

	productDomain.SetID(fmt.Sprintf("%d", entity.ID))

	return productDomain
}
