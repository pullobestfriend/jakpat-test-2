package service

import (
	"jakpat-test-2/entity"
	"jakpat-test-2/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(input entity.Users) (int, error) {
	userID, err := s.repo.CreateUser(input)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *UserService) GetUserByIdAndStatus(id int, status int) (entity.Users, error) {
	user, err := s.repo.GetUserByIdAndStatus(id, status)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByNameAndPassword(name, password string) (entity.Users, error) {
	user, err := s.repo.GetUserByNameAndPassword(name, password)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}
