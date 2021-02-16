package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/Flix14/casbin-example-2/models"
	"github.com/Flix14/casbin-example-2/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dir, _ := os.Getwd()
	err = os.Setenv("CURR_DIR", dir)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("GIN_MODE", os.Getenv("APP_ENV"))
	if err != nil {
		panic(err)
	}

	DB = InitDB()

	r := routes.Init()
	err = r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
}
