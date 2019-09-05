package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *server) {{cookiecutter.project_name}}FirstRoute() chi.Router {
	r := chi.NewRouter()
	r.With(setQueriesOnCtx).Method(http.MethodGet, "/", errorHandler(s.handler.ExampleHandler))
	return r
}
