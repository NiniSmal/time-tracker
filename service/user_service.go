package service

import (
	"context"
	"strconv"
	"strings"
	"time"
	"time-tracker/entity"
	"time-tracker/storage"
)

type UserService struct {
	repo *storage.UserRepo
}

func NewUserService(r *storage.UserRepo) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) error {
	var err error
	arr := strings.Split(user.PassportNumber, " ")
	PassportSeries := arr[0]
	PassportNumber := arr[1]
	user.PassportNum, err = strconv.ParseInt(PassportNumber, 10, 64)
	if err != nil {
		return err
	}
	user.PassportSeries, err = strconv.ParseInt(PassportSeries, 10, 64)
	if err != nil {
		return err
	}
	user.CreatedAt = time.Now()

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Users(ctx context.Context, filters entity.UserFilter) ([]entity.User, error) {
	users, err := s.repo.Users(ctx, filters)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, passportNumber string, user entity.User) error {
	err := s.repo.UpdateUser(ctx, passportNumber, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserService) UserByPassport(ctx context.Context, passportSerie, passportNumber int64) (entity.User, error) {
	user, err := s.repo.UserByPassport(ctx, passportSerie, passportNumber)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
