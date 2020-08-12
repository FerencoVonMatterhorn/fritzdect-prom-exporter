package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := setLogLevel(getCLArgLoglevel())
	if err != nil {
		log.Error(err)
		return
	}

}

func getCLArgLoglevel() string {
	loglevel := flag.String("Loglevel",log.DebugLevel.String(),"Set the Loglevel")
	flag.Parse()
	return *loglevel
}

func setLogLevel(loglevel string) error {
	lvl, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}