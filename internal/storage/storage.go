package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
)

const path = "data/data.json"

type TaskRepository interface {
	WriteTask(task.Task) error
	PrintList() error
	Done(int) error
}

type JSONRepository struct{}

func (JSONRepository) WriteTask(newTask task.Task) error {

	var (
		tasks []task.Task
		err   error
	)

	tasks, err = parsList()
	if err != nil {
		return err
	}

	newIndex := createIndex(&tasks)
	newTask.ID = newIndex

	tasks = append(tasks, newTask)

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (JSONRepository) PrintList() error {
	tasks, err := parsList()
	if err != nil {
		return err
	}
	for _, i := range tasks {
		fmt.Printf("%d. %s ", i.ID, i.Text)
		if i.Done {
			fmt.Printf("[✓]\n")
		} else {
			fmt.Printf("[ ]\n")
		}
	}
	return nil
}

func (JSONRepository) Done(number int) error {
	tasks, err := parsList()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == number {
			tasks[i].Done = true
			return nil
		}
	}
	err = errors.New("задачи с таким номером нет")
	return err
}

func parsList() ([]task.Task, error) {
	var tasks []task.Task

	jsonData, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []task.Task{}
		} else {
			return nil, err
		}
	} else {
		if len(jsonData) > 0 {
			err := json.Unmarshal(jsonData, &tasks)
			if err != nil {
				return nil, err
			}
		} else {
			tasks = []task.Task{}
		}
	}
	return tasks, nil
}

func createIndex(tasks *[]task.Task) int {
	var newIndex int

	if len(*tasks) == 0 {
		return 1
	}

	for _, i := range *tasks {
		if i.ID > newIndex {
			newIndex = i.ID
		}
	}

	return newIndex + 1
}
