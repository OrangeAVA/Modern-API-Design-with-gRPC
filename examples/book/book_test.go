package book_test

import (
	"testing"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/examples/book"
	"github.com/stretchr/testify/assert"
)

func Test_NewBook(t *testing.T) {
	b := book.NewBook("1234", "Ava Orange")

	assert.Equal(t, "1234", b.GetIsbn())
	assert.Equal(t, "Ava Orange", b.GetPublisher())
}
