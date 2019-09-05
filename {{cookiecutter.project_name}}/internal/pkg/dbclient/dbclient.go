package dbclient

import (
	"github.com/sirupsen/logrus"
)

type DBclient interface{

}

type DBclientImpl struct{

}
func New{{cookiecutter.db_type}}client(cfg *config.Config) DBClient {
}