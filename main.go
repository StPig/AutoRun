package main

import (
	"McDailyAutoRun/config"
	"McDailyAutoRun/service"
	"net/http"
)

var processChan chan bool

func main() {
	processChan = make(chan bool)
	go service.RunService()
	go pproff()
	<-processChan
}

func pproff() {
	http.ListenAndServe(":"+config.PprofPort, nil)
}
