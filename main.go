package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jsaevecke/todo-wingman/task"
)

type Model struct {
	isInitialized bool
	focusedList   uint
	lists         []list.Model
	err           error
}

func (model *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)

	model.lists = []list.Model{
		defaultList,
		defaultList,
		defaultList,
	}

	model.lists[task.TODO].Title = "To Do"
	model.lists[task.TODO].SetItems([]list.Item{
		task.New(
			"Learn Golang",
			task.WithDescription("Learn the Go programming language"),
			task.WithTags([]string{"programming", "go"}),
			task.WithStatus(task.TODO),
		),
	})

	model.lists[task.INPROGRESS].Title = "In Progress"
	model.lists[task.DONE].Title = "Done"
}

func New() *Model {
	return &Model{}
}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !model.isInitialized {
			model.initLists(msg.Width, msg.Height)

			model.isInitialized = true
		}
	}

	var command tea.Cmd

	model.lists[model.focusedList], command = model.lists[model.focusedList].Update(msg)

	return model, command
}

func (model Model) View() string {
	if model.isInitialized {

		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			model.lists[task.TODO].View(),
			model.lists[task.INPROGRESS].View(),
			model.lists[task.DONE].View(),
		)
	}

	return "Initializing..."
}

func main() {
	program := tea.NewProgram(Model{})

	_, err := program.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
