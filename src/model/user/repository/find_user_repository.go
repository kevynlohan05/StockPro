package repository

import (
	"database/sql"
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
	"github.com/kevynlohan05/StockPro/src/model/converter"
	userModel "github.com/kevynlohan05/StockPro/src/model/user"
	"github.com/kevynlohan05/StockPro/src/model/user/repository/entity"
)

func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Iniciando FindUserById repository")

	query := `SELECT id, name, email, password, sector, role, is_active FROM users WHERE id = ?`

	row := ur.db.QueryRow(query, id)

	var entity entity.UserEntity

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Sector,
		&entity.Role,
		&entity.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Usuário não encontrado")
		}
		log.Println("Erro ao escanear resultado da query: ", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar usuário")
	}

	return converter.ConvertEntityToDomain(entity), nil
}
