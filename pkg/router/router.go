package router

import (
	"net/http"

	"github.com/auth0-developer-hub/api_standard-library_golang_hello-world/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func Router(audience, domain string) http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(middleware.NotFoundHandler)

	router.HandlerFunc(http.MethodGet, "/api/messages/public", middleware.PublicApiHandler)
	router.Handler(http.MethodGet, "/api/messages/protected", middleware.ValidateJWT(audience, domain, http.HandlerFunc(middleware.ProtectedApiHandler)))
	router.Handler(http.MethodGet, "/api/messages/admin", middleware.ValidateJWT(audience, domain, http.HandlerFunc(middleware.AdminApiHandler)))

	return middleware.HandleCacheControl(router)
}
