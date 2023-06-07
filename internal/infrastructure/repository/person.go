package repository

import (
	"log"

	"github.com/quizardhq/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/quizardhq/golang-ddd-boilerplate/internal/infrastructure/persistence"
)

// PersonRepositoryImpl is self explanatory
type PersonRepositoryImpl struct {
	Database *persistence.Database
}

// GetOne returns a persisted Person
func (pr *PersonRepositoryImpl) GetPerson(id int32) entity.Person {
	person := entity.Person{}
	result := pr.Database.GormDB.First(&person, id)
	if result.Error != nil {
		log.Println("Error on Query: ", result.Error.Error())
		return person
	}
	return person
}

// GetAll returns all persisted Person
func (pr *PersonRepositoryImpl) GetPeople() []entity.Person {
	people := []entity.Person{}
	result := pr.Database.GormDB.Find(&people)
	if result.Error != nil {
		log.Println("Error on Query: ", result.Error.Error())
		return nil
	}
	return people
}

// Create inserts a person into Database
func (pr *PersonRepositoryImpl) CreatePerson(person entity.Person) entity.Person {
	p := entity.Person{}
	result := pr.Database.GormDB.Create(&person)
	if result.Error != nil {
		log.Println("Error on create: ", result.Error.Error())
		return entity.Person{}
	}
	p = person
	return p
}

// Update updates an person
func (pr *PersonRepositoryImpl) UpdatePerson(id int32, person entity.Person) entity.Person {
	p := entity.Person{}
	db := pr.Database.GormDB
	result := db.Model(&p).Where("id = ?", id).Updates(map[string]interface{}{
		"name": person.Name,
		"age":  person.Age,
	})

	if result.Error != nil {
		log.Println("Error on Update: ", result.Error.Error())
		return entity.Person{}
	}

	// Retrieve the updated person
	err := db.First(&p, id).Error
	if err != nil {
		log.Println("Error on Query: ", err.Error())
		return entity.Person{}
	}

	return p
}

// Delete deeltes a persisted Person
func (pr *PersonRepositoryImpl) DeletePerson(id int32) bool {
	db := pr.Database.GormDB
	result := db.Delete(&entity.Person{}, id)

	if result.Error != nil {
		log.Println("Error on Delete: ", result.Error.Error())
		return false
	}

	return true
}

// ProvidePersonRepository providers person repository
func ProvidePersonRepository(database *persistence.Database) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{
		Database: database,
	}
}
