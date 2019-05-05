package impl

import (
	// Go native packages
	"context"
	"encoding/json"
	"strconv"
	"testing"

	// Our packages
	models "github.com/tfiroozian/phonebook/go/models/tmpl"

	// Dep packages
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestInsertContactSuccessfully(t *testing.T) {
	db, mock, err := NewTestDB(t)
	assert.NoError(t, err, "not expected error when opening a stub database connection")
	defer db.Close()

	var (
		firstName   = "first-name"
		lastName    = "last-name"
		phoneNumber = "01234567890"
		email       = "first-last@gmail.com"
		insertId    = int64(1)
	)

	mock.ExpectQuery(`INSERT INTO phone_book[.]contacts [(]first_name, last_name, phone_number, email[)]
	VALUES[(].+[)] RETURNING id`).
		WithArgs(firstName, lastName, phoneNumber, email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(insertId))

	contactId, err := db.InsertContact(context.TODO(), firstName, lastName, phoneNumber, email)
	assert.NoError(t, err, "error was not expected while inserting contact")
	assert.Equal(t, insertId, contactId, "ID must be "+strconv.FormatInt(insertId, 10))

	mock.ExpectationsWereMet()
	assert.NoError(t, err, "error was not expected while checking expectation")
}

func TestDeleteContactIdSuccessfully(t *testing.T) {
	db, mock, err := NewTestDB(t)
	assert.NoError(t, err, "not expected error when opening a stub database connection")
	defer db.Close()

	var contactId int64 = 1

	mock.ExpectExec(`DELETE FROM phone_book[.]contacts WHERE id=.+`).
		WithArgs(contactId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = db.DeleteContactWithId(context.TODO(), contactId)
	assert.NoError(t, err, "error was not expected while deleting contact with id")

	mock.ExpectationsWereMet()
	assert.NoError(t, err, "error was not expected while checking expectation")
}

func TestSelectContactWithIdSuccessfully(t *testing.T) {
	db, mock, err := NewTestDB(t)
	assert.NoError(t, err, "not expected error when opening a stub database connection")
	defer db.Close()

	contact := models.Contact{
		Id:          1,
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "first-last@gmail.com",
		PhoneNumber: "01234567890",
	}

	var contactId int64 = 1
	mock.ExpectQuery(`SELECT id, first_name, last_name, email, phone_number FROM phone_book.contacts WHERE id=.+`).
		WithArgs(contactId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"email", "phone_number"}).AddRow(contact.Id, contact.FirstName, contact.LastName,
			contact.Email, contact.PhoneNumber))

	contactStr, _ := json.Marshal(contact)
	resultContact, err := db.SelectContactWithId(context.TODO(), contactId)
	assert.NoError(t, err, "error was not expected while selecting contact with id")
	assert.Equal(t, resultContact, &contact, "contact must be "+string(contactStr))

	mock.ExpectationsWereMet()
	assert.NoError(t, err, "error was not expected while checking expectation")
}

func TestSelectContactSuccessfully(t *testing.T) {
	db, mock, err := NewTestDB(t)
	assert.NoError(t, err, "not expected error when opening a stub database connection")
	defer db.Close()

	contacts := []models.Contact{
		models.Contact{Id: 1,
			FirstName:   "first-name",
			LastName:    "last-name",
			Email:       "first-last@gmail.com",
			PhoneNumber: "01234567890",
		},

		models.Contact{
			Id:          2,
			FirstName:   "first-name",
			LastName:    "last-name",
			PhoneNumber: "09876543210",
		},
	}

	var (
		firstName   = "first-name"
		lastName    = "last-name"
		email       = ""
		phoneNumber = ""
	)

	mock.ExpectQuery(`SELECT [*] FROM phone_book[.]contacts WHERE [(].+='' OR first_name=.+[)] AND 
	[(].+='' OR last_name=.+[)] AND [(].+='' OR phone_number=.+[)] AND [(].+='' OR email=.+[)]`).
		WithArgs(firstName, lastName, phoneNumber, email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name",
			"email", "phone_number"}).AddRow(contacts[0].Id, contacts[0].FirstName,
			contacts[0].LastName, contacts[0].Email, contacts[0].PhoneNumber).AddRow(
			contacts[1].Id, contacts[1].FirstName,
			contacts[1].LastName, contacts[1].Email, contacts[1].PhoneNumber))

	contactsStr, _ := json.Marshal(contacts)
	resultContacts, err := db.SelectContact(context.TODO(), firstName, lastName, phoneNumber, email)
	assert.NoError(t, err, "error was not expected while selecting contact")
	assert.Equal(t, resultContacts, &contacts, "contacts must be "+string(contactsStr))

	mock.ExpectationsWereMet()
	assert.NoError(t, err, "error was not expected while checking expectation")
}
func TestUpdateContactSuccessfully(t *testing.T) {
	db, mock, err := NewTestDB(t)
	assert.NoError(t, err, "not expected error when opening a stub database connection")
	defer db.Close()

	var (
		firstName   = "first-name"
		lastName    = "last-name"
		phoneNumber = "01234567890"
		email       = "first-last@gmail.com"
		contactId   = int64(1)
	)

	mock.ExpectExec(`UPDATE phone_book[.]contacts SET first_name=.+, last_name=.+, phone_number=.+, email=.+
	WHERE id=.+`).WithArgs(firstName, lastName, phoneNumber, email, contactId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = db.UpdateContact(context.TODO(), contactId, firstName, lastName, phoneNumber, email)
	assert.NoError(t, err, "error was not expected while updating contact")

	mock.ExpectationsWereMet()
	assert.NoError(t, err, "error was not expected while checking expectation")
}
