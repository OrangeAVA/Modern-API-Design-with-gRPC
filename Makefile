.PHONY: gen-person-proto 
gen-person-proto: 
	protoc --go_out=. --go_opt=paths=source_relative examples/person/person.proto

.PHONY: gen-book-proto 
gen-book-proto: 
	protoc --go_out=. --go_opt=paths=source_relative examples/book/book.proto