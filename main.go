package main

import (
	"main.go/config"
	mongodb "main.go/connection/database/mongoDb"
	"main.go/router"
)

var app router.Server

func main() {
	c := config.SetConfig()

	db := mongodb.ConnectDatabase(c.Env)

	app.GofrEngine(c, db)

}
