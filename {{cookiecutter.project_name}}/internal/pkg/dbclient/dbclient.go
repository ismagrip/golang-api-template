package dbclient

import (
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/config"
)

type DBclient interface{
	Close()
}

type DBclientImpl struct{
}
func New{{cookiecutter.db_type}}client(cfg *config.Config) DBclient {
	return nil
}

func (d *DBclientImpl) Close(){

}