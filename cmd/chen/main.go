package main

import (
	"chen/db"
	"chen/server"
	"log"
)

func main() {

	SQLiteConnection, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatal("could establish SQLite connection")
		return
	}

	server := server.NewServer(SQLiteConnection.Get())

	// if err := SQLiteConnection.InitialMigration(server.Router); err != nil {
	// 	log.Fatal("could not make initial migration")
	// 	return
	// }

	server.Listen()
}
