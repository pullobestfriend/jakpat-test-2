package service

import (
	"jakpat-test-2/entity"
	mock_service "jakpat-test-2/mock/service"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestItemService_AddItem(t *testing.T) {
	type args struct {
		input entity.Items
	}
	tests := []struct {
		name    string
		args    args
		mock    func(m *mock_service.MockItem)
		want    int
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				input: entity.Items{
					ID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().AddItem(entity.Items{ID: 1}).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				input: entity.Items{
					ID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().AddItem(entity.Items{ID: 1}).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			got, err := s.AddItem(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ItemService.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_GetItemByIdAndStatus(t *testing.T) {
	type args struct {
		id     int
		status int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    entity.Items
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				id:     1,
				status: 1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), int(1)).Return(entity.Items{}, assert.AnError)
			},
			want:    entity.Items{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				id:     1,
				status: 1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), int(1)).Return(entity.Items{ID: 1}, nil)
			},
			want:    entity.Items{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			got, err := s.GetItemByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.GetItemByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemService.GetItemByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_UpdateItemById(t *testing.T) {
	type args struct {
		id    int
		input entity.Items
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				id:    1,
				input: entity.Items{Stock: 2},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().UpdateItemById(int(1), entity.Items{Stock: 2}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				id:    1,
				input: entity.Items{Stock: 2},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().UpdateItemById(int(1), entity.Items{Stock: 2}).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			if err := s.UpdateItemById(tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ItemService.UpdateItemById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemService_GetItemsBySellerIdAndStatus(t *testing.T) {
	type args struct {
		sellerID int
		status   int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    []entity.Items
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				sellerID: 1,
				status:   1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemsBySellerIdAndStatus(int(1), int(1)).Return(nil, assert.AnError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				sellerID: 1,
				status:   1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemsBySellerIdAndStatus(int(1), int(1)).Return([]entity.Items{
					{
						ID: 1,
					},
				}, nil)
			},
			want: []entity.Items{
				{
					ID: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			got, err := s.GetItemsBySellerIdAndStatus(tt.args.sellerID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.GetItemsBySellerIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemService.GetItemsBySellerIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_CreateOrder(t *testing.T) {
	type args struct {
		input entity.Oders
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "error_item_not_found",
			args: args{
				input: entity.Oders{
					ItemID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), statusItemActive).Return(entity.Items{}, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error_item_no_stock",
			args: args{
				input: entity.Oders{
					ItemID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), statusItemActive).Return(entity.Items{Stock: 0}, nil)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error_create_order",
			args: args{
				input: entity.Oders{
					ItemID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), statusItemActive).Return(entity.Items{Stock: 1}, nil)
				m.EXPECT().CreateOrder(entity.Oders{ItemID: 1}).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error_update_item",
			args: args{
				input: entity.Oders{
					ItemID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), statusItemActive).Return(entity.Items{
					ID:    2,
					Stock: 1,
				}, nil)
				m.EXPECT().CreateOrder(entity.Oders{ItemID: 1}).Return(1, nil)
				m.EXPECT().UpdateItemById(int(2), entity.Items{
					ID:    2,
					Stock: 0,
				}).Return(assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				input: entity.Oders{
					ItemID: 1,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(1), statusItemActive).Return(entity.Items{
					ID:    2,
					Stock: 1,
				}, nil)
				m.EXPECT().CreateOrder(entity.Oders{ItemID: 1}).Return(1, nil)
				m.EXPECT().UpdateItemById(int(2), entity.Items{
					ID:    2,
					Stock: 0,
				}).Return(nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			got, err := s.CreateOrder(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ItemService.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_GetOrderById(t *testing.T) {
	mockTime := time.Date(2023, 1, 1, 1, 1, 1, 1, time.Local)
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    entity.Oders
		wantErr bool
	}{
		{
			name: "error_getting_order",
			args: args{
				id: 1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(1)).Return(entity.Oders{}, assert.AnError)
			},
			want:    entity.Oders{},
			wantErr: true,
		},
		{
			name: "error_update_expired_order",
			args: args{
				id: 1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(1)).Return(entity.Oders{
					ID:          1,
					Status:      orderStatusWaiting,
					ExpiredDate: mockTime,
				}, nil)
				now = func() time.Time {
					return time.Date(2024, 1, 1, 1, 1, 1, 1, time.Local)
				}
				m.EXPECT().UpdateOrderById(int(1), entity.Oders{
					ID:          1,
					Status:      orderStatusExpired,
					ExpiredDate: mockTime,
				}).Return(assert.AnError)
			},
			want:    entity.Oders{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				id: 1,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(1)).Return(entity.Oders{
					ID:          1,
					Status:      orderStatusWaiting,
					ExpiredDate: mockTime,
				}, nil)
				now = func() time.Time {
					return time.Date(2024, 1, 1, 1, 1, 1, 1, time.Local)
				}
				m.EXPECT().UpdateOrderById(int(1), entity.Oders{
					ID:          1,
					Status:      orderStatusExpired,
					ExpiredDate: mockTime,
				}).Return(nil)
			},
			want: entity.Oders{
				ID:          1,
				Status:      orderStatusExpired,
				ExpiredDate: mockTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			got, err := s.GetOrderById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.GetOrderById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemService.GetOrderById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_UpdateOrderById(t *testing.T) {
	type args struct {
		id    int
		input entity.Oders
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				id:    1,
				input: entity.Oders{Status: 2},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().UpdateOrderById(int(1), entity.Oders{Status: 2}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				id:    1,
				input: entity.Oders{Status: 2},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().UpdateOrderById(int(1), entity.Oders{Status: 2}).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockRepo := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockRepo)
			s := NewItemService(mockRepo)
			if err := s.UpdateOrderById(tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ItemService.UpdateOrderById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
