package models

import (
	"database/sql"

	"github.com/rafialariq/go-auth/config"
	"github.com/rafialariq/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: db,
	}
}

func (u *UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("SELECT id, first_name, last_name, username, email, password FROM mst_user WHERE "+fieldName+" = $1 LIMIT 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UserModel) Create(user *entities.User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO mst_user (first_name, last_name, username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

// CREATE TABLE mst_user (
//     id SERIAL PRIMARY KEY,
//     first_name VARCHAR(255),
//     last_name VARCHAR(255),
//     username VARCHAR(255),
//     email VARCHAR(255),
//     password VARCHAR(255),
//     created_at TIMESTAMP,
//     updated_at TIMESTAMP
// );
