package behaviors

import (
	"context"

	"github.com/andrewronscki/lib-golang-teste/pkg/commons/cqrs"
	validation "github.com/andrewronscki/lib-golang-teste/pkg/commons/validation"
)

type ValidationBehavior struct{}

func (b *ValidationBehavior) Handle(ctx context.Context, request any, next cqrs.NextFunc) (any, error) {
	validatable, ok := request.(validation.Validatable)

	if !ok {
		return next()
	}

	if err := validatable.Validate(); err != nil {
		return nil, err
	}

	return next()
}

func NewValidationBehavior() *ValidationBehavior {
	return &ValidationBehavior{}
}
