package main

import (
	"net/http"

	httpHandlers "github.com/b-pagis/go-public-errors-example/http"
	"github.com/b-pagis/go-public-errors-example/repository/current"
	"github.com/b-pagis/go-public-errors-example/repository/legacy"
	"github.com/b-pagis/go-public-errors-example/users"
)

func main() {
	users := users.UserFinder{
		DB:       current.Repository{},
		LegacyDB: legacy.Repository{},
	}
	handler := httpHandlers.Handlers{
		UserFinder: users,
	}
	http.HandleFunc("/internal", handler.FindOnlyInternalError)
	http.HandleFunc("/public", handler.FindPublicErrors)
	http.HandleFunc("/mid", httpHandlers.HandleError(handler.FindPublicErrorsForMid))
	http.ListenAndServe(":8080", nil)
}
