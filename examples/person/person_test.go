package person_test

import (
	"testing"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/examples/person"
	"github.com/stretchr/testify/assert"
)

func Test_NewPerson(t *testing.T) {
	p := person.NewPerson("Hitesh Pattanayak", 29, []string{"9012345678", "9087654321"})

	assert.Equal(t, "Hitesh Pattanayak", p.GetName())
	assert.Equal(t, int32(29), p.GetAge())
	assert.Equal(t, []string{"9012345678", "9087654321"}, p.GetPhoneNumbers())
}
