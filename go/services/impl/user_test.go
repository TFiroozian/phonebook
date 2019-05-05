package impl

import (
	// Go native packages
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"testing"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	models "github.com/tfiroozian/phonebook/go/models/mock"
	structs "github.com/tfiroozian/phonebook/go/models/tmpl"

	// Dep packages
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListContactSuccessfully(t *testing.T) {
	defer GetEnv(t).Finish()
	query := ListContactQuery{
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "",
		PhoneNumber: "",
	}

	contacts := []structs.Contact{
		structs.Contact{Id: 1,
			FirstName:   "first-name",
			LastName:    "last-name",
			Email:       "first-last@gmail.com",
			PhoneNumber: "01234567890",
		},

		structs.Contact{
			Id:          2,
			FirstName:   "first-name",
			LastName:    "last-name",
			PhoneNumber: "09876543210",
		},
	}

	env.Environment.DataStore.(*models.MockDataStore).EXPECT().SelectContact(gomock.Any(),
		query.FirstName, query.LastName, query.PhoneNumber, query.Email).Return(&contacts, nil)

	router := SetupRouter()
	w := performRequest(t, router, "GET", "/api/v0/contacts?first_name=first-name&last_name=last-name", nil)

	body, _ := ioutil.ReadAll(w.Result().Body)
	m := make(map[string]interface{})
	m["contacts"] = contacts
	b, _ := json.Marshal(m)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, b)
}

func TestGetContactSuccessfully(t *testing.T) {
	defer GetEnv(t).Finish()
	var contactId int64 = 1

	contact := structs.Contact{
		Id:          1,
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "first-last@gmail.com",
		PhoneNumber: "01234567890",
	}

	env.Environment.DataStore.(*models.MockDataStore).EXPECT().SelectContactWithId(gomock.Any(),
		contactId).Return(&contact, nil)

	router := SetupRouter()
	w := performRequest(t, router, "GET", path.Join("/api/v0/contacts",
		strconv.FormatInt(contactId, 10)), nil)

	body, _ := ioutil.ReadAll(w.Result().Body)
	m := make(map[string]interface{})
	m["contact"] = contact
	b, _ := json.Marshal(m)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, b)
}

func TestCreateContactSuccessfully(t *testing.T) {
	defer GetEnv(t).Finish()
	form := CreateContactRequest{
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "first-last@gmail.com",
		PhoneNumber: "01234567890",
	}

	var (
		contactId int64 = 1
		newId     int64 = 1
	)

	env.Environment.DataStore.(*models.MockDataStore).EXPECT().InsertContact(gomock.Any(),
		form.FirstName, form.LastName, form.PhoneNumber, form.Email).Return(newId, nil)

	router := SetupRouter()
	w := performRequest(t, router, "POST", "/api/v0/contacts", form)

	body, _ := ioutil.ReadAll(w.Result().Body)
	m := make(map[string]interface{})
	m["contact_id"] = contactId
	b, _ := json.Marshal(m)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, b)
}

func TestDeleteContactSuccessfully(t *testing.T) {
	defer GetEnv(t).Finish()
	var contactId int64 = 1

	env.Environment.DataStore.(*models.MockDataStore).EXPECT().DeleteContactWithId(gomock.Any(),
		contactId).Return(nil)

	router := SetupRouter()
	w := performRequest(t, router, "DELETE", path.Join("/api/v0/contacts",
		strconv.FormatInt(contactId, 10)), nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateContactSuccessfully(t *testing.T) {
	defer GetEnv(t).Finish()
	form := UpdateContactRequest{
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "first-last@gmail.com",
		PhoneNumber: "01234567890",
	}

	var contactId int64 = 1
	env.Environment.DataStore.(*models.MockDataStore).EXPECT().UpdateContact(gomock.Any(),
		contactId, form.FirstName, form.LastName, form.PhoneNumber, form.Email).Return(nil)

	router := SetupRouter()
	w := performRequest(t, router, "PUT", path.Join("/api/v0/contacts",
		strconv.FormatInt(contactId, 10)), form)
	assert.Equal(t, http.StatusOK, w.Code)
}
