package main

import (
	"log"
	"os"

	inc "./include"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func main() {

	// ENV FILE
	err := godotenv.Load("p-auth-back.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DATABASE CONNECTION
	inc.Db, inc.Err = gorm.Open("sqlite3", "w2g-auth-back.db")
	if inc.Err != nil {
		panic("failed to connect database")
	}
	defer inc.Db.Close()
	inc.Db.LogMode(inc.DbLogMode)

	// OPEN AND USE LOG FILE
	f, err := os.OpenFile("w2g-auth-back.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

}
