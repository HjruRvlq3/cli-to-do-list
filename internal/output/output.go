package output

import (
	"fmt"

	"github.com/HjruRvlq3/cli-to-do-list/internal/task"
)

func SuccessfulСreation(newTask task.Task) {
	fmt.Print(newTask)
}

func ConsoleClean() {
	fmt.Print("\033[H\033[2J")
}
