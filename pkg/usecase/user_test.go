package usecase

import (
	"context"
	"jakpat-test-2/entity"
	mock_service "jakpat-test-2/mock/service"
	"jakpat-test-2/pkg/service"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_CreateUser(t *testing.T) {
	type args struct {
		input entity.Users
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockUser)
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				input: entity.Users{
					Name:     "asd",
					Password: "qwe",
					Role:     1,
					Status:   1,
				},
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().CreateUser(gomock.Any()).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				input: entity.Users{
					Name:     "asd",
					Password: "qwe",
					Role:     1,
					Status:   1,
				},
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().CreateUser(gomock.Any()).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				User: mockService,
			}, "hash_salt", []byte{}, 10)
			got, err := u.CreateUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetUserByIdAndStatus(t *testing.T) {
	type args struct {
		id     int
		status int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockUser)
		args    args
		want    entity.Users
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				id:     1,
				status: 1,
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByIdAndStatus(int(1), int(1)).Return(entity.Users{}, assert.AnError)
			},
			want:    entity.Users{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				id:     1,
				status: 1,
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByIdAndStatus(int(1), int(1)).Return(entity.Users{ID: 1}, nil)
			},
			want:    entity.Users{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				User: mockService,
			}, "hash_salt", []byte{}, 10)
			got, err := u.GetUserByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetUserByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetUserByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetUserByNameAndPassword(t *testing.T) {
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockUser)
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				name:     "asd",
				password: "qwe",
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByNameAndPassword("asd", gomock.Any()).Return(entity.Users{}, assert.AnError)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name:     "asd",
				password: "qwe",
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByNameAndPassword("asd", gomock.Any()).Return(entity.Users{
					ID: 1,
				}, nil)
			},
			want:    gomock.Any().String(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				User: mockService,
			}, "hash_salt", []byte{}, 10)
			_, err := u.GetUserByNameAndPassword(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetUserByNameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// TODO add want matching unit test
			// if got != tt.want {
			// 	t.Errorf("Usecase.GetUserByNameAndPassword() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUsecase_ParseToken(t *testing.T) {
	type args struct {
		ctx         context.Context
		accessToken string
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockUser)
		args    args
		want    *entity.Users
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				ctx:         context.Background(),
				accessToken: "asd",
			},
			mock:    func(m *mock_service.MockUser) {},
			want:    &entity.Users{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				User: mockService,
			}, "hash_salt", []byte{10}, 10)
			_, err := u.ParseToken(tt.args.ctx, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// TODO add want matching unit test
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Usecase.ParseToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
