package impl

import (
	// Go native packages
	"context"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	models "github.com/tfiroozian/phonebook/go/models/tmpl"
)

func (db *DBImpl) SelectUserWithUsername(c context.Context, username string) (*models.User, error) {
	query := "SELECT id, username, password FROM " + env.UserTable + " WHERE username=$1"
	var user models.User
	err := db.QueryRowxContext(c, query, username).StructScan(&user)
	return &user, err
}
