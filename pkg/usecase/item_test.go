package usecase

import (
	"jakpat-test-2/entity"
	mock_service "jakpat-test-2/mock/service"
	"jakpat-test-2/pkg/service"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_AddItem(t *testing.T) {
	type args struct {
		user  entity.Users
		input entity.Items
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "error_not_seller",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				input: entity.Items{
					ID: 3,
				},
			},
			mock:    func(m *mock_service.MockItem) {},
			want:    0,
			wantErr: true,
		},
		{
			name: "error_add_item",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				input: entity.Items{
					ID: 3,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().AddItem(entity.Items{
					ID:       3,
					SellerID: 1,
				}).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				input: entity.Items{
					ID: 3,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().AddItem(entity.Items{
					ID:       3,
					SellerID: 1,
				}).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			got, err := u.AddItem(tt.args.user, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetItemByIdAndStatus(t *testing.T) {
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
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			got, err := u.GetItemByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetItemByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetItemByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_UpdateItemById(t *testing.T) {
	type args struct {
		user  entity.Users
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
			name: "error_not_seller",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				id: 3,
				input: entity.Items{
					ID: 3,
				},
			},
			mock:    func(m *mock_service.MockItem) {},
			wantErr: true,
		},
		{
			name: "error_getting_item",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id: 3,
				input: entity.Items{
					ID: 3,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(3), statusItemActive).Return(entity.Items{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "error_item_is_not_sellers",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id: 3,
				input: entity.Items{
					ID: 3,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(3), statusItemActive).Return(entity.Items{
					ID:       3,
					SellerID: 5,
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "error_update_item",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id: 3,
				input: entity.Items{
					ID:    3,
					Stock: 5,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(3), statusItemActive).Return(entity.Items{
					ID:       3,
					SellerID: 1,
				}, nil)
				m.EXPECT().UpdateItemById(int(3), entity.Items{
					ID:    3,
					Stock: 5,
				}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id: 3,
				input: entity.Items{
					ID:    3,
					Stock: 5,
				},
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetItemByIdAndStatus(int(3), statusItemActive).Return(entity.Items{
					ID:       3,
					SellerID: 1,
				}, nil)
				m.EXPECT().UpdateItemById(int(3), entity.Items{
					ID:    3,
					Stock: 5,
				}).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			if err := u.UpdateItemById(tt.args.user, tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.UpdateItemById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_GetItemsBySellerIdAndStatus(t *testing.T) {
	type args struct {
		user     entity.Users
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
			name: "error_not_seller",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				sellerID: 3,
				status:   1,
			},
			mock:    func(m *mock_service.MockItem) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error_wrong_seller",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				sellerID: 3,
				status:   1,
			},
			mock:    func(m *mock_service.MockItem) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error_getting_item",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
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
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
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
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			got, err := u.GetItemsBySellerIdAndStatus(tt.args.user, tt.args.sellerID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetItemsBySellerIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetItemsBySellerIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_CreateOrder(t *testing.T) {
	type args struct {
		user   entity.Users
		itemID int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "error_not_buyer",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				itemID: 5,
			},
			mock:    func(m *mock_service.MockItem) {},
			want:    0,
			wantErr: true,
		},
		{
			name: "error_create_order",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				itemID: 5,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().CreateOrder(gomock.Any()).Return(0, assert.AnError)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				itemID: 5,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().CreateOrder(gomock.Any()).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			got, err := u.CreateOrder(tt.args.user, tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetOrderById(t *testing.T) {
	type args struct {
		user entity.Users
		id   int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		want    entity.Order
		wantErr bool
	}{
		{
			name: "error_getting_order",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id: 5,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{}, assert.AnError)
			},
			want:    entity.Order{},
			wantErr: true,
		},
		{
			name: "error_user_not_eligible",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				id: 5,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 3,
					BuyerID:  4,
				}, nil)
			},
			want:    entity.Order{},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 2,
				},
				id: 5,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 3,
					BuyerID:  1,
					Status:   1,
				}, nil)
			},
			want: entity.Order{
				ID:         5,
				SellerID:   3,
				BuyerID:    1,
				Status:     1,
				StatusName: orderStatusMap[1],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			got, err := u.GetOrderById(tt.args.user, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetOrderById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetOrderById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_UpdateOrderStatusByIdAndStatus(t *testing.T) {
	type args struct {
		user  entity.Users
		id    int
		input int
	}
	tests := []struct {
		name    string
		mock    func(m *mock_service.MockItem)
		args    args
		wantErr bool
	}{
		{
			name: "error_getting_order",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id:    5,
				input: 3,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "error_not_correct_seller",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id:    5,
				input: 3,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 2,
					Status:   1,
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "error_not_correct_status",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id:    5,
				input: 3,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 1,
					Status:   1,
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "error_update",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id:    5,
				input: 3,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 1,
					Status:   2,
				}, nil)
				m.EXPECT().UpdateOrderById(int(5), entity.Order{
					ID:       5,
					SellerID: 1,
					Status:   3,
				}).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				user: entity.Users{
					ID:   1,
					Role: 1,
				},
				id:    5,
				input: 3,
			},
			mock: func(m *mock_service.MockItem) {
				m.EXPECT().GetOrderById(int(5)).Return(entity.Order{
					ID:       5,
					SellerID: 1,
					Status:   2,
				}, nil)
				m.EXPECT().UpdateOrderById(int(5), entity.Order{
					ID:       5,
					SellerID: 1,
					Status:   3,
				}).Return(nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockService := mock_service.NewMockItem(mockCtrl)
			tt.mock(mockService)
			u := NewUsecase(&service.Service{
				Item: mockService,
			}, "", []byte{}, 0)
			if err := u.UpdateOrderStatusByIdAndStatus(tt.args.user, tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.UpdateOrderStatusByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
