package mapAndStruct

import "fmt"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type person struct {
	name   string
	age    int
	gender Gender
}

func isValidName(name string) bool {
	return len(name) != 0
}

func isValidAge(age int) bool {
	return age > 0
}

func isValidGender(gender Gender) bool {
	return gender == Male || gender == Female
}

func validateCreatePersonInput(name string, age int, gender Gender) error {
	if !isValidName(name) {
		return fmt.Errorf("error: length of name cannot be zero")
	}
	if !isValidAge(age) {
		return fmt.Errorf("error: age cannot be less than 0")
	}
	if !isValidGender(gender) {
		return fmt.Errorf("error: gender must be either male or female")
	}
	return nil
}

func CreatePerson(name string, age int, gender Gender) (p *person, err error) {
	err = validateCreatePersonInput(name, age, gender)
	if err != nil {
		return nil, err
	}
	p = &person{
		name, age, gender,
	}
	return p, nil
}
