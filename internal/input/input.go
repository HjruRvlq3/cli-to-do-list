// обработка ввода

package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/HjruRvlq3/cli-to-do-list/internal/output"
	"github.com/HjruRvlq3/cli-to-do-list/internal/service"
	clearscreen "github.com/HjruRvlq3/cli-to-do-list/pkg/clearScreen"
)

func MovManager() {
	var (
		move string
		err  error
	)

	fmt.Printf("Go to-do \n\n")

	for move != "exit" && move != "ex" {

		move, err = readInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if move == "" {
			continue
		}

		switch move {
		case "list", "ls":
			output.PrintList()
		case "add", "a":
			clearscreen.ClearScreen()
			err = addTask()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "help", "h":
			clearscreen.ClearScreen()
			output.Help()
		case "exit", "ex":
			clearscreen.ClearScreen()
			output.Exit()
			return
		default:
			clearscreen.ClearScreen()
			output.Err(move)
		}
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text := strings.TrimSpace(strings.ToLower(in))
	return text, nil
}

func addTask() error {
	var (
		newTaskText string
		err         error
	)

	clearscreen.ClearScreen()
	fmt.Printf("\nВведите текст задачи в одну строку:\n\n")

	newTaskText, err = readInput()
	if err != nil {
		return err
	}

	err = service.AddTask(newTaskText)
	if err != nil {
		return err
	}
	output.Success(newTaskText, true)

	return nil
}
