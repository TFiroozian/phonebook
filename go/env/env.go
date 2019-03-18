package env

import (
	// Our packages
	models "github.com/tfiroozian/phonebook/go/models/tmpl"

	// Dep packages
	"github.com/sirupsen/logrus"
)

type Env struct {
	DataStore models.DataStore
	Logger    *logrus.Logger
}

var Environment Env

const (
	// In production mode we're going to read postgres schema name from k8s configmaps
	schemaName   = "app"
	ContactTable = schemaName + ".contact"
)
