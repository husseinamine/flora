package views

import (
	"log"
	"net/http"

	"github.com/husseinamine/florasrv/controllers"
)

type Users struct {
	logger *log.Logger
}

func NewUsers(logger *log.Logger) *Users {
	return &Users{logger}
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GET(rw, r)
	default:
		rw.WriteHeader(http.StatusNotImplemented)
	}
}

func (u *Users) GET(rw http.ResponseWriter, r *http.Request) {
	if err := controllers.Users.ToJSON(rw); err != nil {
		http.Error(rw, "Something Went Wrong!", http.StatusInternalServerError)
	}
}
