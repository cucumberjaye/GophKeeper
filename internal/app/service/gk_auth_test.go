package service

import (
	"testing"

	"github.com/cucumberjaye/GophKeeper/internal/app/repository/serverrepository/mock_service"
)

// TestKeeperService_CreateToken - тесты для функции CreateToken.
func TestKeeperService_CreateToken(t *testing.T) {
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				login:    "test",
				password: "test",
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				login:    "fail",
				password: "test",
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()

	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.CreateToken(tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
