package service

import (
	"testing"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/repository/serverrepository/mock_service"
)

// TestKeeperService_SetLoginPasswordData - тесты для SetLoginPasswordData.
func TestKeeperService_SetLoginPasswordData(t *testing.T) {
	type args struct {
		userID string
		data   models.LoginPasswordData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.LoginPasswordData{
					Login:        "test",
					Password:     "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.LoginPasswordData{
					Login:        "test",
					Password:     "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.SetLoginPasswordData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.SetLoginPasswordData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_SetTextData - тесты для SetTextData.
func TestKeeperService_SetTextData(t *testing.T) {
	type args struct {
		userID string
		data   models.TextData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.TextData{
					Data:         "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.TextData{
					Data:         "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.SetTextData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.SetTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_SetBinaryData - тесты для SetBinaryData.
func TestKeeperService_SetBinaryData(t *testing.T) {
	type args struct {
		userID string
		data   models.BinaryData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.BinaryData{
					Data:         []byte("test"),
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.BinaryData{
					Data:         []byte("test"),
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.SetBinaryData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.SetBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_SetBankCardData - тесты для SetBankCardData.
func TestKeeperService_SetBankCardData(t *testing.T) {
	type args struct {
		userID string
		data   models.BankCardData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.BankCardData{
					Number:       "test",
					ValidThru:    "test",
					CVV:          "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.BankCardData{
					Number:       "test",
					ValidThru:    "test",
					CVV:          "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.SetBankCardData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.SetBankCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_Sync - тесты для Sync.
func TestKeeperService_Sync(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Sync(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.Sync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// TestKeeperService_DeleteData - тесты для DeleteData.
func TestKeeperService_DeleteData(t *testing.T) {
	type args struct {
		key    string
		userID string
	}
	tests := []struct {
		name    string
		s       *KeeperService
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.DeleteData(tt.args.key, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_UpdateLoginPasswordData - тесты для UpdateLoginPasswordData.
func TestKeeperService_UpdateLoginPasswordData(t *testing.T) {
	type args struct {
		userID string
		data   models.LoginPasswordData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.LoginPasswordData{
					Login:        "test",
					Password:     "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.LoginPasswordData{
					Login:        "test",
					Password:     "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.UpdateLoginPasswordData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.UpdateLoginPasswordData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_UpdateTextData - тесты для UpdateTextData.
func TestKeeperService_UpdateTextData(t *testing.T) {
	type args struct {
		userID string
		data   models.TextData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.TextData{
					Data:         "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.TextData{
					Data:         "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.UpdateTextData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.UpdateTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_UpdateBinaryData - тесты для UpdateBinaryData.
func TestKeeperService_UpdateBinaryData(t *testing.T) {
	type args struct {
		userID string
		data   models.BinaryData
	}
	tests := []struct {
		name    string
		s       *KeeperService
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.BinaryData{
					Data:         []byte("test"),
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.BinaryData{
					Data:         []byte("test"),
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.UpdateBinaryData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.UpdateBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestKeeperService_UpdateBankCardData - тесты для UpdateBankCardData.
func TestKeeperService_UpdateBankCardData(t *testing.T) {
	type args struct {
		userID string
		data   models.BankCardData
	}
	tests := []struct {
		name    string
		s       *KeeperService
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				userID: "test",
				data: models.BankCardData{
					Number:       "test",
					ValidThru:    "test",
					CVV:          "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "Repository error",
			args: args{
				userID: "fail",
				data: models.BankCardData{
					Number:       "test",
					ValidThru:    "test",
					CVV:          "test",
					Description:  "test",
					LastModified: 0,
				},
			},
			wantErr: true,
		},
	}

	m := mock_service.NewMockKeeperRepository()
	svc := New(m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := svc.UpdateBankCardData(tt.args.userID, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("KeeperService.UpdateBankCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
