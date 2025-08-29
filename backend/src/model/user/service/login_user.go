package service

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ud *userDomainService) LoginUserServices(userDomain userModel.UserDomainInterface) (string, *rest_err.RestErr) {
	log.Println("Iniciando LoginUserServices service")

	userDomain.EncryptPassword()

	userResult, err := ud.userRepository.LoginUser(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		log.Println("Erro ao fazer login do usuário no banco de dados: ", err)
		return "", err
	}

	if !userResult.GetIsActive() {
		log.Println("Usuário inativo: ", userDomain.GetEmail())
		return "", rest_err.NewBadRequestError("User is inactive")
	}

	log.Println("Gerando token JWT para o usuário: ", userDomain.GetEmail())
	token, err := userResult.GenerateToken()
	if err != nil {
		log.Println("Erro ao gerar token JWT: ", err)
		return "", err
	}

	return token, nil
}
