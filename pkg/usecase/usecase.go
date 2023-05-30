package usecase

import (
	"context"
	"jakpat-test-2/entity"
	"jakpat-test-2/pkg/service"
	"time"
)

type User interface {
	CreateUser(input entity.Users) (int, error)
	GetUserByIdAndStatus(id int, status int) (entity.Users, error)
	GetUserByNameAndPassword(name, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entity.Users, error)
}

type Item interface {
	AddItem(user entity.Users, input entity.Items) (int, error)
	GetItemByIdAndStatus(id int, status int) (entity.Items, error)
	UpdateItemById(user entity.Users, id int, input entity.Items) error
	GetItemsBySellerIdAndStatus(user entity.Users, sellerID int, status int) ([]entity.Items, error)
	CreateOrder(user entity.Users, itemID int) (int, error)
	GetOrderById(user entity.Users, id int) (entity.Oders, error)
	UpdateOrderStatusByIdAndStatus(user entity.Users, id int, input int) error
}

type Usecase struct {
	services       *service.Service
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewUsecase(services *service.Service,
	hashSalt string,
	signingKey []byte,
	tokenTTL time.Duration) *Usecase {
	return &Usecase{
		services:       services,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}
