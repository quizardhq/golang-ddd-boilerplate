package mapper

import (
	"github.com/google/wire"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/dto"
)

type PersonMapper interface {
	ToDTO(entity.Person) dto.PersonDTO
	FromDTO(dto.PersonDTO) entity.Person
	ToManyDTO([]entity.Person) []dto.PersonDTO
}

var PersonMapperSet = wire.NewSet(
	ProvidePersonMapper,
	wire.Bind(
		new(PersonMapper),
		new(*PersonMapperImpl),
	),
)
