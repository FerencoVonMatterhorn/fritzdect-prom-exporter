package connection

import (
	"github.com/bpicode/fritzctl/fritz"
	"github.com/ferencovonmatterhorn/fritzdect-prom-exporter/pkg/config"
	"reflect"
	"testing"
)

type fakeHomeAuto struct {
	listresponse *fritz.Devicelist
	err          error
}

func (fha fakeHomeAuto) Login() error {
	return fha.err
}

func (fha fakeHomeAuto) List() (*fritz.Devicelist, error) {
	return fha.listresponse, fha.err
}

func (fha fakeHomeAuto) On(names ...string) error {
	return fha.err
}

func (fha fakeHomeAuto) Off(names ...string) error {
	return fha.err
}

func (fha fakeHomeAuto) Toggle(names ...string) error {
	return fha.err
}

func (fha fakeHomeAuto) Temp(value float64, names ...string) error {
	return fha.err
}

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
		{
			name: "",
			args: args{
				{
					Username: "",
					Password: "",
				},
			},
			want:    fritz.HomeAuto,
			wantErr: false,
		},
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
