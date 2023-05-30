package service

import (
	"errors"
	"jakpat-test-2/entity"
	"jakpat-test-2/pkg/repository"
	"time"
)

const (
	orderStatusWaiting = 1
	orderStatusExpired = 5

	statusItemActive = 1
)

var (
	errItemOutOfStock = errors.New("item out of stock")

	// unit test purpose
	now = time.Now
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) AddItem(input entity.Items) (int, error) {
	itemID, err := s.repo.AddItem(input)
	if err != nil {
		return 0, err
	}
	return itemID, nil
}

func (s *ItemService) GetItemByIdAndStatus(id int, status int) (entity.Items, error) {
	item, err := s.repo.GetItemByIdAndStatus(id, status)
	if err != nil {
		return entity.Items{}, err
	}
	return item, nil
}

func (s *ItemService) UpdateItemById(id int, input entity.Items) error {
	err := s.repo.UpdateItemById(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *ItemService) GetItemsBySellerIdAndStatus(sellerID int, status int) ([]entity.Items, error) {
	items, err := s.repo.GetItemsBySellerIdAndStatus(sellerID, status)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemService) CreateOrder(input entity.Oders) (int, error) {
	item, err := s.repo.GetItemByIdAndStatus(input.ItemID, statusItemActive)
	if err != nil {
		return 0, err
	}

	if item.Stock <= 0 {
		return 0, errItemOutOfStock
	}

	// add tx
	orderID, err := s.repo.CreateOrder(input)
	if err != nil {
		return 0, err
	}

	item.Stock--
	err = s.repo.UpdateItemById(item.ID, item)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (s *ItemService) GetOrderById(id int) (entity.Oders, error) {
	order, err := s.repo.GetOrderById(id)
	if err != nil {
		return entity.Oders{}, err
	}

	if now().After(order.ExpiredDate) && order.Status ==  orderStatusWaiting{
		order.Status = orderStatusExpired
		err = s.repo.UpdateOrderById(id, order)
		if err != nil {
			return entity.Oders{}, err
		}
	}

	return order, nil
}

func (s *ItemService) UpdateOrderById(id int, input entity.Oders) error {
	err := s.repo.UpdateOrderById(id, input)
	if err != nil {
		return err
	}
	return nil
}
