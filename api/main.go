package main

import (
	"flareAPI/database"
	"flareAPI/server"
	"flareAPI/util"
	"log"
	"os"
)

//type application struct {
//	PortAddr string
//	db       *database.DB
//}

func main() {
	// Initialize or connect to the DB in the beginning, The interfaces has its methods develop each one on its own
	//===================MySQL DB==================================
	//storage, err := database.NewSQLDb("SqlDbDSN")
	//===================Local DB===================================
	storage, err := database.NewLocalDB()
	//======================================================
	defer storage.Close()
	if err != nil {
		log.Fatal(err)
	}
	logger := util.NewLogger(os.Stdout, util.LevelAll, true)
	servInstance := &server.Server{
		PortAddr: ":3001",
		Logger:   logger,
		Store:    storage}
	servInstance.Start()
}
