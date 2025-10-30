package helloworld

import (
	"context"

	user "github.com/andrewronscki/lib-golang-teste/internal/user/domain"
	"github.com/andrewronscki/lib-golang-teste/pkg/commons/cqrs"
)

type Command struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CommandHandler struct{}

func (h *CommandHandler) Handle(ctx context.Context, command *Command) (*Model, error) {
	userCreated := user.NewUser(command.FirstName, command.LastName)

	model := &Model{}
	userCreated.Marshal(model)

	return model, nil
}

func NewCommandHandler() cqrs.ICommandHandler[*Command, *Model] {
	return &CommandHandler{}
}
