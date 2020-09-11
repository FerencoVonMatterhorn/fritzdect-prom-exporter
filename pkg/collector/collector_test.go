package collector

import (
	"errors"
	"github.com/bpicode/fritzctl/fritz"
	"testing"
)

type fakeClient struct {
	listresponse *fritz.Devicelist
	err          error
}

func (fc fakeClient) Login() error {
	return fc.err
}

func (fc fakeClient) List() (*fritz.Devicelist, error) {
	return fc.listresponse, fc.err
}

func (fc fakeClient) On(names ...string) error {
	return fc.err
}

func (fc fakeClient) Off(names ...string) error {
	return fc.err
}

func (fc fakeClient) Toggle(names ...string) error {
	return fc.err
}

func (fc fakeClient) Temp(value float64, names ...string) error {
	return fc.err
}

func TestCollectMetrics(t *testing.T) {
	type args struct {
		connection fritz.HomeAuto
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "abort on error",
			args: args{connection: fakeClient{
				listresponse: nil,
				err:          errors.New("test error"),
			}},
			wantErr: true,
		},
		{
			name: "pass empty temperature",
			args: args{connection: fakeClient{
				listresponse: &fritz.Devicelist{
					Devices: []fritz.Device{
						{
							Identifier:      "123",
							ID:              "1234",
							Functionbitmask: "512",
							Fwversion:       "",
							Manufacturer:    "",
							Productname:     "",
							Present:         0,
							Name:            "testDevice",
							Switch:          fritz.Switch{},
							Powermeter:      fritz.Powermeter{},
							Temperature: fritz.Temperature{
								Celsius: "",
								Offset:  "",
							},
							Thermostat: fritz.Thermostat{},
						},
					},
					Groups: nil,
				},
				err: nil,
			}},
			wantErr: true,
		},
		{
			name: "pass empty power in w",
			args: args{connection: fakeClient{
				listresponse: &fritz.Devicelist{
					Devices: []fritz.Device{
						{
							Identifier:      "123",
							ID:              "1234",
							Functionbitmask: "512",
							Fwversion:       "",
							Manufacturer:    "",
							Productname:     "",
							Present:         0,
							Name:            "testDevice",
							Switch:          fritz.Switch{},
							Powermeter: fritz.Powermeter{
								Power:  "",
								Energy: "",
							},
							Temperature: fritz.Temperature{
								Celsius: "20",
								Offset:  "",
							},
							Thermostat: fritz.Thermostat{},
						},
					},
					Groups: nil,
				},
				err: nil,
			}},
			wantErr: true,
		},
		{
			name: "pass empty energy in wh",
			args: args{connection: fakeClient{
				listresponse: &fritz.Devicelist{
					Devices: []fritz.Device{
						{
							Identifier:      "123",
							ID:              "1234",
							Functionbitmask: "512",
							Fwversion:       "",
							Manufacturer:    "",
							Productname:     "",
							Present:         0,
							Name:            "testDevice",
							Switch:          fritz.Switch{},
							Powermeter: fritz.Powermeter{
								Power:  "50",
								Energy: "",
							},
							Temperature: fritz.Temperature{
								Celsius: "20",
								Offset:  "",
							},
							Thermostat: fritz.Thermostat{},
						},
					},
					Groups: nil,
				},
				err: nil,
			}},
			wantErr: true,
		},
		{
			name: "pass empty switch state",
			args: args{connection: fakeClient{
				listresponse: &fritz.Devicelist{
					Devices: []fritz.Device{
						{
							Identifier:      "123",
							ID:              "1234",
							Functionbitmask: "512",
							Fwversion:       "",
							Manufacturer:    "",
							Productname:     "",
							Present:         0,
							Name:            "testDevice",
							Switch: fritz.Switch{
								State:      "",
								Mode:       "",
								Lock:       "",
								DeviceLock: "",
							},
							Powermeter: fritz.Powermeter{
								Power:  "50",
								Energy: "50",
							},
							Temperature: fritz.Temperature{
								Celsius: "20",
								Offset:  "",
							},
							Thermostat: fritz.Thermostat{},
						},
					},
					Groups: nil,
				},
				err: nil,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CollectMetrics(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("CollectMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
