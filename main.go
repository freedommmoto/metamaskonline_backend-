package main

import (
	"database/sql"
	api "github.com/freedommmoto/metamaskonline_backend/api"
	db "github.com/freedommmoto/metamaskonline_backend/model/sqlc"
	tool "github.com/freedommmoto/metamaskonline_backend/tool"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var MainQueries *db.Queries
var Config tool.ConfigObject

func connectDB() {
	log.Println("connectDB")
	conn, err := sql.Open(Config.DBDriver, Config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	MainQueries = db.New(conn)
	//user, err := MainQueries.SelectUserID(context.Background(), int32(1))
	//log.Println("user", user)
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
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}
	Config = config
	
	connectDB()

	//make gin server
	server, err := api.NewServer(config, MainQueries)
	if err != nil {
		log.Fatal("can't NewServer", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start server with gin", err)
	}

}
