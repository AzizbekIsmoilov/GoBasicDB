package user

import "database/sql"

type Repository interface {
	FindAll() ([]User, error)
	Create(CreateUserRequest) (*User, error)
}

type postgresRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}