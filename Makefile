.PHONY: gen-person-proto 
gen-person-proto: 
	protoc --go_out=. --go_opt=paths=source_relative examples/person/person.proto