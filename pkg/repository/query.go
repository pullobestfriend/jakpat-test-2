package repository

import (
	"jakpat-test-2/entity"

	"github.com/jmoiron/sqlx"
)

type DBPostgres struct {
	db *sqlx.DB
}

func NewDBPostgres(db *sqlx.DB) *DBPostgres {
	return &DBPostgres{db: db}
}

func (r *DBPostgres) CreateUser(input entity.Users) (int, error) {
	query := "INSERT INTO users (name, password, role, status) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.db.Get(&id, query, input.Name, input.Password, input.Role, input.Status)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DBPostgres) GetUserByIdAndStatus(id int, status int) (entity.Users, error) {
	var user entity.Users
	query := "SELECT id, name, role, join_date FROM users WHERE id = $1 AND status = $2"
	err := r.db.Get(&user, query, id, status)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r *DBPostgres) GetUserByNameAndPassword(name, password string) (entity.Users, error) {
	var user entity.Users
	query := "SELECT id, name, role, join_date, status FROM users WHERE name = $1 AND password = $2 AND status = 1"
	err := r.db.Get(&user, query, name, password)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (r *DBPostgres) AddItem(input entity.Items) (int, error) {
	query := "INSERT INTO items (seller_id, name, stock, status) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.db.Get(&id, query, input.SellerID, input.Name, input.Stock, input.Status)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DBPostgres) GetItemByIdAndStatus(id int, status int) (entity.Items, error) {
	var item entity.Items
	query := "SELECT id, seller_id, name, stock FROM items WHERE id = $1 AND status = $2"
	err := r.db.Get(&item, query, id, status)
	if err != nil {
		return entity.Items{}, err
	}
	return item, nil
}

func (r *DBPostgres) UpdateItemById(id int, input entity.Items) error {
	query := "UPDATE items SET seller_id = $1, name = $2, stock = $3, status = $4 WHERE id = $5"
	_, err := r.db.Exec(query, input.SellerID, input.Name, input.Stock, input.Status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DBPostgres) GetItemsBySellerIdAndStatus(sellerID int, status int) ([]entity.Items, error) {
	var items []entity.Items
	query := "SELECT id, seller_id, name, stock, status FROM items WHERE seller_id = $1 AND status = $2"
	err := r.db.Select(&items, query, sellerID, status)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *DBPostgres) CreateOrder(input entity.Order) (int, error) {
	query := "INSERT INTO orders (item_id, buyer_id, status, expired_date) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.db.Get(&id, query, input.ItemID, input.BuyerID, input.Status, input.ExpiredDate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DBPostgres) GetOrderById(id int) (entity.Order, error) {
	var order entity.Order
	query := `
	SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
	FROM orders
	JOIN items on orders.item_id = items.id
	WHERE orders.id = $1`
	err := r.db.Get(&order, query, id)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (r *DBPostgres) GetOrdersBySellerId(sellerID int) ([]entity.Order, error) {
	var orders []entity.Order
	query := `
	SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
	FROM orders
	JOIN items on orders.item_id = items.id
	WHERE items.seller_id = $1`
	err := r.db.Select(&orders, query, sellerID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *DBPostgres) UpdateOrderById(id int, input entity.Order) error {
	query := "UPDATE orders SET status = $1, last_updated = NOW() WHERE id = $2"
	_, err := r.db.Exec(query, input.Status, id)
	if err != nil {
		return err
	}
	return nil
}
