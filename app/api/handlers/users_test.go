package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filatovw/46klpd6x/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

type userManager struct {
	mock.Mock
}

func (s *userManager) CreateUser(user service.User) error { return nil }
func (s *userManager) DeleteUser(user service.User) error { return nil }
func (s *userManager) Users(limit int, offset int) ([]service.User, error) {
	return []service.User{}, nil
}

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
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	srv := userManager{}
	handler := CreateUserHandler{sugar, &srv}
	handler.ServeHTTP(rr, req)
	result := rr.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	response := genericResponse("")
	json.NewDecoder(result.Body).Decode(&response)
	assert.Equal(t, "OK", string(response))
}

func TestDeleteUserHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sugar := logger.Sugar()
	body := []byte{}
	buf := bytes.NewBuffer(body)
	json.NewEncoder(buf).Encode(deleteUserRequest{"test@gmail.com"})

	req, err := http.NewRequest("DELETE", "/users/", buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	srv := userManager{}
	handler := DeleteUserHandler{sugar, &srv}
	handler.ServeHTTP(rr, req)
	result := rr.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	response := genericResponse("")
	json.NewDecoder(result.Body).Decode(&response)
	assert.Equal(t, "OK", string(response))
}

func TestListUsersHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sugar := logger.Sugar()
	body := []byte{}
	buf := bytes.NewBuffer(body)
	json.NewEncoder(buf).Encode(listUsersRequest{15, 2})

	req, err := http.NewRequest("GET", "/users/", buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	srv := userManager{}
	handler := ListUsersHandler{sugar, &srv}
	handler.ServeHTTP(rr, req)
	result := rr.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	response := listUsersResponse([]User{})
	json.NewDecoder(result.Body).Decode(&response)
	assert.Equal(t, []User{}, []User(response))
}
