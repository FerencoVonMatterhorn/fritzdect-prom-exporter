package main

import (
	"flag"
	"github.com/bpicode/fritzctl/fritz"
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
	connection,err := connectToFritzbox(userName, password)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(connection.List())
	devs, err := connection.List()
	if err != nil {
		log.Error(err)
		return
	}
	for _, dev := range devs.Switches(){
		log.Info(dev.Temperature.FmtCelsius())
	}

}

func connectToFritzbox(username string, password string) (fritz.HomeAuto, error) {
	fritzConnection := fritz.NewHomeAuto(
		fritz.SkipTLSVerify(),
		fritz.Credentials(username, password),
		)
	err := fritzConnection.Login()
	if err != nil {
		return nil,err
	}
	return fritzConnection,err
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