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
