package impl

import (
	// Go native packages
	"context"
	"strconv"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	models "github.com/tfiroozian/phonebook/go/models/tmpl"
)

func (db *DBImpl) SelectContactWithId(c context.Context, userId, contactId int64) (*models.Contact, error) {
	query := "SELECT id, first_name, last_name, email, phone_number FROM " +
		env.ContactTable + " WHERE id=$1 AND user_id=$2"
	var contact models.Contact
	err := db.QueryRowxContext(c, query, contactId, userId).StructScan(&contact)
	return &contact, err
}

func (db *DBImpl) SelectContact(c context.Context, userId int64, firstName, lastName, phoneNumber,
	email string) (*[]models.Contact, error) {
	query := `SELECT id, first_name, last_name, email FROM ` + env.ContactTable +
		` WHERE ($1='' OR first_name=$1) AND ($2='' OR last_name=$2) AND 
		($3='' OR phone_number=$3) AND ($4='' OR email=$4) AND user_id=$5`
	rows, err := db.QueryxContext(c, query, firstName, lastName, phoneNumber, email, userId)
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

func (db *DBImpl) DeleteContactWithId(c context.Context, userId, contactId int64) error {
	query := `DELETE FROM ` + env.ContactTable + ` WHERE id=$1 AND user_id=$2`
	result, err := db.ExecContext(c, query, contactId, userId)
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

func (db *DBImpl) InsertContact(c context.Context, userId int64, firstName, lastName,
	phoneNumber, email string) (int64, error) {
	var contactId int64
	query := `INSERT INTO ` + env.ContactTable + ` (first_name, last_name, phone_number, email, user_id) 
	VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRowxContext(c, query, firstName, lastName, phoneNumber, email, userId).Scan(&contactId)
	return contactId, err
}

func (db *DBImpl) UpdateContact(c context.Context, userId, contactId int64, firstName, lastName, phoneNumber,
	email string) error {
	query := `UPDATE  ` + env.ContactTable + ` SET first_name=$1, last_name=$2, phone_number=$3, 
	email=$4 WHERE id=$5 AND user_id=$6`
	result, err := db.ExecContext(c, query, firstName, lastName, phoneNumber, email, contactId, userId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		env.Environment.Logger.Error("number of affected rows for updating contact with id = "+
			strconv.FormatInt(contactId, 10)+"is:", strconv.FormatInt(rows, 10))
	}

	return nil
}
