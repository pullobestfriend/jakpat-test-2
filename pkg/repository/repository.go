package repository

import (
	"jakpat-test-2/entity"

	"github.com/jmoiron/sqlx"
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
	CreateOrder(input entity.Order) (int, error)
	GetOrderById(id int) (entity.Order, error)
	UpdateOrderById(id int, input entity.Order) error
}

type Repository struct {
	User
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewDBPostgres(db),
		Item: NewDBPostgres(db),
	}
}
