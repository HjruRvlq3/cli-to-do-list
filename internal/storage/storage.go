package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
)

const path = "data/data.json"

func WriteTask(newTask task.Task) error {

	var (
		tasks []task.Task
		err   error
	)

	tasks, err = parsList()

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

func PrintList() error {
	tasks, err := parsList()
	if err != nil {
		return err
	}
	for _, i := range tasks {
		fmt.Printf("%d. %s\n", i.ID, i.Text)
	}
	return nil
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
