package api

import (
	"github.com/filatovw/46klpd6x/internal/service/auth"
	"github.com/filatovw/46klpd6x/internal/service/user"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func routes(logger *zap.SugaredLogger, userService user.Service, authService auth.Service) *mux.Router {
	r := mux.NewRouter()
	cmw := ContentTypeMiddleware{ContentTypes: []string{"application/json"}}
	r.Use(cmw.Middleware)

	anyAmw := AuthMiddleware{authService: authService, admin: false}
	adminAmw := AuthMiddleware{authService: authService, admin: true}
	r.Handle("/users/", CreateUserHandler{logger, userService}).Methods("POST")
	r.Handle("/users/", adminAmw.Middleware(DeleteUserHandler{logger, userService})).Methods("DELETE")
	r.Handle("/users/", adminAmw.Middleware(ListUsersHandler{logger, userService})).Methods("GET")

	sub := r.PathPrefix("/auth/").Subrouter()
	sub.Handle("/signin/", SignInHandler{logger, authService}).Methods("POST")
	sub.Handle("/signout/", anyAmw.Middleware(SignOutHandler{logger, authService})).Methods("POST")
	return r
}
