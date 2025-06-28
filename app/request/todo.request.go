package request

import (
	"github.com/rizama/go-fiber-starter/app/dao"
	"github.com/rizama/go-fiber-starter/app/model"
)

type CreateTodoRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Status string `json:"status" validate:"required,oneof=pending in_progress completed"`
}

type UpdateTodoRequest struct {
	Title  string `json:"title" validate:"omitempty,min=3,max=100"`
	Status string `json:"status" validate:"required,oneof=pending in_progress completed"`
}

func (r *CreateTodoRequest) ToModel() *model.TodoModel {
	return &model.TodoModel{
		Todo: dao.Todo{
			Title:  r.Title,
			Status: r.Status,
		},
	}
}

func (r *UpdateTodoRequest) ToModel() *model.TodoModel {
	return &model.TodoModel{
		Todo: dao.Todo{
			Title:  r.Title,
			Status: r.Status,
		},
	}
}
