package impl

import (
	// Go native packages
	"testing"

	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	models "github.com/tfiroozian/phonebook/go/models/mock"

	// Dep packages
	"github.com/golang/mock/gomock"
)

func GetEnv(t *testing.T) *gomock.Controller {
	ctrl := gomock.NewController(t)
	env.Environment = env.Env{DataStore: models.NewMockDataStore(ctrl)}
	return ctrl
}
