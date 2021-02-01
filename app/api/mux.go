package api

import (
	"net/http"

	"github.com/filatovw/46klpd6x/internal/service/auth"
	"github.com/filatovw/46klpd6x/internal/service/user"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func routes(logger *zap.SugaredLogger, userService user.Service, authService auth.Service) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/users", CreateUserHandler{logger, userService}).Methods("POST")
	r.Handle("/users", DeleteUserHandler{logger, userService}).Methods("DELETE")
	r.Handle("/users", ListUsersHandler{logger, userService}).Methods("GET")

	sub := r.PathPrefix("/auth").Subrouter()
	sub.Handle("/signin", SignInHandler{logger, authService}).Methods("POST")
	sub.Handle("/signout", SignOutHandler{logger, authService}).Methods("POST")
	return r
}

type CreateUserHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type DeleteUserHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

func (h DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type ListUsersHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

func (h ListUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type SignInHandler struct {
	logger      *zap.SugaredLogger
	authService auth.Service
}

func (h SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type SignOutHandler struct {
	logger      *zap.SugaredLogger
	authService auth.Service
}

func (h SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
