package collector

import (
	"github.com/bpicode/fritzctl/fritz"
	"testing"
)

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
			name:    "pass null object",
			args:    struct{ connection fritz.HomeAuto }{connection: fritz.NewHomeAuto()},
			wantErr: true,
		},
		{
			name: "pass object with empty password",
			args: struct{ connection fritz.HomeAuto }{connection: fritz.NewHomeAuto(fritz.SkipTLSVerify(),
				fritz.Credentials("test", ""))},
			wantErr: true,
		},
		{
			name: "pass object with empty user",
			args: struct{ connection fritz.HomeAuto }{connection: fritz.NewHomeAuto(fritz.SkipTLSVerify(),
				fritz.Credentials("", "test"))},
			wantErr: true,
		},
		{
			name: "pass object with empty password and user",
			args: struct{ connection fritz.HomeAuto }{connection: fritz.NewHomeAuto(fritz.SkipTLSVerify(),
				fritz.Credentials("", ""))},
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
