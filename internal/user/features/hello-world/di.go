package helloworld

import (
	cqrsdig "github.com/andrewronscki/lib-golang-teste/pkg/commons/cqrs-dig"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func Configure(container *dig.Container) error {
	return multierr.Combine(
		cqrsdig.ProvideCommandHandler[*Command, *Model](
			container,
			NewCommandHandler,
		),
	)
}
