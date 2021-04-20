package main

import (
	db "github.com/achmadsy/go-project-template/go-project/configs"
	"github.com/achmadsy/go-project-template/go-project/routes"
)

func main() {
	db.InitDB()
	routes.Run()
}
