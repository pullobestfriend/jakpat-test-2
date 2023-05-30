package entity

import "time"

type Users struct {
	ID       int       `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password string    `json:"password" db:"password"` // implement encryption
	Role     int       `json:"role" db:"role"`         // 1 for seller, 2 for buyer
	JoinDate time.Time `json:"joinDate" db:"join_date"`
	Status   int       `json:"status" db:"status"` // set 0 for soft delete
}

type Items struct {
	ID       int    `json:"id" db:"id"`
	SellerID int    `json:"sellerId" db:"seller_id"`
	Name     string `json:"name" db:"name"`
	Stock    int    `json:"stock" db:"stock"`
	Status   int    `json:"status" db:"status"` // set 0 for soft delete
}

type Oders struct {
	ID      int `json:"id" db:"id"`
	ItemID  int `json:"itemId" db:"item_id"`
	BuyerID int `json:"buyerId" db:"buyer_id"`
	SellerID int `json:"sellerID" db:"seller_id"`

	// 1 for waiting
	// 2 for on process
	// 3 for shipping
	// 4 for delivered
	// 5 for expired
	Status     int    `json:"status" db:"status"`
	StatusName string `json:"statusName"`

	ExpiredDate time.Time `json:"expiredDate" db:"expired_date"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	LastUpdated time.Time `json:"lastUpdated" db:"last_updated"`
}
