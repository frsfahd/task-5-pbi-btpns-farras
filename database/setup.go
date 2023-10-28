package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// var db = os.Getenv("NAME_DB")
	// var host = os.Getenv("HOST_DB")
	// var user = os.Getenv("USER_DB")
	// var pass = os.Getenv("PASS_DB")
	// var port, _ = strconv.Atoi(os.Getenv("PORT_DB"))
	var err error

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, pass, db, port)
	dsn := "postgres://vaplgfdx:I_pJSa2fh-3Tvu-257LLEVJZhhtSjxdd@satao.db.elephantsql.com/vaplgfdx"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
