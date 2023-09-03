package book

import (
	"encoding/hex"

	"google.golang.org/protobuf/proto"
)

func NewBook(isbn, publisher string) *BookInfo {
	return &BookInfo{
		Isbn:      isbn,
		Publisher: publisher,
	}
}

func SeriallizeBookInfoInHexaDecimal(isbn, publisher string) (string, error) {
	b := NewBook(isbn, publisher)

	serializedBook, err := proto.Marshal(b)
	if err != nil {
		return "", err
	}

	hexString := hex.EncodeToString(serializedBook)

	return hexString, nil
}
