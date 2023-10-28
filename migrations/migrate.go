package main

import (
	"gin-photo-api/database"
	"gin-photo-api/models"
)

func init() {
	database.LoadEnv()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
