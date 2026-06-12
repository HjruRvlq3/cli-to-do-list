package output

import (
	"fmt"

	"github.com/HjruRvlq3/cli-to-do-list/internal/storage"
	clearscreen "github.com/HjruRvlq3/cli-to-do-list/pkg/clearScreen"
)

func Exit() {
	fmt.Printf(
		"\n%s\n",
		"Go to-do off",
	)
}

func Err(inputMove string) {
	fmt.Printf("\n%s%s%s\n\n", "команды '", inputMove, "' нет, попробуйте другую")
	Help()
}

func Help() {
	fmt.Printf(
		"%s\n\n%s\n%s\n%s\n%s\n%s\n\n",
		"Go to-do",
		"add / a		- добавить задачу",
		"done / d	- выполнить задачу (выполненная задача будет удалена)",
		"list / ls	- активные задачи",
		"help / h	- информация о приложении",
		"exit / ex	- выход",
	)
}

// taskText - текст задачи | move true добавлена | move false удалена
func Success(taskText string, move bool) {
	switch move {
	case true:
		clearscreen.ClearScreen()
		fmt.Printf("\nЗадача '%s' добавлена\n\n", taskText)
	default:
		fmt.Printf("Задача '%s' удалена", taskText)
	}
}

func PrintList() {
	clearscreen.ClearScreen()
	storage.PrintList()
	fmt.Printf("\n")
}
