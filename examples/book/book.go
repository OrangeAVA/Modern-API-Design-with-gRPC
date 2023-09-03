package book

func NewBook(isbn, publisher string) *BookInfo {
	return &BookInfo{
		Isbn:      isbn,
		Publisher: publisher,
	}
}
