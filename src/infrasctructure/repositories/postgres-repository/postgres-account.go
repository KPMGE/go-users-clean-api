package postgresrepository

import (
	"database/sql"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type PostgresAccountRepository struct {
	db *sql.DB
}

func (repo *PostgresAccountRepository) CheckAccountByEmail(email string) bool {
	query := `SELECT email FROM accounts WHERE email = ($1)`
	rows, err := repo.db.Query(query, email)
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var foundEmail string

		err = rows.Scan(&foundEmail)
		CheckError(err)

		if foundEmail != "" {
			return true
		}
	}

	return false
}

func (repo *PostgresAccountRepository) CheckAccountByUserName(userName string) bool {
	query := `SELECT user_name FROM accounts WHERE user_name = ($1)`
	rows, err := repo.db.Query(query, userName)
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var foundUserName string

		err = rows.Scan(&foundUserName)
		CheckError(err)

		if foundUserName != "" {
			return true
		}
	}

	return false
}

func (repo *PostgresAccountRepository) Save(account *entities.Account) error {
	query := `INSERT INTO "accounts"("id", "created_at", "updated_at", "user_name", "email", "password") VALUES($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(query, account.ID, account.CreatedAt, account.UpdatedAt, account.UserName, account.Email, account.Password)
	return err
}

func (repo *PostgresAccountRepository) DeleteAccountById(accountId string) bool {
	query := `DELETE FROM accounts WHERE id = ($1)`
	res, err := repo.db.Exec(query, accountId)
	CheckError(err)

	n, err := res.RowsAffected()
	CheckError(err)

	return n > 0
}

func NewPostgresAccountRepository(db *sql.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}
