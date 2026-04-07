package models

import "time"


type Item struct{
	ID int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}