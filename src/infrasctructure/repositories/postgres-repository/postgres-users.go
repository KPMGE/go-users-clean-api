package postgresrepository

import (
	"database/sql"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func (repo *PostgresUserRepository) Save(user *entities.User) error {
	query := `INSERT INTO "users"("id", "created_at", "updated_at", "name", "user_name", "email") VALUES($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(query, user.ID, user.CreatedAt, user.UpdatedAt, user.Name, user.UserName, user.Email)

	return err
}

func (repo *PostgresUserRepository) CheckByEmail(email string) bool {
	query := `SELECT email FROM users WHERE email = ($1)`
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

func (repo *PostgresUserRepository) CheckByUserName(userName string) bool {
	query := `SELECT user_name FROM users WHERE user_name = ($1)`
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

func (repo *PostgresUserRepository) List() []*entities.User {
	var users []*entities.User

	query := `SELECT * FROM users`
	rows, err := repo.db.Query(query)

	defer rows.Close()
	CheckError(err)

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.UserName, &user.Email)
		CheckError(err)
		users = append(users, &user)
	}

	return users
}

func (repo *PostgresUserRepository) Delete(userId string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := repo.db.Exec(query, userId)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostgresUserRepository) CheckById(userId string) bool {
	query := `SELECT id FROM users WHERE id = ($1)`
	rows, err := repo.db.Query(query, userId)
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var foundUserId string

		err = rows.Scan(&foundUserId)
		CheckError(err)

		if foundUserId != "" {
			return true
		}
	}

	return false
}

func (repo *PostgresUserRepository) GetById(userId string) (*entities.User, error) {
	query := `SELECT id, created_at, updated_at, name, user_name, email FROM users WHERE id = ($1)`
	rows := repo.db.QueryRow(query, userId)

	var foundUser entities.User

	err := rows.Scan(&foundUser.ID, &foundUser.CreatedAt, &foundUser.UpdatedAt,
		&foundUser.Name, &foundUser.UserName, &foundUser.Email)

	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}
