package service

import (
	"net/http"

	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/config"
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/dbclient"
	"github.com/go-chi/chi"
)

type server struct {
	handler *handler
	router  *chi.Mux
}

func StartServer(connection dbclient.DBclient,  cfg *config.Config) error {

	server, err := newServer(connection,cfg)
	if err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Server.Port, server.router)
}

func newServer(connection dbclient.DBclient, cfg *config.Config) (*server, error) {
	handler := newHandler(connection)
	server := server{
		handler,
		nil,
	}
	server.setUpRouter(cfg)
	return &server, nil
}
