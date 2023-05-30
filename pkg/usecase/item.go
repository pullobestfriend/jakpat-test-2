package usecase

import (
	"errors"
	"jakpat-test-2/entity"
	"time"
)

const (
	statusItemActive = 1

	roleSeller = 1
	roleBuyer  = 2

	orderExpiryTimeInDays = 3
)

var (
	errUserNotSeller           = errors.New("user not seller")
	errUserNotBuyer            = errors.New("user not buyer")
	errUserNotEligibleForOrder = errors.New("user not eligible to view this order")
	errItemAndSellerMismatch   = errors.New("this item is not in this seller inventory")
	errWrongOrderStatusUpdate  = errors.New("cannot update order status to this status")

	orderStatusMap = map[int]string{
		1: "waiting",
		2: "on process",
		3: "shipping",
		4: "delivered",
		5: "expired",
	}
)

func (u *Usecase) AddItem(user entity.Users, input entity.Items) (int, error) {
	if user.Role != roleSeller {
		return 0, errUserNotSeller
	}

	input.SellerID = user.ID
	itemID, err := u.services.AddItem(input)
	if err != nil {
		return 0, err
	}
	return itemID, nil
}

func (u *Usecase) GetItemByIdAndStatus(id int, status int) (entity.Items, error) {
	item, err := u.services.GetItemByIdAndStatus(id, status)
	if err != nil {
		return entity.Items{}, err
	}
	return item, nil
}

func (u *Usecase) UpdateItemById(user entity.Users, id int, input entity.Items) error {
	if user.Role != roleSeller {
		return errUserNotSeller
	}

	item, err := u.services.GetItemByIdAndStatus(id, statusItemActive)
	if err != nil {
		return err
	}

	if item.SellerID != user.ID {
		return errItemAndSellerMismatch
	}

	err = u.services.UpdateItemById(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) GetItemsBySellerIdAndStatus(user entity.Users, sellerID int, status int) ([]entity.Items, error) {
	if user.Role != roleSeller {
		return nil, errUserNotSeller
	}

	if user.ID != sellerID {
		return nil, errItemAndSellerMismatch
	}

	items, err := u.services.GetItemsBySellerIdAndStatus(sellerID, status)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (u *Usecase) CreateOrder(user entity.Users, itemID int) (int, error) {
	if user.Role != roleBuyer {
		return 0, errUserNotBuyer
	}

	orderID, err := u.services.CreateOrder(entity.Order{
		ItemID:      itemID,
		BuyerID:     user.ID,
		Status:      1,
		ExpiredDate: time.Now().AddDate(0, 0, orderExpiryTimeInDays),
	})
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (u *Usecase) GetOrderById(user entity.Users, id int) (entity.Order, error) {
	order, err := u.services.GetOrderById(id)
	if err != nil {
		return entity.Order{}, err
	}

	order.StatusName = orderStatusMap[order.Status]
	if user.ID == order.BuyerID || user.ID == order.SellerID {
		return order, nil
	} else {
		return entity.Order{}, errUserNotEligibleForOrder
	}
}

func (u *Usecase) UpdateOrderStatusByIdAndStatus(user entity.Users, id int, input int) error {
	order, err := u.services.GetOrderById(id)
	if err != nil {
		return err
	}

	if order.SellerID != user.ID {
		return errUserNotEligibleForOrder
	}

	if input-order.Status != 1 {
		return errWrongOrderStatusUpdate
	}

	order.Status = input
	err = u.services.UpdateOrderById(id, order)
	if err != nil {
		return err
	}

	return nil
}
