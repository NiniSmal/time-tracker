package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"time-tracker/entity"
)

type TaskRepo struct {
	conn *pgx.Conn
}

func NewTaskRepo(conn *pgx.Conn) *TaskRepo {
	return &TaskRepo{
		conn: conn,
	}
}

func (t *TaskRepo) CreateTask(ctx context.Context, task entity.Task) error {
	query := "INSERT INTO tasks ( name, status, created_at, finished_at, lead_time, user_id ) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := t.conn.Exec(ctx, query, task.Name, task.Status, task.CreatedAt, task.FinishedAt, task.LeadTime, task.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepo) UpdateStatus(ctx context.Context, id int64, task entity.TasksFilter) error {
	query := "UPDATE tasks SET status = $1, finished_at = $2, lead_time =$3  WHERE id = $4"

	_, err := t.conn.Exec(ctx, query, task.Status, task.FinishedAt, task.LeadTime, id)
	if err != nil {
		return err
	}
	return nil
}
func (t *TaskRepo) TaskByID(ctx context.Context, id int64) (entity.Task, error) {
	query := "SELECT id, name, status, created_at, finished_at, lead_time, user_id FROM tasks WHERE id = $1"

	var task entity.Task

	err := t.conn.QueryRow(ctx, query, id).Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt, &task.FinishedAt, &task.LeadTime, &task.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Task{}, entity.ErrNotFound
		} else {
			return entity.Task{}, err
		}
	}
	return task, nil
}

func (t *TaskRepo) UserTasks(ctx context.Context, task entity.TasksFilter) ([]entity.Task, error) {
	query := "SELECT id,  lead_time, user_id FROM tasks WHERE user_id = $1 AND (created_at BETWEEN $2 AND $3) AND (finished_at BETWEEN $2 AND $3)  ORDER BY lead_time DESC "
	var tasks []entity.Task

	rows, err := t.conn.Query(ctx, query, task.UserID, task.CreatedAt, task.FinishedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		err = rows.Scan(&task.ID, &task.LeadTime, &task.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}
