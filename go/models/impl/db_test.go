package impl

import (
	// Go native packages
	"testing"

	// Dep packages
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func NewTestDB(t *testing.T) (*DBImpl, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("an error will occur on db.Begin() call", err)
		return nil, mock, err
	}
	return &DBImpl{sqlx.NewDb(db, "sqlmock")}, mock, err
}
