package main

import (
	"chen/db"
	"chen/route"
)

func main() {
	db.OpenSQLiteConnection()
	server := route.NewServer(db.GetSQLiteConnection())
	server.ServerListen()
}
