package main

import (
	"database/sql"
	"log"

	"github.com/fozyz/gocom/cmd/api"
	"github.com/fozyz/gocom/config"
	"github.com/fozyz/gocom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8888", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} 

	log.Println("DB: Successfully connected!")
}
