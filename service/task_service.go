package service

import (
	"context"
	"time"
	"time-tracker/entity"
	"time-tracker/storage"
)

type TaskService struct {
	repo *storage.TaskRepo
}

func NewTaskService(r *storage.TaskRepo) *TaskService {
	return &TaskService{
		repo: r,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	task.CreatedAt = time.Now()
	task.FinishedAt.IsZero()
	task.LeadTime = 0
	task.Status = false

	err := s.repo.CreateTask(ctx, task)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (s *TaskService) UpdateStatus(ctx context.Context, id int64, task entity.TasksFilter) (entity.Task, error) {
	taskDB, err := s.repo.TaskByID(ctx, id)
	if err != nil {
		return entity.Task{}, err
	}

	task.FinishedAt = time.Now()
	task.LeadTime = task.FinishedAt.Sub(task.CreatedAt) / time.Minute

	err = s.repo.UpdateStatus(ctx, id, task)
	if err != nil {
		return entity.Task{}, err
	}
	taskDB, err = s.repo.TaskByID(ctx, id)
	if err != nil {
		return entity.Task{}, err
	}

	return taskDB, nil
}

func (s *TaskService) TasksTimeByUserID(ctx context.Context, task entity.TasksFilter) ([]entity.Task, error) {
	tasks, err := s.repo.UserTasks(ctx, task)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
