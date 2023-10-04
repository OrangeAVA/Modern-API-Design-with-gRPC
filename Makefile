REST_BOOKS_SERVER_EXECUTABLE_LOC="./books-app/build/out/$(REST_BOOKS_SERVER_APP_NAME)"
REST_BOOKS_SERVER_APP_NAME=rest-books-server
REST_BOOKS_SERVER_APP_VERSION:=0.0.1
GRPC_BOOKS_SERVER_EXECUTABLE_LOC="./books-app/build/out/$(GRPC_BOOKS_SERVER_APP_NAME)"
GRPC_BOOKS_SERVER_APP_NAME=grpc-books-server
GRPC_BOOKS_SERVER_APP_VERSION:=0.0.1
GRPC_BOOKS_CLIENT_EXECUTABLE_LOC="./books-app/build/out/$(GRPC_BOOKS_CLIENT_APP_NAME)"
GRPC_BOOKS_CLIENT_APP_NAME=grpc-books-client
GRPC_BOOKS_CLIENT_APP_VERSION:=0.0.1
REST_BOOKS_CLIENT_EXECUTABLE_LOC="./books-app/build/out/$(REST_BOOKS_CLIENT_APP_NAME)"
REST_BOOKS_CLIENT_APP_NAME=rest-books-client
REST_BOOKS_CLIENT_APP_VERSION:=0.0.1
APP_COMMIT:=$(shell git rev-parse HEAD)
REST_CONFIG_FILE="./books-app/configs/rest-books-server.yaml"
GRPC_CONFIG_FILE="./books-app/configs/grpc-books-server.yaml"
GRPC_CLIENT_CONFIG_FILE="./books-app/configs/grpc-books-client.yaml"
REST_CLIENT_CONFIG_FILE="./books-app/configs/rest-books-client.yaml"

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

.PHONY: gen-books-app-proto 
gen-books-app-proto: 
	protoc --go_out=books-app/internal/pkg --go-grpc_out=books-app/internal/pkg books-app/internal/pkg/proto/*.proto

.PHONY: compile-rest-server
compile-rest-server:
	go build -a -ldflags "-X main.version=$(REST_BOOKS_SERVER_APP_VERSION) -X main.commit=$(APP_COMMIT)" -o ./books-app/build/out/$(REST_BOOKS_SERVER_APP_NAME) books-app/cmd/rest-books-server/main.go

.PHONY: compile-grpc-server
compile-grpc-server:
	go build -a -ldflags "-X main.version=$(GRPC_BOOKS_SERVER_APP_VERSION) -X main.commit=$(APP_COMMIT)" -o ./books-app/build/out/$(GRPC_BOOKS_SERVER_APP_NAME) books-app/cmd/grpc-books-server/main.go

.PHONY: compile-grpc-client
compile-grpc-client:
	go build -a -ldflags "-X main.version=$(GRPC_BOOKS_CLIENT_APP_VERSION) -X main.commit=$(APP_COMMIT)" -o ./books-app/build/out/$(GRPC_BOOKS_CLIENT_APP_NAME) books-app/cmd/grpc-books-client/main.go

.PHONY: compile-rest-client
compile-rest-client:
	go build -a -ldflags "-X main.version=$(REST_BOOKS_CLIENT_APP_VERSION) -X main.commit=$(APP_COMMIT)" -o ./books-app/build/out/$(REST_BOOKS_CLIENT_APP_NAME) books-app/cmd/rest-books-client/main.go


.PHONY: deps
deps:
	go mod download

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	gotestsum --format=testname  --packages ./... --junitfile report.xml -- -coverprofile=coverage.out ./...

.PHONY: clean
clean:
	go clean -testcache
	rm -rf ./build/out

.PHONY: http-serve
http-serve:
	$(REST_BOOKS_SERVER_EXECUTABLE_LOC) -configFile=$(REST_CONFIG_FILE)

.PHONY: grpc-serve
grpc-serve:
	$(GRPC_BOOKS_SERVER_EXECUTABLE_LOC) -configFile=$(GRPC_CONFIG_FILE)

.PHONY: main-http-serve
main-http-serve:
	go run books-app/cmd/rest-books-server/main.go -configFile=$(REST_CONFIG_FILE)

.PHONY: main-grpc-serve
main-grpc-serve:
	go run books-app/cmd/grpc-books-server/main.go -configFile=$(GRPC_CONFIG_FILE)

.PHONY: execute-grpc-client
execute-grpc-client:
	go run books-app/cmd/grpc-books-client/main.go -configFile=$(GRPC_CLIENT_CONFIG_FILE)

.PHONY: execute-rest-client
execute-rest-client:
	go run books-app/cmd/rest-books-client/main.go -configFile=$(REST_CLIENT_CONFIG_FILE)

.PHONY: start-pg
start-pg:
	docker compose up db

.PHONY: exec-pg
exec-pg:
	docker exec -it <container-id> psql -U postgres