package config

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Credentials FritzBoxCredentials
	Exporter    ExporterConfig
}

type FritzBoxCredentials struct {
	Username string
	Password string
}

type ExporterConfig struct {
	Loglevel string
	Interval int
}

func (c Config) String() string {
	return fmt.Sprintf("\nSet loglevel to %s\nusername: %s\npassword: %s\ninterval: %d", c.Exporter.Loglevel, c.Credentials.Username, c.Credentials.Password, c.Exporter.Interval)
}

func Parse() (Config, error) {
	loglevel := flag.String("l", log.InfoLevel.String(), "Set the Loglevel")
	interval := flag.Int("i", 30, "Set the interval in which the exporter scrapes metrics from the dect devices")
	username := flag.String("u", "", "Set the FritzBox User for authentication")
	password := flag.String("p", "", "Set the Fritzbox password for authentication")
	flag.Parse()
	c := Config{
		Exporter: ExporterConfig{
			Loglevel: *loglevel,
			Interval: *interval,
		},
		Credentials: FritzBoxCredentials{
			Username: *username,
			Password: *password,
		},
	}
	return c, setLogLevel(*loglevel)
}

func setLogLevel(loglevel string) error {
	lvl, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
