package database

import(
	"database/sql"
	"github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewDB(connStr string) (*sqlx.DB, error)  {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil{
		 return nil, err
	}

	db.SetMaxOpenConns(25)
 	db.SetMaxIdleConns(10)
	return db, nil
}