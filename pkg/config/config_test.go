package config

import (
	"testing"
)

func Test_setLogLevel(t *testing.T) {
	t.Parallel()
	type args struct {
		loglevel string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid loglevel",
			args:    args{"info"},
			wantErr: false,
		},
		{
			name:    "invalid loglevel",
			args:    args{"infooooooooooooooooooooo"},
			wantErr: true,
		},
		{
			name:    "insert integer",
			args:    args{"5"},
			wantErr: true,
		},
		{
			name:    "insert special char",
			args:    args{"?"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setLogLevel(tt.args.loglevel); (err != nil) != tt.wantErr {
				t.Errorf("setLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_String(t *testing.T) {
	type fields struct {
		Credentials FritzBoxCredentials
		Exporter    ExporterConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "insert normal values",
			fields: fields{
				Credentials: FritzBoxCredentials{
					Username: "dasisteinUser",
					Password: "test",
				},
				Exporter: ExporterConfig{Loglevel: "info"},
			},
			want: "\nSet loglevel to info\nusername: dasisteinUser\npassword: test\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Credentials: tt.fields.Credentials,
				Exporter:    tt.fields.Exporter,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
