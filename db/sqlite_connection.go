package db

import (
	"chen/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbSQLite *gorm.DB

func OpenSQLiteConnection() {
	dbSQLiteConnection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	dbSQLite = dbSQLiteConnection
	autoUpdate()
}

func GetSQLiteConnection() *gorm.DB {
	return dbSQLite
}

func autoUpdate() {
	err := dbSQLite.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.Organization{},
		&model.Branch{},
		&model.Hall{},
		&model.Table{},
		&model.Item{},
		&model.Order{},
		&model.OrderList{},
	)
	fmt.Println("Auto update done")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}
