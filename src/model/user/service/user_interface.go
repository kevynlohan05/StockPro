package service

import (
	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
	repositoryUser "github.com/kevynlohan05/StockPro/src/model/user/repository"
)

func NewUserDomainService(userRepository repositoryUser.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repositoryUser.UserRepository
}

type UserDomainService interface {
	CreateUserServices(userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(id string) (userModel.UserDomainInterface, *rest_err.RestErr)
	DeleteUserServices(userId string) *rest_err.RestErr
}
