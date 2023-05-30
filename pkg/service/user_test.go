package service

import (
	"jakpat-test-2/entity"
	mock_service "jakpat-test-2/mock/service"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser(t *testing.T) {
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
					ID: 1,
				},
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().CreateUser(entity.Users{ID: 1}).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				input: entity.Users{
					ID: 1,
				},
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().CreateUser(entity.Users{ID: 1}).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockRepo)
			s := NewUserService(mockRepo)
			got, err := s.CreateUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByIdAndStatus(t *testing.T) {
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
			mockRepo := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockRepo)
			s := NewUserService(mockRepo)
			got, err := s.GetUserByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByNameAndPassword(t *testing.T) {
	type args struct {
		name     string
		password string
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
				name:     "asd",
				password: "qwe",
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByNameAndPassword("asd", "qwe").Return(entity.Users{}, assert.AnError)
			},
			want:    entity.Users{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name:     "asd",
				password: "qwe",
			},
			mock: func(m *mock_service.MockUser) {
				m.EXPECT().GetUserByNameAndPassword("asd", "qwe").Return(entity.Users{ID: 1}, nil)
			},
			want:    entity.Users{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockUser(mockCtrl)
			tt.mock(mockRepo)
			s := NewUserService(mockRepo)
			got, err := s.GetUserByNameAndPassword(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByNameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByNameAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
