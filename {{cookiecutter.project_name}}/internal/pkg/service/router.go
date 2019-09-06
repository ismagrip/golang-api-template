package service

import (
	log "github.com/sirupsen/logrus"

	"net/http"

	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *server) setUpRouter(cfg *config.Config) error {

	s.router = chi.NewRouter()
	//Universal middlewares
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	if cfg.Server.Benchmark {
		s.router.Mount("/debug", middleware.Profiler())
	}

	//Custom routes
	s.router.Mount("/first_route",s.{{cookiecutter.project_name}}FirstRoute())

	//for debugging
	printWorkingRoutes(s.router)

	return nil
}

func printWorkingRoutes(router chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.WithFields(log.Fields{
			"METHOD": method,
			"ROUTE":  route,
		}).Info("Register route")
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.WithField("ERROR", err.Error()).Warning("Error walking route")
	}
}
