package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/husseinamine/flora/apps"
)

type Users struct {
	router *mux.Router
	app    *apps.Users
}

func (u *Users) Initialize() {
	router := u.router
	handlers := u.app

	CUrouter := router.NewRoute().Subrouter() // create/update router
	CUrouter.Use(apps.UserExistsMiddleware)   // register middlewares

	// register routes
	router.HandleFunc("/", handlers.GET).Methods(http.MethodGet)
	CUrouter.HandleFunc("/", handlers.POST).Methods(http.MethodPost)
	CUrouter.HandleFunc("/{id:[0-9]+}/", handlers.PUT).Methods(http.MethodPut)
}

func NewUsers(router *mux.Router, app *apps.Users) *Users {
	router = router.PathPrefix("/users").Subrouter()

	return &Users{router, app}
}
