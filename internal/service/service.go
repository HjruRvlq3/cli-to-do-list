package service

import (
	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
	"github.com/HjruRvlq3/cli-to-do-list/internal/taskBase"
)

func Create(newTaskText string) error {
	newTask := task.New(newTaskText)

	err := taskBase.WriteTask(newTask)
	if err != nil {
		return err
	}

	return nil
}
