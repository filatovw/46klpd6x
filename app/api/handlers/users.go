package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/filatovw/46klpd6x/internal/helpers"
	"github.com/filatovw/46klpd6x/internal/service/user"
	"go.uber.org/zap"
)

// CreateUserHandler create one user
type CreateUserHandler struct {
	Logger      *zap.SugaredLogger
	UserService user.Service
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.Logger.Errorf("failed to read request body: %s", err)
		helpers.WriteBadRequest(w, "can't read body")
		return
	}
	item := user.User{}
	if err := h.UserService.CreateUser(item); err != nil {
		helpers.WriteServerError(w, "failed to create user")
		return
	}
	helpers.WriteOK(w, "OK")
}

// DeleteUserHandler delete one user by email
type DeleteUserHandler struct {
	Logger      *zap.SugaredLogger
	UserService user.Service
}

type deleteUserRequest struct {
	Email string `json:"email"`
}

func (h DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req deleteUserRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.Logger.Errorf("failed to read request body: %s", err)
		helpers.WriteBadRequest(w, "can't read body")
		return
	}
	item := user.User{}
	if err := h.UserService.DeleteUser(item); err != nil {
		helpers.WriteServerError(w, "failed to create user")
		return
	}
	helpers.WriteOK(w, "OK")
}

type listUsersRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type User struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
type listUsersResponse []User

// ListUsersHandler list users in a system
type ListUsersHandler struct {
	Logger      *zap.SugaredLogger
	UserService user.Service
}

func (h ListUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req listUsersRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.Logger.Errorf("failed to read request body: %s", err)
		helpers.WriteBadRequest(w, "can't read body")
		return
	}
	users, err := h.UserService.Users(req.Limit, req.Offset)
	if err != nil {
		helpers.WriteServerError(w, "failed to create user")
		return
	}
	helpers.WriteOK(w, users)
}
