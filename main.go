package main

import (
	"net/http"

	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/collector"

	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := config.Parse()
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

	collector.CollectMetrics(connection)

	log.Debug("starting http endpoint")
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Error(err)
		return
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
