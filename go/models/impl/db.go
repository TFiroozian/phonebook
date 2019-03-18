package impl

import (
	// Dep packages
	"github.com/jmoiron/sqlx"
)

type DBImpl struct {
	*sqlx.DB
}

func NewDBImpl(dsn string) (*DBImpl, error) {
	db, err := sqlx.Connect("postgres", dsn)
	return &DBImpl{db}, err
}
