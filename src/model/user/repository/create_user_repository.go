package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ur *userRepository) CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando CreateUser repository")

	value := converter.ConvertUserDomainToEntity(userDomain)

	query := `
	INSERT INTO users (name, email, password, sector, role, is_active) 
	VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := ur.db.Exec(query,
		value.Name,
		value.Email,
		value.Password,
		value.Sector,
		value.Role,
		value.IsActive,
	)

	if err != nil {
		log.Println("Erro ao inserir usuário no banco de dados: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao inserir usuário no banco de dados")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Erro ao obter ID do usuário inserido: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao obter ID do usuário inserido")
	}

	value.ID = int(id)

	log.Println("Usuário criado com sucesso!")
	return converter.ConvertEntityToDomain(*value), nil
}
