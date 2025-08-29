package repository

import (
	"log"

	"github.com/kevynlohan05/StockPro/src/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	log.Println("Iniciando DeleteUser repository")

	query := `DELETE FROM users WHERE id = ?`

	result, err := ur.db.Exec(query, userId)
	if err != nil {
		log.Println("Erro ao executar query de deleção de usuário: ", err)
		return rest_err.NewInternalServerError("Erro ao deletar usuário")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas: ", err)
		return rest_err.NewInternalServerError("Erro ao deletar usuário")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Usuário não encontrado")
	}

	log.Println("Usuário deletado com sucesso!")

	return nil
}
