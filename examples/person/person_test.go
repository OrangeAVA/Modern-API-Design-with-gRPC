package person_test

import (
	"strings"
	"testing"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/examples/person"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewPerson(t *testing.T) {
	p := person.NewPerson("Hitesh Pattanayak", 29, []string{"9012345678", "9087654321"})

	assert.Equal(t, "Hitesh Pattanayak", p.GetName())
	assert.Equal(t, int32(29), p.GetAge())
	assert.Equal(t, []string{"9012345678", "9087654321"}, p.GetPhoneNumbers())
}

func Test_SeriallizePersonMessage(t *testing.T) {
	msg, err := person.SeriallizePersonMessage("Hitesh Pattanayak", 29, []string{"9012345678", "9087654321"})
	require.NoError(t, err)
	assert.True(t, strings.Contains(msg, "Hitesh Pattanayak"))
	assert.True(t, strings.Contains(msg, "9012345678"))
	assert.True(t, strings.Contains(msg, "9087654321"))
}

func Test_DeseriallizePersonMessage(t *testing.T) {
	msg, err := person.SeriallizePersonMessage("Hitesh Pattanayak", 29, []string{"9012345678", "9087654321"})
	require.NoError(t, err)

	deseriallizePerson, err := person.DeseriallizePersonMessage([]byte(msg))
	require.NoError(t, err)

	assert.Equal(t, "Hitesh Pattanayak", deseriallizePerson.GetName())
	assert.Equal(t, int32(29), deseriallizePerson.GetAge())
	assert.Equal(t, []string{"9012345678", "9087654321"}, deseriallizePerson.GetPhoneNumbers())
}
