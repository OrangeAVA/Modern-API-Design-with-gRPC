package grpcbooksserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) GetBookInfoWithReviews(ctx context.Context, req *proto.GetBookInfoRequest) (*proto.GetBookInfoResponse, error) {
	log.Println("fetching book and reviews")

	// Set a deadline for the entire operation.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Define a function for getting a book.
	getBookFunc := func() (interface{}, error) {
		return a.bookServerClient.GetBook(ctx, &proto.GetBookRequest{Isbn: req.GetIsbn()})
	}

	// Execute getBookFunc function using the circuit breaker.
	result, err := a.cb.Execute(ctx, getBookFunc)

	if err != nil {
		return nil, err
	}

	// Display the result.
	book, ok := result.(*proto.Book)
	if !ok {
		return nil, fmt.Errorf("could not fetch book with isbn(%v)", req.GetIsbn())
	}

	// Define a function for getting book reviews.
	getBookReviewsFunc := func() (interface{}, error) {
		return a.reviewServerClient.GetBookReviews(ctx, &proto.GetBookReviewsRequest{Isbn: req.GetIsbn()})
	}

	// Execute getBookReviewsFunc function using the circuit breaker.
	result, err = a.cb.Execute(ctx, getBookReviewsFunc)

	if err != nil {
		return nil, err
	}

	// Display the result.
	resp, ok := result.(*proto.GetBookReviewsResponse)
	if !ok {
		return nil, fmt.Errorf("could not fetch book reviews with isbn(%v)", req.GetIsbn())
	}

	return &proto.GetBookInfoResponse{
		Isbn:      req.GetIsbn(),
		Name:      book.Name,
		Publisher: book.Publisher,
		Reviews:   resp.GetReviews(),
	}, nil
}
