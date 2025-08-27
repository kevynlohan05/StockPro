package converter

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
	imagesJSON, err := json.Marshal(productDomain.GetImages())
	if err != nil {
		log.Printf("Error marshaling attachment URLs: %v\n", err)
		imagesJSON = []byte("[]") // fallback to empty
	}

	purchasePrice, err := strconv.ParseFloat(productDomain.GetPurchasePrice(), 64)
	if err != nil {
		log.Printf("Erro ao converter PurchasePrice: %v\n", err)
	}

	salePrice, err := strconv.ParseFloat(productDomain.GetSalePrice(), 64)
	if err != nil {
		log.Printf("Erro ao converter SalePrice: %v\n", err)
	}

	entity := &entity.ProductEntity{
		Name:          productDomain.GetName(),
		Description:   productDomain.GetDescription(),
		Mark:          productDomain.GetMark(),
		PurchasePrice: purchasePrice,
		SalePrice:     salePrice,
		Images:        string(imagesJSON),
	}

	if productDomain.GetID() != "" {
		var id int
		if _, err := fmt.Sscanf(productDomain.GetID(), "%d", &id); err == nil {
			entity.ID = id
		}
	}

	return entity
}
