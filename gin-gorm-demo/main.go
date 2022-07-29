package main

import (
	"gorm-gin-test/model"
	"gorm-gin-test/routes"
	"log"
)

func main() {
	_, err := model.SetupDatabase("db.sqlite")
	if err != nil {
		log.Println("failed to setup database", err.Error())
	}
	r := routes.SetUpRouter()
	r.Run("0.0.0.0:8080")
}
