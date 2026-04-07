package models

type User struct {
	ID int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required"`
	Email string `db:"json" json:"email" validate:"required,email"`
}