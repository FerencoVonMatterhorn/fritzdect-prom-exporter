package collector

import (
	"strconv"
	"time"

	"github.com/bpicode/fritzctl/fritz"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/prometheus/common/log"
)

func CollectMetrics(connection fritz.HomeAuto) error {
	for {
		devs, err := connection.List()
		if err != nil {
			return err
		}

		switches := devs.Switches()
		for _, dev := range switches {
			log.Debug("getting temp from dect device")
			temp, err := strconv.ParseFloat(dev.Temperature.FmtCelsius(), 64)
			if err != nil {
				return err
			}
			log.Debug("getting power in W from dect Device")
			power, err := strconv.ParseFloat(dev.Powermeter.FmtPowerW(), 64)
			if err != nil {
				return err
			}
			log.Debug("getting total power consumption from dect device")
			consumption, err := strconv.ParseFloat(dev.Powermeter.FmtEnergyWh(), 64)
			if err != nil {
				return err
			}
			log.Debug("gettig current switch state of dect device")
			switchState, err := strconv.ParseFloat(dev.Switch.State, 64)
			if err != nil {
				return err
			}
			dect_power.Set(power)
			dect_temperature.Set(temp)
			dect_total_power_consumption.Set(consumption)
			dect_switch_status.Set(switchState)
		}
		time.Sleep(2 * time.Minute)
	}
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

var (
	dect_total_power_consumption = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dect_total_power_consumption",
		Help: "Total Power Consumption over Time",
	})
)

var (
	dect_switch_status = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "dect_switch_status",
		Help: "Status of the Switch",
	})
)
