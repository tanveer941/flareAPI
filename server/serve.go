package server

import (
	"flareAPI/cloud"
	"flareAPI/database"
	"flareAPI/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	PortAddr  string
	Logger    *util.Logger
	Store     database.Storer
	S3Actions *cloud.S3Action
}

func (s *Server) Start() error {
	gin.ForceConsoleColor()
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "traditional route",
		})
	})
	router.GET("/api1", s.getMyData)
	router.GET("/api2", s.getExp)
	router.GET("/api3", s.getHugeData)
	router.GET("/api4", s.getHugeDataConcurrency)
	return router.Run(s.PortAddr)
}
