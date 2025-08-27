package converter

import (
	"fmt"

	"github.com/kevynlohan05/StockPro/src/model/entity"
	productModel "github.com/kevynlohan05/StockPro/src/model/product"
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
	entity := &entity.ProductEntity{
		Name:          productDomain.GetName(),
		Description:   productDomain.GetDescription(),
		Mark:          productDomain.GetMark(),
		PurchasePrice: productDomain.GetPurchasePrice(),
		SalePrice:     productDomain.GetSalePrice(),
		Image:         productDomain.GetDescription(),
	}

	if productDomain.GetID() != "" {
		var id int
		if _, err := fmt.Sscanf(productDomain.GetID(), "%d", &id); err == nil {
			entity.ID = id
		}
	}

	return entity
}
