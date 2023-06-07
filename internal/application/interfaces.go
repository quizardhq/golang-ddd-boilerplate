package application

import (
	"github.com/google/wire"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/dto"
)

type PersonApplication interface {
	CreatePerson(dto.PersonDTO) dto.PersonDTO
	FindOne(int32) dto.PersonDTO
	FindAll() []dto.PersonDTO
	UpdatePerson(int32, dto.PersonDTO) dto.PersonDTO
	DeletePerson(int32) bool
}

var PersonApplicationSet = wire.NewSet(
	ProvidePersonApplication,
	wire.Bind(
		new(PersonApplication),
		new(*PersonApplicationImpl),
	),
)
