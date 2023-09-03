package person

func NewPerson(name string, age int, phoneNumbers []string) *Person {
	phNos := make([]string, 0)
	phNos = append(phNos, phoneNumbers...)
	return &Person{
		Name:         name,
		Age:          int32(age),
		PhoneNumbers: phNos,
	}
}
