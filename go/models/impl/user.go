package impl

import (
	// Go native packages
	"context"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	models "github.com/tfiroozian/phonebook/go/models/tmpl"
)

func (db *DBImpl) SelectContactWithId(c context.Context, contactId int64) (*models.Contact, error) {
	query := "SELECT id, first_name, last_name, email, phone_number FROM " + env.ContactTable + " WHERE id=$1"
	var contact models.Contact
	err := db.QueryRowxContext(c, query, contactId).StructScan(&contact)
	return &contact, err
}

func (db *DBImpl) SelectContact(c context.Context, firstName, lastName, phoneNumber,
	email string) (*[]models.Contact, error) {
	query := `SELECT * FROM ` + env.ContactTable + ` WHERE ($1='' OR first_name=$1) AND 
	($2='' OR last_name=$2) AND ($3='' OR phone_number=$3) AND ($4='' OR email=$4)`
	rows, err := db.QueryxContext(c, query, firstName, lastName, phoneNumber, email)
	if err != nil {
		return nil, err
	}

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err = rows.StructScan(&contact)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return &contacts, nil
}

func (db *DBImpl) DeleteContactWithId(c context.Context, contactId int64) error {
	query := `DELETE FROM ` + env.ContactTable + ` WHERE id=$1`
	result, err := db.ExecContext(c, query, contactId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
	}

	return nil
}

func (db *DBImpl) InsertContact(c context.Context, firstName, lastName, phoneNumber, email string) (int64, error) {
	var contactId int64
	query := `INSERT INTO ` + env.ContactTable + ` (first_name, last_name, phone_number, email) VALUES($1,
	$2, $3, $4) RETURNING id`
	err := db.QueryRowxContext(c, query, firstName, lastName, phoneNumber, email).Scan(&contactId)
	return contactId, err
}
