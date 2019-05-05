package impl

import (
	// Go native packages
	"net/http"
	"strconv"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"

	// Dep packages
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ListContactQuery struct {
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
}

func ListContact(c *gin.Context) {
	var query ListContactQuery
	err := c.BindQuery(&query)
	if err != nil {
		env.Environment.Logger.Error("Bind query error: " + err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	contacts, err := env.Environment.DataStore.SelectContact(c, query.FirstName, query.LastName,
		query.PhoneNumber, query.Email)
	if err != nil {
		params := make(map[string]interface{})
		params["firtname"] = query.FirstName
		params["lastname"] = query.FirstName
		params["phoneNumber"] = query.PhoneNumber
		params["email"] = query.Email
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(
			"Select contacts using query params error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

func GetContact(c *gin.Context) {
	contactId, err := strconv.ParseInt(c.Param("contact-id"), 10, 64)
	if err != nil {
		// For the cases that contact-id is nil
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	contact, err := env.Environment.DataStore.SelectContactWithId(c, contactId)
	if err != nil {
		params := make(map[string]interface{})
		params["contactId"] = contactId
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(
			"Select contact with contact Id: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"contact": contact})
}

type CreateContactRequest struct {
	FirstName   string `json:"first_name" form:"first_name" binding:"required"`
	LastName    string `json:"last_name" form:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Email       string `json:"email" form:"email"`
}

func CreateContact(c *gin.Context) {
	var request CreateContactRequest
	err := c.BindJSON(&request)
	if err != nil {
		env.Environment.Logger.Error("Bind json error: " + err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	contactId, err := env.Environment.DataStore.InsertContact(c, request.FirstName,
		request.LastName, request.PhoneNumber, request.Email)
	if err != nil {
		params := make(map[string]interface{})
		params["firtname"] = request.FirstName
		params["lastname"] = request.FirstName
		params["phoneNumber"] = request.PhoneNumber
		params["email"] = request.Email
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(
			"Insert contacts error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"contact_id": contactId})
}

func DeleteContact(c *gin.Context) {
	contactId, err := strconv.ParseInt(c.Param("contact-id"), 10, 64)
	if err != nil {
		// For the cases that contact-id is nil
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = env.Environment.DataStore.DeleteContactWithId(c, contactId)
	if err != nil {
		params := make(map[string]interface{})
		params["contactId"] = contactId
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(
			"Delete contact with contact id error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type UpdateContactRequest struct {
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
}

func UpdateContact(c *gin.Context) {
	var request UpdateContactRequest
	err := c.BindJSON(&request)
	if err != nil {
		env.Environment.Logger.Error("Bind json error: " + err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	contactId, err := strconv.ParseInt(c.Param("contact-id"), 10, 64)
	if err != nil {
		env.Environment.Logger.Error("Bind json error: " + err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = env.Environment.DataStore.UpdateContact(c, contactId, request.FirstName,
		request.LastName, request.PhoneNumber, request.Email)
	if err != nil {
		params := make(map[string]interface{})
		params["email"] = request.Email
		params["first-name"] = request.FirstName
		params["last-name"] = request.LastName
		params["phone-number"] = request.PhoneNumber
		env.Environment.Logger.WithFields(logrus.Fields{"params": params}).Error(
			"Update contact error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
