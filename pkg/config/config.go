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
	return fmt.Sprintf("\nSet loglevel to %s\nusername: %s\npassword: %s", c.Loglevel, c.Credentials.Username, c.Credentials.Password)
}

func Parse() Config {
	loglevel := flag.String("l", log.InfoLevel.String(), "Set the Loglevel")
	username := flag.String("u", "", "Set the FritzBox User for authentication")
	password := flag.String("p", "", "Set the Fritzbox password for authentication")
	flag.Parse()
	c := Config{
		Loglevel: *loglevel,
		Credentials: FritzBoxCredentials{
			Username: *username,
			Password: *password,
		},
	}
	return c
}
