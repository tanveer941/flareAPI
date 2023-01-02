package server

import (
	"flareAPI/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func (s *Server) getMyData(context *gin.Context) {
	s.Logger.Info("API 1: Fetch Data")
	userData := s.Store.GetAdmin(1)
	fmt.Println(userData)
	//context.Request.
	context.JSON(http.StatusOK, gin.H{"message": "From the api folder"})
}

func (s *Server) getExp(context *gin.Context) {
	s.Logger.Info("API 2: Get S3 Bucket tags")
	tags, _ := s.S3Actions.GetBucketTags("icm-pe-dev-us-east-1-data")
	context.JSON(http.StatusOK, gin.H{"message": tags})
}

func (s *Server) getHugeData(context *gin.Context) {
	s.Logger.Info("API 3: Huge data fetch")
	fetchChan := make(chan string, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	util.Fetch3SecData(fetchChan, &wg)
	util.Fetch5SecData(fetchChan, &wg)
	close(fetchChan)
	// Add data in array
	dataArray := make([]string, 0)
	for data := range fetchChan {
		dataArray = append(dataArray, data)
	}
	context.PureJSON(http.StatusOK, dataArray)
}

func (s *Server) getHugeDataConcurrency(context *gin.Context) {
	s.Logger.Info("API 4: Huge data concurrency fetch")
	fetchChan := make(chan string, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go util.Fetch3SecData(fetchChan, &wg)
	go util.Fetch5SecData(fetchChan, &wg)
	wg.Wait()
	close(fetchChan)
	// Add data in array
	dataArray := make([]string, 0)
	for data := range fetchChan {
		dataArray = append(dataArray, data)
	}
	context.PureJSON(http.StatusOK, dataArray)
}
