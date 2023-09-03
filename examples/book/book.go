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

func NewBookRequest(id int) *BookRequest {

	return &BookRequest{
		Id: int32(id),
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

func SeriallizeBookRequestInHexaDecimal(id int) (string, error) {
	b := NewBookRequest(id)

	serializedBookReq, err := proto.Marshal(b)
	if err != nil {
		return "", err
	}

	hexString := hex.EncodeToString(serializedBookReq)

	return hexString, nil
}
