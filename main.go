package main

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := config.Parse()
	err := setLogLevel(c.Loglevel)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("%s", c)
	connection, err := connectToFritzbox(c.Credentials)
	if err != nil {
		log.Error(err)
		return
	}

	devs, err := connection.List()
	if err != nil {
		log.Error(err)
		return
	}

	for _, dev := range devs.Switches() {
		log.Infof("Temperatur: %s Celsius", dev.Temperature.FmtCelsius())
	}

}

func connectToFritzbox(credentials config.FritzBoxCredentials) (fritz.HomeAuto, error) {
	fritzConnection := fritz.NewHomeAuto(
		fritz.SkipTLSVerify(),
		fritz.Credentials(credentials.Username, credentials.Password),
	)
	err := fritzConnection.Login()
	if err != nil {
		return nil, err
	}
	return fritzConnection, err
}

func setLogLevel(loglevel string) error {
	lvl, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
