package service

import (
	mcdaily "McDailyAutoRun/McDaily"
	"McDailyAutoRun/config"
	lineapi "McDailyAutoRun/lineAPI"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
)

// RunService Run service
func RunService() {
	go startRestfulAPI()

	gocron.Every(1).Day().Do(task)
	<-gocron.Start()
}

func startRestfulAPI() {
	fmt.Printf("Start listen port: %s\n", config.TCPPort)

	r := mux.NewRouter()

	// line bot router
	r.HandleFunc("/linebot", lineapi.EventHandle)
	r.HandleFunc("/line/quota", lineapi.GetMessageQuota)

	http.ListenAndServe(":"+config.TCPPort, r)
}

func task() {
	for _, v := range config.User {
		go mcdaily.GetLottery(v)
	}
}
