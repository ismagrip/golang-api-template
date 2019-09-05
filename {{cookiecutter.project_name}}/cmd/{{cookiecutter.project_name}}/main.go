package main

import (
	"os"
	"os/signal"
	"syscall"

	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/config"
	"{{cookiecutter.git_hoster}}/{{cookiecutter.owner}}/{{cookiecutter.project_name}}/internal/pkg/dbclient"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	cfg := config.LoadConfig()
	connection := dbclient.New{{cookiecutter.db_type}}client(cfg)
	defer connection.Close()
	handleSigterm(func() {
		log.Infoln("Captured Ctrl+C")
		connection.Close()
	})

	log.Infof("PID: %d", os.Getpid())
	log.Infof("Server port: %s", cfg.Server.Port)
	log.Errorf("Server could not be started: %s", service.StartServer(connection, storage, cfg).Error())

}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
