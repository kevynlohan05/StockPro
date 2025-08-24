package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userDomainService) CreateUserServices(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando CreateUserServices service")

	userDomain.EncryptPassword()

	userResult, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		log.Println("Erro ao criar usuário no banco de dados: ", err)
		return nil, err
	}

	return userResult, nil
}
