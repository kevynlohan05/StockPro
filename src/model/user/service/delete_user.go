package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUserServices(userId string) *rest_err.RestErr {
	log.Println("Iniciando DeleteUserServices service")

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		log.Println("Erro ao deletar usuário no banco de dados: ", err)
		return err
	}

	return nil
}
