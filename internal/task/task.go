package task

func New(newTaskText string) Task {
	newTask := Task{0, newTaskText, false}
	return newTask
}

type Task struct {
	ID   int
	Text string
	Done bool
}
