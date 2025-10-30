package helloworld

import (
	"context"

	helloworld "github.com/andrewronscki/lib-golang-teste/internal/user/features/hello-world"
	"github.com/andrewronscki/lib-golang-teste/pkg/cqrs"
)

type HandlerDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func Handler(dto *HandlerDTO) (*helloworld.Model, error) {
	ctx := context.Background()
	command := &helloworld.Command{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	if model, err := cqrs.Send[*helloworld.Command, *helloworld.Model](ctx, command); err != nil {
		return nil, err
	} else {
		return model, nil
	}
}
