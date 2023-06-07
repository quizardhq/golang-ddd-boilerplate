package entity

// Person is a dummy model
type Person struct {
	ID   int32  `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int32  `gorm:"column:age"`
}

// ProvidePerson provides a person model
func ProvidePerson() Person {
	return Person{}
}