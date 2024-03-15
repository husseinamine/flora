package views

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/husseinamine/florasrv/controllers"
)

var ErrorBadRequest = errors.New("Bad Request!")

type Users struct {
	logger *log.Logger
}

func (u *Users) getIdFromUrlPath(path string) (int, error) {
	r := regexp.MustCompile(`/([0-9]+)`)

	matches := r.FindAllSubmatch([]byte(path), -1)

	if len(matches) != 1 || len(matches[0]) != 2 {
		return -1, ErrorBadRequest
	}

	match := string(matches[0][1])

	return strconv.Atoi(match)
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.GET(rw, r)
	case http.MethodPost:
		u.POST(rw, r)
	case http.MethodPut:
		u.PUT(rw, r)
	default:
		rw.WriteHeader(http.StatusNotImplemented)
	}
}

func (u *Users) GET(rw http.ResponseWriter, r *http.Request) {
	if err := controllers.Users.ToJSON(rw); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (u *Users) POST(rw http.ResponseWriter, r *http.Request) {
	user, err := controllers.Users.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err := controllers.Users.AddUser(user); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%#v\n", controllers.Users[len(controllers.Users)-1])
}

func (u *Users) PUT(rw http.ResponseWriter, r *http.Request) {
	id, err := u.getIdFromUrlPath(r.URL.Path)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := controllers.Users.FromJSON(r.Body)

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
