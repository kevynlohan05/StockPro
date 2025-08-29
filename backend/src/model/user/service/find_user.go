package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userDomainService) FindUserByIdServices(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindUserByIdServices service")

	userResult, err := ud.userRepository.FindUserById(id)
	if err != nil {
		log.Println("Erro ao buscar usuário no banco de dados: ", err)
		return nil, err
	}

	return userResult, nil
}

func (ud *userDomainService) FindAllUsersServices() ([]userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindAllUsersServices service")

	users, err := ud.userRepository.FindAllUsers()
	if err != nil {
		log.Println("Erro ao buscar usuários no banco de dados: ", err)
		return nil, err
	}

	return users, nil
}
