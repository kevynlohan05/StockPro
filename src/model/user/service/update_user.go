package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando UpdateUser service")

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		log.Println("Erro ao atualizar usuário no banco de dados: ", err)
		return err
	}

	return nil
}
