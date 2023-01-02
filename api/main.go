package main

import (
	"flareAPI/cloud"
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
	cfg.DatabaseConfig.Dsn = "root:password@tcp(localhost:3306)/test?parseTime=true"
	// Initialize or connect to the DB in the beginning, The interfaces has its methods develop each one on its own
	//===================MySQL DB==================================
	//storage, err := database.NewSQLDb(cfg.DatabaseConfig.Dsn)
	//===================Local DB===================================
	storage, err := database.NewLocalDB()
	//======================================================
	defer storage.Close()
	if err != nil {
		log.Fatal(err)
	}
	//===================Logger Instance===================================
	logger := util.NewLogger(os.Stdout, util.LevelAll, true)
	logger.Info("Running for %s Environment", cfg.Environment)
	//====================S3 Client Action==================================
	s3Inst := cloud.S3Action{}
	s3Action, _ := s3Inst.NewClient()
	//======================================================
	servInstance := &server.Server{
		PortAddr:  cfg.PortAddress,
		Logger:    logger,
		Store:     storage,
		S3Actions: s3Action}
	servInstance.Start()
}
