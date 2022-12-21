package main

import (
	"flareAPI/database"
	"flareAPI/server"
	"flareAPI/types"
	"flareAPI/util"
	"log"
	"os"
)

func main() {
	cfg := types.Configuration{}
	cfg.PortAddress = ":5000"
	cfg.Environment = "Development"
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
	logger.Info("Running for %s Environment", cfg.Environment)
	servInstance := &server.Server{
		PortAddr: cfg.PortAddress,
		Logger:   logger,
		Store:    storage}
	servInstance.Start()
}
