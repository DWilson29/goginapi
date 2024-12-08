package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var r *gin.Engine

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "weapons.db")
	if err != nil {
		log.Println(err)
	}
	// Create table
	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS weapons (id INTEGER PRIMARY KEY, type INTEGER, name VARCHAR(64), damage INTEGER, speed INTEGER, reach INTEGER)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table Weapons!")
	}
	statement.Exec()

	// Create
	statement, _ = DB.Prepare("INSERT INTO weapons (type, name, damage, speed, reach) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(0, "Lame Sword", 100, 100, 100)
	log.Println("Inserted the book into database!")

	r = gin.Default()
	initializeRoutes()
	r.Run()
}
