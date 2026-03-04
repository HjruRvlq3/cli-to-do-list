package input

import (
	"fmt"

	"github.com/HjruRvlq3/cli-to-do-list/internal/output"
	"github.com/HjruRvlq3/cli-to-do-list/internal/service"
)

func MovManager() {
	var mov string

	fmt.Scan(&mov)

	switch mov {
	case "create":
		err := create()
		if err != nil {
			fmt.Println("Ошибка", err)
		} else {
			output.ConsoleClean()
			fmt.Printf("\nЗадача добавлена\n")
		}
	}
}

func create() error {
	output.ConsoleClean()
	fmt.Printf("\nВведите текст задачи в одну строку:\n")

	var newTaskText string
	fmt.Scan(&newTaskText)

	err := service.Create(newTaskText)
	if err != nil {
		return err
	}

	return nil
}
