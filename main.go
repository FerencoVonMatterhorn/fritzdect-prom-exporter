package main

import log "github.com/sirupsen/logrus"

func main() {
	initLogger()
}

func initLogger()  {
	log.SetLevel(log.DebugLevel)
}