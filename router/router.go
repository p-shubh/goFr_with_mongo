package router

import (
	"gofr.dev/pkg/gofr"
	"main.go/config"
	mongodb "main.go/connection/database/mongoDb"
)

type Server struct {
	app    *gofr.App
	db     mongodb.Database
	config *config.Config
}

func (s *Server) GofrEngine(c *config.Config, db mongodb.Database) Server {

	app := gofr.New()
	app.Logger().Debug()
	app.GET("/healthCheck", s.healthCheck)
	app.Run()

	return Server{app: app, db: db, config: c}
}

type statement struct {
	message string
}

func (s *Server) healthCheck(c *gofr.Context) (interface{}, error) {
	return "Server is Up !", nil
}

// func (s server) start() {
// 	s.app.Run()
// }
