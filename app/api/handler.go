package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/filatovw/46klpd6x/internal/service/auth"
	"github.com/filatovw/46klpd6x/internal/service/user"
	"go.uber.org/zap"
)

func writeResponse(w http.ResponseWriter, response interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ok(w http.ResponseWriter, response interface{}) {
	writeResponse(w, response, http.StatusOK)
}

func serverError(w http.ResponseWriter, response interface{}) {
	writeResponse(w, response, http.StatusInternalServerError)
}

func badRequest(w http.ResponseWriter, response interface{}) {
	writeResponse(w, response, http.StatusBadRequest)
}

// CreateUserHandler create one user
type CreateUserHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

// CreateUserRequest
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req CreateUserHandler
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.logger.Errorf("failed to read request body: %s", err)
		badRequest(w, "can't read body")
		return
	}
	item := user.User{}
	if err := h.userService.CreateUser(item); err != nil {
		serverError(w, "failed to create user")
		return
	}
	ok(w, "OK")
}

// DeleteUserHandler delete one user by email
type DeleteUserHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

type DeleteUserRequest struct {
	Email string `json:"email"`
}

func (h DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.logger.Errorf("failed to read request body: %s", err)
		badRequest(w, "can't read body")
		return
	}
	item := user.User{}
	if err := h.userService.DeleteUser(item); err != nil {
		serverError(w, "failed to create user")
		return
	}
	ok(w, "OK")
}

type ListUsersRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type User struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
type ListUsersResponse []User

// ListUsersHandler list users in a system
type ListUsersHandler struct {
	logger      *zap.SugaredLogger
	userService user.Service
}

func (h ListUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req ListUsersRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.logger.Errorf("failed to read request body: %s", err)
		badRequest(w, "can't read body")
		return
	}
	users, err := h.userService.Users(req.Limit, req.Offset)
	if err != nil {
		serverError(w, "failed to create user")
		return
	}
	ok(w, users)
}

type SignInRequest struct {
	User
}

// SignInHandler sign-in user by (email, password)
type SignInHandler struct {
	logger      *zap.SugaredLogger
	authService auth.Service
}

func (h SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.logger.Errorf("failed to read request body: %s", err)
		badRequest(w, "can't read body")
		return
	}
	item := auth.User{}
	token, err := h.authService.SignIn(item)
	if err != nil {
		h.logger.Errorf("failed to sign in user: %s", err)
		badRequest(w, "bad request")
		return
	}
	ok(w, token)
}

// SignOutHandler sign-out user
type SignOutHandler struct {
	logger      *zap.SugaredLogger
	authService auth.Service
}

func (h SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if err := h.authService.SignOut(token); err != nil {
		h.logger.Errorf("failed to signout user with token: %s", token)
		badRequest(w, "unable to signout user")
		return
	}
	ok(w, "OK")
}
