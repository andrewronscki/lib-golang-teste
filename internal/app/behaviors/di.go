package behaviors

import (
	cqrsdig "github.com/andrewronscki/lib-golang-teste/pkg/commons/cqrs-dig"
	"go.uber.org/dig"
	"go.uber.org/multierr"
)

func Configure(container *dig.Container) error {
	return multierr.Combine(
		cqrsdig.ProvideCommandBehavior[*LoggingBehavior](
			container,
			0,
			NewLoggingBehavior,
		),
		cqrsdig.ProvideQueryBehavior[*LoggingBehavior](
			container,
			0,
			NewLoggingBehavior,
		),
		cqrsdig.ProvideCommandBehavior[*ValidationBehavior](
			container,
			1,
			NewValidationBehavior,
		),
	)
}
