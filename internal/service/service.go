package service

import (
	"github.com/HjruRvlq3/cli-to-do-list/internal/storage"
	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
)

func AddTask(newTaskText string) error {
	newTask := task.New(newTaskText)

	err := storage.WriteTask(newTask)
	if err != nil {
		return err
	}

	return nil
}

/*
	type TaskReposetory interface {
		WriteTask(task.Task) error
	}

	type Service struct {
		repo TaskReposetory
	}

	func New(repo TaskReposetory) *Service {
		return &Service{
			repo: repo,
		}
	}

	func (s *Service) AddTask(newTaskText string) error {
		newTask := task.New(newTaskText)
		return s.repo.WriteTask(newTask)
	}
*/
