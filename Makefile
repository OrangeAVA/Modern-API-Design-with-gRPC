.PHONY: gen-person-proto 
gen-person-proto: 
	protoc --go_out=. --go_opt=paths=source_relative chapter-2/person/person.proto

.PHONY: gen-book-proto 
gen-book-proto: 
	protoc --go_out=. --go_opt=paths=source_relative chapter-2/book/book.proto

.PHONY: gen-info-proto 
gen-info-proto: 
	protoc --go_out=. --go_opt=paths=source_relative chapter-2/size/info.proto
	protoc --go_out=. --go_opt=paths=source_relative chapter-2/serialization/info.proto

.PHONY: run-benchmark
run-benchmark:
	go test -bench=. ./... > benchmark.txt