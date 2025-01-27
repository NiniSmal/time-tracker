package storage

import (
	"context"
	"errors"
	sq "github.com/elgris/sqrl"
	"github.com/jackc/pgx/v5"
	"time-tracker/entity"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(c *pgx.Conn) *UserRepo {
	return &UserRepo{
		conn: c,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	query := "INSERT INTO users( passport_num, passport_series, surname, name, patronymic, address, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.conn.Exec(ctx, query, user.PassportNum, user.PassportSeries, user.Surname, user.Name, user.Patronymic, user.Address, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Users(ctx context.Context, filter entity.UserFilter) (users []entity.User, err error) {
	builder := sq.Select("id", "passport_num", "passport_series", "surname", "name", "patronymic", "address", "created_at").
		From("users")
	if filter.Name != "" {
		builder = builder.Where(sq.Eq{"name": filter.Name})
	}
	if filter.Surname != "" {
		builder = builder.Where(sq.Eq{"surname": filter.Surname})
	}
	if filter.Patronymic != "" {
		builder = builder.Where(sq.Eq{"patronymic": filter.Patronymic})
	}
	if filter.Address != "" {
		builder = builder.Where(sq.Eq{"address": filter.Address})
	}
	if !filter.CreatedAt.IsZero() {
		builder = builder.Where(sq.Eq{"created_at": filter.CreatedAt})
	}
	builder = builder.Limit(5)

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var user entity.User
	_, err = pgx.ForEachRow(rows, []any{&user.ID, &user.PassportNum, &user.PassportSeries, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt}, func() error {
		users = append(users, user)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int64, user entity.User) error {
	query := "UPDATE users SET name = $1, surname=$2, patronymic =$3, address =$4 WHERE id = $5"
	_, err := r.conn.Exec(ctx, query, user.Name, user.Surname, user.Patronymic, user.Address, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int64) error {
	query := "DELETE  FROM users WHERE id = $1"
	_, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UserByPassport(ctx context.Context, passportSerie, passportNumber int64) (entity.User, error) {
	query := "SELECT surname, name, patronymic, address FROM users WHERE passport_series = $1 AND passport_num = $2"
	var user entity.User

	err := r.conn.QueryRow(ctx, query, passportSerie, passportNumber).Scan(&user.Surname, &user.Name, &user.Patronymic, &user.Address)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, entity.ErrBadRequest
		}
		return entity.User{}, err

	}
	return user, nil
}
