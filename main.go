package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	db "github.com/freedommmoto/test_simplebank/db/sqlc"
	tool "github.com/freedommmoto/test_simplebank/tool"
)

var MainQueries *db.Queries
var Config tool.ConfigObject

func connectDB() {
	config, err := tool.LoadConfig(".")
	fmt.Printf("%v", config)
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	MainQueries = db.New(conn)
}

func main() {

	//set time-zone
	loc, errTime := time.LoadLocation("UTC")
	if errTime != nil {
		log.Println("unable to set time zone:", errTime)
	}
	time.Local = loc

	//load config
	config, err := tool.LoadConfig(".")
	//fmt.Printf("%v", config)
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}
	Config = config
	connectDB()

}
