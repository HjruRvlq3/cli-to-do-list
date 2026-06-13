package service

import (
	"fmt"

	"github.com/HjruRvlq3/cli-to-do-list/internal/storage"
	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
	clearscreen "github.com/HjruRvlq3/cli-to-do-list/pkg/clearScreen"
)

type Service struct {
	repo storage.TaskRepository
}

func New(repo storage.TaskRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddTask(newTaskText string) error {
	newTask := task.New(newTaskText)

	err := s.repo.WriteTask(newTask)

	if err != nil {
		return err
	}

	return nil
}

func PrintList(srv *Service) {
	clearscreen.ClearScreen()

	srv.repo.PrintList()

	fmt.Printf("\n")
}

func (s *Service) Dune(number int) error {
	err := s.repo.Done(number)
	return err
}
