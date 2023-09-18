package service

import (
	todo "github.com/WORUS/Golang-CRUD"
	"github.com/WORUS/Golang-CRUD/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)

}

func (s *TodoListService) GetALL(userId int) ([]todo.TodoList, error) {
	return s.repo.GetALL(userId)

}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)

}
