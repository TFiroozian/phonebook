package tmpl

import (
	// Go native package
	"context"
)

type DataStore interface {
	SelectContactWithId(c context.Context, contactId int64) (*Contact, error)
	SelectContact(c context.Context, firstName, lastName, phoneNumber, email string) (*[]Contact, error)
	DeleteContactWithId(c context.Context, contactId int64) error
	InsertContact(c context.Context, firstName, lastName, phoneNumber, email string) (int64, error)
	UpdateContact(c context.Context, id int64, firstName, lastName, phoneNumber, email string) error
}
