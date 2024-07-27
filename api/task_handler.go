package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"time-tracker/entity"
	"time-tracker/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: s,
	}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	taskStart, err := t.service.CreateTask(ctx, task)
	if err != nil {
		sendError(ctx, w, err)
		return
	}

	err = sendJson(w, taskStart)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

func (t *TaskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idP := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idP, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return
	}

	var task entity.TasksFilter
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendError(ctx, w, err)
		return
	}

	taskDB, err := t.service.UpdateStatus(ctx, id, task)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = sendJson(w, taskDB)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}

// Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
func (t *TaskHandler) TaskTimeByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idP := r.URL.Query().Get("user_id")
	dateStart := r.URL.Query().Get("start")
	dateEnd := r.URL.Query().Get("end")

	var task entity.TasksFilter
	var err error

	task.UserID, err = strconv.ParseInt(idP, 10, 64)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	task.CreatedAt, err = time.Parse("2006-01-02", dateStart)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	task.FinishedAt, err = time.Parse("2006-01-02", dateEnd)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	taskHours, err := t.service.TasksTimeByUserID(ctx, task)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
	err = sendJson(w, taskHours)
	if err != nil {
		sendError(ctx, w, err)
		return
	}
}
