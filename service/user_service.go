package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"time-tracker/entity"
	"time-tracker/storage"
)

type UserService struct {
	repo   *storage.UserRepo
	client *http.Client
	appURL string
}

func NewUserService(r *storage.UserRepo, client *http.Client, appURL string) *UserService {
	return &UserService{
		repo:   r,
		client: client,
		appURL: appURL,
	}
}

func (s *UserService) CreateUser(ctx context.Context, userPassport entity.UserPassport) error {
	arr := strings.Split(userPassport.PassportNumber, " ")
	PassportSeries := arr[0]
	PassportNumber := arr[1]

	userPassportNum, err := strconv.ParseInt(PassportNumber, 10, 64)
	if err != nil {
		return err
	}
	userPassportSeries, err := strconv.ParseInt(PassportSeries, 10, 64)
	if err != nil {
		return err
	}
	userInfo, err := s.userInfo(ctx, userPassportNum, userPassportSeries)
	if err != nil {
		return err
	}
	user := entity.User{
		PassportNum:    userPassportNum,
		PassportSeries: userPassportSeries,
		Surname:        userInfo.Surname,
		Name:           userInfo.Name,
		Patronymic:     userInfo.Patronymic,
		Address:        userInfo.Address,
		CreatedAt:      time.Time{},
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) userInfo(ctx context.Context, userPassportNum, userPassportSeries int64) (entity.UserInfo, error) {
	url := fmt.Sprintf("%s?passportSerie=%d&passportNumber=%d", s.appURL, userPassportSeries, userPassportNum)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entity.UserInfo{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		if resp.StatusCode != http.StatusOK {
			return entity.UserInfo{}, entity.ErrBadRequest
		}
		return entity.UserInfo{}, err
	}

	defer resp.Body.Close()

	var user entity.UserInfo

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return entity.UserInfo{}, err
	}
	return user, nil
}

func (s *UserService) Users(ctx context.Context, filters entity.UserFilter) ([]entity.User, error) {
	users, err := s.repo.Users(ctx, filters)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, user entity.User) error {
	err := s.repo.UpdateUser(ctx, id, user)
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
