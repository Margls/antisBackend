package repositories

import(
	// "context"
	// "database/sql"
    // "antis/backend/models"
    "github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{db:db}
}