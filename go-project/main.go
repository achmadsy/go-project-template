package main

import (
	"github.com/achmadsy/go-project-template/go-project/db"
	"github.com/achmadsy/go-project-template/go-project/routes"
)

func main() {
	db.InitDB()
	routes.Run()
}
