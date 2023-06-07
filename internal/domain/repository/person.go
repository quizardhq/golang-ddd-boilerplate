package repository

import (
	"github.com/google/wire"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/infrastructure/repository"
)

// PersonRepository is a interface for person repository
var PersonRepositorySet = wire.NewSet(
	repository.ProvidePersonRepository,
	wire.Bind(
		new(PersonRepository),
		new(*repository.PersonRepositoryImpl),
	),
)
