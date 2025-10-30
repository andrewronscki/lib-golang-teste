package helloworld

import (
	"context"
	"fmt"

	"github.com/andrewronscki/lib-golang-teste/pkg/cqrs"
)

type Command struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CommandHandler struct {
}

func (h *CommandHandler) Handle(ctx context.Context, command *Command) (*Model, error) {
	name := fmt.Sprintf("%s %s", command.FirstName, command.LastName)

	model := &Model{
		Name: name,
	}

	return model, nil
}

func NewCommandHandler() cqrs.ICommandHandler[*Command, *Model] {
	return &CommandHandler{}
}
