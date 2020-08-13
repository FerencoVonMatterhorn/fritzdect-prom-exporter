package config

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Loglevel    string
	Credentials FritzBoxCredentials
}

type FritzBoxCredentials struct {
	Username string
	Password string
}

func (c Config) String() string {
	return fmt.Sprintf("Set loglevel to %s \n username: %s \n password: %s", c.Loglevel, c.Credentials.Username, c.Credentials.Password)
}

func Parse() Config {
	c := Config{
		Loglevel: *flag.String("l", log.DebugLevel.String(), "Set the Loglevel"),
		Credentials: FritzBoxCredentials{
			Username: *flag.String("u", "", "Set the FritzBox User for authentication"),
			Password: *flag.String("p", "", "Set the Fritzbox password for authentication"),
		},
	}
	flag.Parse()
	return c
}
