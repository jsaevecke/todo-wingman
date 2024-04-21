package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jsaevecke/todo-wingman/task"
)

type Model struct {
	list list.Model
	err  error
}

func (model *Model) initList(width, height int) {
	model.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)

	model.list.Title = "Tasks"
	model.list.SetItems([]list.Item{
		task.New(
			"Learn Golang",
			task.WithDescription("Learn the Go programming language"),
			task.WithTags([]string{"programming", "go"}),
			task.WithStatus(task.TODO),
		),
	})
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
		model.initList(msg.Width, msg.Height)
	}

	var command tea.Cmd

	model.list, command = model.list.Update(msg)

	return model, command
}

func (model Model) View() string {
	return model.list.View()
}

func main() {
	program := tea.NewProgram(Model{})

	_, err := program.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
