#!/bin/bash

# install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# [gRPC] list all services
grpcurl -plaintext 0.0.0.0:50051 list

# [gRPC] list all supported methods
grpcurl -plaintext 0.0.0.0:50051 list prot.BookService

# [gRPC] get all books
grpcurl -plaintext 0.0.0.0:50051  prot.BookService/ListBooks

# [gRPC] get a book
grpcurl -plaintext -d '{"isbn": "12345"}' 0.0.0.0:50051  prot.BookService/GetBook

# [gRPC] add a book
grpcurl -plaintext -d '{"isbn": "12348", "name": "test name", "publisher": "test publisher"}' 0.0.0.0:50051  prot.BookService/AddBook

# [gRPC] update a book
grpcurl -plaintext -d '{"isbn": "12348", "name": "test name", "publisher": "new publisher"}' 0.0.0.0:50051  prot.BookService/UpdateBook

# [gRPC] remove a book
grpcurl -plaintext -d '{"isbn": "12348"}' 0.0.0.0:50051  prot.BookService/RemoveBook

# [gRPC] submit a review
grpcurl -plaintext -d '{"isbn": "12347", "reviewer": "hitesh pattanayak", "comment": "good book", "rating": "5"}' 0.0.0.0:50052  prot.ReviewService/SubmitReviews
grpcurl -plaintext -d '{"isbn": "12347", "reviewer": "another person", "comment": "great book", "rating": "4"}' 0.0.0.0:50052  prot.ReviewService/SubmitReviews
grpcurl -plaintext -d '{"isbn": "12347"}' 0.0.0.0:50052  prot.ReviewService/GetBookReviews

# [gRPC] get a book with reviews
grpcurl -plaintext -d '{"isbn": "12347"}' 0.0.0.0:50053  prot.BookInfoService.GetBookInfoWithReviews