package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
)

func (ur *userRepository) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {
	log.Println("Iniciando UpdateUser repository")

	var currentUser struct {
		Name     string
		Email    string
		Password string
		Sector   string
		Role     string
		IsActive bool
	}

	queryFind := `SELECT name, email, password, sector, role, is_active FROM users WHERE id = ?`

	err := ur.db.QueryRow(queryFind, userId).Scan(
		&currentUser.Name,
		&currentUser.Email,
		&currentUser.Password,
		&currentUser.Sector,
		&currentUser.Role,
		&currentUser.IsActive,
	)
	if err != nil {
		log.Println("Erro ao buscar usuário existente: ", err)
		return rest_err.NewInternalServerError("Erro ao buscar usuário existente")
	}

	value := converter.ConvertUserDomainToEntity(userDomain)

	if value.Name == "" {
		value.Name = currentUser.Name
	}
	if value.Email == "" {
		value.Email = currentUser.Email
	}
	if value.Password == "" {
		value.Password = currentUser.Password
	}
	if value.Sector == "" {
		value.Sector = currentUser.Sector
	}
	if value.Role == "" {
		value.Role = currentUser.Role
	}

	queryUpdate := `UPDATE users SET name = ?, email = ?, password = ?, sector = ?, role = ? WHERE id = ?`

	result, err := ur.db.Exec(
		queryUpdate,
		value.Name,
		value.Email,
		value.Password,
		value.Sector,
		value.Role,
		userId,
	)

	if err != nil {
		log.Println("Erro ao executar query de atualização de usuário: ", err)
		return rest_err.NewInternalServerError("Erro ao atualizar usuário")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas: ", err)
		return rest_err.NewInternalServerError("Erro ao atualizar usuário")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Usuário não encontrado")
	}

	log.Println("Usuário atualizado com sucesso!")
	return nil
}
