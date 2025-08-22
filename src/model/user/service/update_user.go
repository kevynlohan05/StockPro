package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {
	return nil
}
