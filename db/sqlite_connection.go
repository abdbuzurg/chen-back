package db

import (
	"chen/db/migration"
	"chen/model"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteConnection interface {
	Get() *gorm.DB
	InitialMigration(r *gin.Engine) error
	autoUpdate() error
}

type sqliteConnection struct {
	db *gorm.DB
}

func NewSQLiteConnection() (SQLiteConnection, error) {
	connection := sqliteConnection{}
	dbSQLiteConnection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("could not get db connection")
		return connection, err
	}

	connection.db = dbSQLiteConnection

	err = connection.autoUpdate()
	if err != nil {
		log.Fatalf("could auto update tables: %v", err)
		return connection, err
	}

	return connection, nil
}

func (liteConn sqliteConnection) Get() *gorm.DB {
	return liteConn.db
}

func (liteConn sqliteConnection) autoUpdate() error {
	return liteConn.db.AutoMigrate(
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
}

func (liteConn sqliteConnection) InitialMigration(r *gin.Engine) error {
	return migration.InitialMigration(r, liteConn.db)
}
