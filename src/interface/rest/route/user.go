package route

import (
	"net/http"

	userHandler "auth-skm/src/interface/rest/handlers/user"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h userHandler.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.Register)
	r.Post("/signin", h.Login)
	return r
}
