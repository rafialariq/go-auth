package entities

import (
	"time"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// `validate:"required"`
// `validate:"required"`
// `validate:"required"`
// `validate:"required"`
// `validate:"required"`
// CREATE TABLE mst_user (
//     id INT PRIMARY KEY,
//     first_name VARCHAR(255),
//     last_name VARCHAR(255),
//     username VARCHAR(255),
//     email VARCHAR(255),
//     password VARCHAR(255),
//     created_at TIMESTAMP,
//     updated_at TIMESTAMP
// );
