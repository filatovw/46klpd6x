package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filatovw/46klpd6x/internal/service/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

type Service struct {
	mock.Mock
}

func (s *Service) CreateUser(user user.User) error                  { return nil }
func (s *Service) DeleteUser(user user.User) error                  { return nil }
func (s *Service) Users(limit int, offset int) ([]user.User, error) { return []user.User{}, nil }

func TestCreateUserHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sugar := logger.Sugar()
	body := []byte{}
	buf := bytes.NewBuffer(body)
	json.NewEncoder(buf).Encode(createUserRequest{"test@gmail.com", "pass1", "ahaha bobob"})

	req, err := http.NewRequest("POST", "/users/", buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, err)
	rr := httptest.NewRecorder()
	srv := Service{}
	handler := CreateUserHandler{sugar, srv}
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}
