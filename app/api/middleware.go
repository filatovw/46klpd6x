package api

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/filatovw/46klpd6x/internal/helpers"
	"github.com/filatovw/46klpd6x/internal/service/auth"
)

// ContentTypeMiddleware should filter out unsupported content types
type ContentTypeMiddleware struct {
	ContentTypes []string
}

// Middleware function
func (m *ContentTypeMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentTypeHeader := r.Header.Get("Content-Type")
		for _, contentType := range m.ContentTypes {
			if contentType == contentTypeHeader {
				next.ServeHTTP(w, r)
				break
			}
		}
		helpers.WriteResponse(w, "bad request", http.StatusUnsupportedMediaType)
	})
}

var bearerTokenPattern *regexp.Regexp

func init() {
	var err error
	bearerTokenPattern, err = regexp.Compile(`^Bearer\s+([\w]+)`)
	if err != nil {
		panic(err)
	}
}

// AuthMiddleware check user authorization
type AuthMiddleware struct {
	pattern     *regexp.Regexp
	authService auth.Service
	admin       bool
}

// Middleware function
func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := strings.Trim(r.Header.Get("Authorization"), " ")
		token := bearerTokenPattern.Find([]byte(bearerToken))
		if token == nil || len(token) == 0 {
			helpers.WriteResponse(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := m.authService.FindByToken(string(token))
		if err != nil || user == nil {
			helpers.WriteResponse(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if m.admin && !m.authService.IsAdminUser(*user) {
			helpers.WriteResponse(w, "forbidden", http.StatusForbidden)
		}
		next.ServeHTTP(w, r)
	})
}
