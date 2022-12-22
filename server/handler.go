package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) getMyData(context *gin.Context) {
	s.Logger.Info("API 1")
	userData := s.Store.GetAdmin(1)
	fmt.Println(userData)
	//context.Request.
	context.JSON(http.StatusOK, gin.H{"message": "From the api folder"})
}

func (s *Server) getExp(context *gin.Context) {
	s.Logger.Info("API 2")
	tags, _ := s.S3Actions.GetBucketTags("icm-pe-dev-us-east-1-data")
	context.JSON(http.StatusOK, gin.H{"message": tags})
}
