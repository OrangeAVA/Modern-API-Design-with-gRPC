#!/bin/bash

# [REST] Fetch all books
curl --request GET --url http://localhost:8090/books/all | json_pp

# [REST] Fetch a book by isbn
curl --request GET --url http://localhost:8090/books/12346 | json_pp

# [REST] Add a book
curl --request POST \
  --url http://localhost:8090/books \
  --header 'Content-Type: application/json' \
  --data '{
	"isbn": "12348",
	"name": "test book",
	"publisher": "test publisher"
}'

# [REST] Update a book
curl --request PUT \
  --url http://localhost:8090/books \
  --header 'Content-Type: application/json' \
  --data '{
	"isbn": "12348",
	"name": "test book",
	"publisher": "new publisher"
}'

# [REST] Delete a book
curl --request DELETE --url http://localhost:8090/books/12348

