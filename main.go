package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

func main() {
	loglevel, userName, password := getCLArgLoglevel()
	err := setLogLevel(loglevel)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Set Loglevel to " + loglevel)
	log.Info("username: " + userName)
	log.Info("password: " + password)
}

func getCLArgLoglevel() (string, string, string) {
	loglevel := flag.String("l",log.DebugLevel.String(),"Set the Loglevel")
	userName := flag.String("u", "", "Set the FritzBox User for authentication")
	password := flag.String("p", "", "Set the Fritzbox password for authentication")
	flag.Parse()
	return *loglevel, *userName, *password
}

func setLogLevel(loglevel string) error {
	lvl, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}