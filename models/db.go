package models

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (db *gorm.DB) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", username, password, host, port)
	dsndb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsndb), &gorm.Config{})
	if err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("DB ReOpen 1: " + err.Error())
			}
			err = CreateDB(db)
			if err != nil {
				panic("DB Create: " + err.Error())
			}
			db, err = gorm.Open(mysql.Open(dsndb), &gorm.Config{})
			if err != nil {
				panic("DB ReOpen 2: " + err.Error())
			}
		} else {
			panic("DB Connection:" + err.Error())
		}
	}

	err = db.AutoMigrate(
		&Customer{},
		&Cashbox{},
		&Table{},
		&Monitor{},
		&User{},
	)
	if err != nil {
		panic("DB Migration: " + err.Error())
	}

	CreateUsers(db)

	return
}

func DropDB(db *gorm.DB) (err error) {
	err = db.Exec("DROP DATABASE " + os.Getenv("DB_DATABASE")).Error
	if err != nil {
		return
	}
	return
}

func CreateDB(db *gorm.DB) (err error) {
	err = db.Exec("CREATE DATABASE " + os.Getenv("DB_DATABASE")).Error
	if err != nil {
		return
	}
	return
}

func UseDB(db *gorm.DB, database string) (err error) {
	err = db.Exec("USE DATABASE " + database).Error
	if err != nil {
		return
	}
	return
}
