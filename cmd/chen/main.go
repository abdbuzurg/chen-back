package main

import (
	"chen/db"
	"chen/routes"
)

func main() {
	db.OpenSQLiteConnection()
	routes.RunRoutes()
}
