package person

import (
	"github.com/golang/protobuf/proto"
)

func NewPerson(name string, age int, phoneNumbers []string) *Person {
	phNos := make([]string, 0)
	phNos = append(phNos, phoneNumbers...)
	return &Person{
		Name:         name,
		Age:          int32(age),
		PhoneNumbers: phNos,
	}
}

func SeriallizePersonMessage(name string, age int, phoneNumbers []string) (string, error) {
	person := NewPerson(name, age, phoneNumbers)

	serializedPerson, err := proto.Marshal(person)
	if err != nil {
		return "", err
	}

	return string(serializedPerson), nil
}

func DeseriallizePersonMessage(serializedPerson []byte) (*Person, error) {
	person := Person{}

	err := proto.Unmarshal(serializedPerson, &person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
