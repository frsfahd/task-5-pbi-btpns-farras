package main

import (
	"gin-photo-api/database"
	"gin-photo-api/routes"
)

func init() {
	database.LoadEnv()
	database.ConnectToDB()
}

func main() {
	routes.Run()
}
