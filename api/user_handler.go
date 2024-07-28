package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"time-tracker/entity"
	"time-tracker/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var userPassport entity.UserPassport

	err := json.NewDecoder(r.Body).Decode(&userPassport)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = entity.Validation(userPassport.PassportNumber)
	if err != nil {
		sendError(ctx, w, err)
		return
	}

	err = u.service.CreateUser(ctx, userPassport)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

func (u *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filters := entity.UserFilter{}
	var err error

	name := r.URL.Query().Get("name")
	if name != "" {
		filters.Name = name
	}

	surname := r.URL.Query().Get("surname")
	if surname != "" {
		filters.Surname = surname
	}

	address := r.URL.Query().Get("address")
	if address != "" {
		filters.Address = address
	}
	date := r.URL.Query().Get("date")
	if date != "" {
		filters.CreatedAt, err = time.Parse("2006-01-02", date)
		if err != nil {
			sendError(ctx, w, err)
			return
		}
	}
	users, err := u.service.Users(ctx, filters)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = sendJson(w, users)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return
	}

	var user entity.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = u.service.UpdateUser(ctx, id, user)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idR := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idR, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return

	}

	err = u.service.DeleteUser(ctx, id)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

func (u *UserHandler) UserByPassport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	myUrl, err := url.Parse(r.RequestURI)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	params, _ := url.ParseQuery(myUrl.RawQuery)

	passportSeriesP := params.Get("passportSerie")
	passportNumberP := params.Get("passportNumber")

	passportSerie, err := strconv.ParseInt(passportSeriesP, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	passportNumber, err := strconv.ParseInt(passportNumberP, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	user, err := u.service.UserByPassport(ctx, passportSerie, passportNumber)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = sendJson(w, user)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}
