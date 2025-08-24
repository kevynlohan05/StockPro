package converter

import (
	"fmt"

	userModel "github.com/kevynlohan05/StockPro/src/model/user"
	"github.com/kevynlohan05/StockPro/src/model/user/repository/entity"
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
