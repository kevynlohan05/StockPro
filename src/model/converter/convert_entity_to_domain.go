package converter

import (
	"fmt"

	userModel "github.com/kevynlohan05/StockPro/src/model/user"
	"github.com/kevynlohan05/StockPro/src/model/user/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) userModel.UserDomainInterface {
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
