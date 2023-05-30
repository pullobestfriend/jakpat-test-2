package service

import (
	"jakpat-test-2/entity"
	"jakpat-test-2/pkg/repository"
)

type User interface {
	CreateUser(input entity.Users) (int, error)
	GetUserByIdAndStatus(id int, status int) (entity.Users, error)
	GetUserByNameAndPassword(name, password string) (entity.Users, error)
}

type Item interface {
	AddItem(input entity.Items) (int, error)
	GetItemByIdAndStatus(id int, status int) (entity.Items, error)
	UpdateItemById(id int, input entity.Items) error
	GetItemsBySellerIdAndStatus(sellerID int, status int) ([]entity.Items, error)
	CreateOrder(input entity.Oders) (int, error)
	GetOrderById(id int) (entity.Oders, error)
	UpdateOrderById(id int, input entity.Oders) error
}

type Service struct {
	User
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Item: NewItemService(repos.Item),
	}
}
