package main

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/collector"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%s", c)

	con := fritz.NewHomeAuto(fritz.SkipTLSVerify(), fritz.Credentials(c.Credentials.Username, c.Credentials.Password))
	if err := con.Login(); err != nil {
		log.Error(err)
		return
	}

	errChan := make(chan error, 1)
	go func(ch chan<- error) {
		ch <- collector.CollectMetrics(con)
	}(errChan)

	log.Debug("starting http endpoint")
	http.Handle("/metrics", promhttp.Handler())
	go func(ch chan<- error) {
		ch <- http.ListenAndServe(":2112", nil)
	}(errChan)
	err = <-errChan
	if err != nil {
		log.Error(err)
		return
	}
}
