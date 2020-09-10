package connection

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"reflect"
	"testing"
)

func TestConnectToFritzbox(t *testing.T) {
	type args struct {
		credentials config.FritzBoxCredentials
	}
	tests := []struct {
		name    string
		args    args
		want    fritz.HomeAuto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectToFritzbox(tt.args.credentials)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToFritzbox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToFritzbox() got = %v, want %v", got, tt.want)
			}
		})
	}
}
