package route

import (
	"net/http"

	respondenHandler "survey-skm/src/interface/rest/handlers/responden"

	"github.com/go-chi/chi/v5"
)

func RespondenRouter(h respondenHandler.RespondenHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/quisioner", h.CreateQuisionerData)
	r.Get("/quisioner", h.GetQuisionerDataEachLayanan)
	return r
}
