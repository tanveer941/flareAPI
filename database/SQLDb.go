package database

import (
	"flareAPI/types"
	"github.com/jmoiron/sqlx"
	"time"
)

type DB struct {
	*sqlx.DB
}

func NewSQLDb(dsn string) (*DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)
	return &DB{db}, nil
}

func (m *DB) GetAdmin(id int) *types.User {
	return &types.User{ID: id, Name: "SQLDbAdmin"}
}
