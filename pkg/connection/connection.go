package connection

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
)

func ConnectToFritzbox(credentials config.FritzBoxCredentials) (fritz.HomeAuto, error) {
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
