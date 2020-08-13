package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	c := Config{
		Loglevel: "xy",
		Credentials: FritzBoxCredentials{
			Username: "user",
			Password: "passwd",
		},
	}
	t.Log(c)
}
