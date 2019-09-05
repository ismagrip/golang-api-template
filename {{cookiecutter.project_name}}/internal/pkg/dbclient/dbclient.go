package dbclient

import (
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/config"
	"github.com/sirupsen/logrus"
)

type DBclient interface{

}

type DBclientImpl struct{

}
func New{{cookiecutter.db_type}}client(cfg *config.Config) DBClient {
}