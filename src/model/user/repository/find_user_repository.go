package repository

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
