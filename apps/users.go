package apps

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/husseinamine/florasrv/controllers"
)

var BadRequestError = errors.New("Bad Request!")

type Users struct {
	logger *log.Logger
}

type UserKey struct{}

func (u *Users) GET(rw http.ResponseWriter, r *http.Request) {
	if err := controllers.Users.ToJSON(rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (u *Users) POST(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey{}).(*controllers.User)

	if err := controllers.Users.AddUser(user); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%#v\n", controllers.Users[len(controllers.Users)-1])
}

func (u *Users) PUT(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey{}).(*controllers.User)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err := controllers.Users.UpdateUser(user, id); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%#v\n", controllers.Users[id])
}

func NewUsers(logger *log.Logger) *Users {
	return &Users{logger}
}

func UserExistsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := controllers.Users.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey{}, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
