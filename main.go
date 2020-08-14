package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

	devs, err := connection.List()
	if err != nil {
		log.Error(err)
		return
	}

	switches := devs.Switches()

	recordMetrics(switches, c.Interval)
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

func recordMetrics(devlist []fritz.Device, interval int) {
	go func() {
		for {
			for _, dev := range devlist {
				log.Debug("getting temp from dect device")
				temp, err := strconv.ParseFloat(dev.Temperature.FmtCelsius(), 64)
				if err != nil {
					panic(err)
				}
				power, err := strconv.ParseFloat(dev.Powermeter.FmtPowerW(), 64)
				if err != nil {
					panic(err)
				}
				dect_power.Set(power)
				dect_temperature.Set(temp)
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}

var (
	dect_temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dect_temperature",
		Help: "Temperature of Dect Device",
	})
)

var (
	dect_power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dect_power",
		Help: "Power consumption of Dect Device",
	})
)
