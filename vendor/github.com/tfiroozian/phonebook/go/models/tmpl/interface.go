package tmpl

import (
	// Go native package
	"context"
)

type DataStore interface {
	// Contacts
	SelectContactWithId(c context.Context, userId, contactId int64) (*Contact, error)
	DeleteContactWithId(c context.Context, userId, contactId int64) error
	UpdateContact(c context.Context, userId, id int64, firstName, lastName, phoneNumber, email string) error
	InsertContact(c context.Context, userId int64, firstName, lastName,
		phoneNumber, email string) (int64, error)
	SelectContact(c context.Context, userId int64, firstName, lastName,
		phoneNumber, email string) (*[]Contact, error)

	// Users
	SelectUserWithUsername(c context.Context, username string) (*User, error)
}
