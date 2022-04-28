package main

import (
	"chen/db"
	"chen/routes"
)

func main() {
	db.OpenSQLiteConnection()
	server := routes.NewServer(db.GetSQLiteConnection())
	server.ServerListen()
}
