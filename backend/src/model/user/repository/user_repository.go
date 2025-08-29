package repository

import (
	"database/sql"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	LoginUser(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr)
	CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr
	FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr)
	FindAllUsers() ([]userModel.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(userId string) *rest_err.RestErr
}
