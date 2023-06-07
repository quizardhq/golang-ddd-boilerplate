package application

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/mock"
)

func TestFindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	someId := int32(10)
	expected := dto.PersonDTO{ID: 10, Name: "Some person", Age: 20}
	target := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}
	repoMock.
		EXPECT().
		GetPerson(someId).
		Return(entity.Person{ID: 10, Name: "Some person", Age: 20}).
		Times(1)
	mapperMock.
		EXPECT().
		ToDTO(entity.Person{ID: 10, Name: "Some person", Age: 20}).
		Return(expected).
		Times(1)

	result := target.FindOne(int32(someId))

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	expected := []dto.PersonDTO{
		{ID: 10, Name: "Some person", Age: 20},
		{ID: 11, Name: "Another person", Age: 38},
	}
	target := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}
	repoMock.
		EXPECT().
		GetPeople().
		Return([]entity.Person{
			{ID: 10, Name: "Some person", Age: 20},
			{ID: 11, Name: "Another person", Age: 38},
		}).
		Times(1)
	mapperMock.
		EXPECT().
		ToManyDTO([]entity.Person{
			{ID: 10, Name: "Some person", Age: 20},
			{ID: 11, Name: "Another person", Age: 38},
		}).
		Return(expected).
		Times(1)

	result := target.FindAll()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestCreatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	payload := dto.PersonDTO{Name: "Created person", Age: 20}
	expected := dto.PersonDTO{ID: 1, Name: "Created person", Age: 20}
	target := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}
	mapperMock.
		EXPECT().
		FromDTO(dto.PersonDTO{Name: "Created person", Age: 20}).
		Return(entity.Person{ID: 0, Name: "Created person", Age: 20}).
		Times(1)
	repoMock.
		EXPECT().
		CreatePerson(entity.Person{ID: 0, Name: "Created person", Age: 20}).
		Return(entity.Person{ID: 1, Name: "Created person", Age: 20}).
		Times(1)
	mapperMock.
		EXPECT().
		ToDTO(entity.Person{ID: 1, Name: "Created person", Age: 20}).
		Return(expected).
		Times(1)

	result := target.CreatePerson(payload)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestUpdatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	personId := int32(10)
	payload := dto.PersonDTO{ID: 1, Name: "UPDATED person", Age: 21}
	expected := dto.PersonDTO{ID: 1, Name: "UPDATED person", Age: 21}
	target := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}
	mapperMock.
		EXPECT().
		FromDTO(payload).
		Return(entity.Person{ID: 1, Name: "UPDATED person", Age: 21}).
		Times(1)
	repoMock.
		EXPECT().
		UpdatePerson(personId, entity.Person{ID: 1, Name: "UPDATED person", Age: 21}).
		Return(entity.Person{ID: 1, Name: "UPDATED person", Age: 21}).
		Times(1)
	mapperMock.
		EXPECT().
		ToDTO(entity.Person{ID: 1, Name: "UPDATED person", Age: 21}).
		Return(expected).
		Times(1)

	result := target.UpdatePerson(personId, payload)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestDeletePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	someId := int32(10)
	expected := true
	target := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}
	repoMock.
		EXPECT().
		DeletePerson(someId).
		Return(true).
		Times(1)

	result := target.DeletePerson(int32(someId))

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestProvidePersonApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockPersonRepository(ctrl)
	mapperMock := mock.NewMockPersonMapper(ctrl)
	expected := &PersonApplicationImpl{
		repository: repoMock,
		mapper:     mapperMock,
	}

	result := ProvidePersonApplication(repoMock, mapperMock)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}
