package main

import (
	"github.com/achmadsy/go-project/db"
	"github.com/achmadsy/go-project/routes"
)

func main() {
	db.InitDB()
	routes.Run()
}
