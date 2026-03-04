package taskBase

import (
	"encoding/json"
	"os"

	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
)

const path = "data/data.json"

func WriteTask(newTask task.Task) error {

	var tasks []task.Task

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []task.Task{}
		} else {
			return err
		}
	} else {
		if len(data) > 0 {
			err := json.Unmarshal(data, &tasks)
			if err != nil {
				return err
			}
		} else {
			tasks = []task.Task{}
		}
	}

	tasks = append(tasks, newTask)

	data, err = json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
