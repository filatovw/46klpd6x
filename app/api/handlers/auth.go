package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/filatovw/46klpd6x/internal/helpers"
	"github.com/filatovw/46klpd6x/internal/service/auth"
	"go.uber.org/zap"
)

type signInRequest struct {
	User
}

// SignInHandler sign-in user by (email, password)
type SignInHandler struct {
	Logger      *zap.SugaredLogger
	AuthService auth.Service
}

func (h SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req signInRequest
	err := json.NewDecoder(io.LimitReader(r.Body, 1024<<2)).Decode(&req)
	if err != nil {
		h.Logger.Errorf("failed to read request body: %s", err)
		helpers.WriteBadRequest(w, "can't read body")
		return
	}
	item := auth.User{}
	token, err := h.AuthService.SignIn(item)
	if err != nil {
		h.Logger.Errorf("failed to sign in user: %s", err)
		helpers.WriteBadRequest(w, "bad request")
		return
	}
	helpers.WriteOK(w, token)
}

// SignOutHandler sign-out user
type SignOutHandler struct {
	Logger      *zap.SugaredLogger
	AuthService auth.Service
}

func (h SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if err := h.AuthService.SignOut(token); err != nil {
		h.Logger.Errorf("failed to signout user with token: %s", token)
		helpers.WriteBadRequest(w, "unable to signout user")
		return
	}
	helpers.WriteOK(w, "OK")
}
