package main

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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

	devs, err := connection.List()
	if err != nil {
		log.Error(err)
		return
	}

	switches := devs.Switches()

	recordMetrics(switches)
	startEndpoint()

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

func startEndpoint() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func recordMetrics(devlist []fritz.Device) {
	go func() {
		for {
			for _, dev := range devlist {
				temp, err := strconv.ParseFloat(dev.Temperature.FmtCelsius(), 64)
				if err != nil {
					panic(err)
				}
				dect_temperature.Set(temp)
			}
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	dect_temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dect_temperature",
		Help: "Temperature of Dect Device",
	})
)
