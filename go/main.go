package main

import (
	// Our packages
	"github.com/tfiroozian/phonebook/go/env"
	logger "github.com/tfiroozian/phonebook/go/logger"
	models "github.com/tfiroozian/phonebook/go/models/impl"
	services "github.com/tfiroozian/phonebook/go/services/impl"

	// Dep packages
	_ "github.com/lib/pq"
)

func main() {
	env.Environment.Logger = logger.NewLogger()
	// In production mode we're going to read postgres username and password from k8s secrets
	// and hostname and dbName from k8s configmaps
	db, err := models.NewDBImpl("postgres://postgres:postgres@127.0.0.1/app?sslmode=disable")
	if err != nil {
		env.Environment.Logger.Fatal("db connection failed: " + err.Error())
	}

	defer db.Close()

	env.Environment.DataStore = db
	env.Environment.Middlewares = new(services.MiddlewaresImpl)
	services.SetupRouter().Run(":8080")
}
