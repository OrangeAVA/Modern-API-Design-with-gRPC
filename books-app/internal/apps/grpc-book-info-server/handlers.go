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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define a retryable function for getting a book.
	retryableGetBook := func() (interface{}, error) {
		return a.bookServerClient.GetBook(ctx, &proto.GetBookRequest{Isbn: req.GetIsbn()})
	}

	// Execute the retryableGetBook function with retry logic.
	result, err := WithRetry(ctx, retryableGetBook, 3, 2*time.Second)
	if err != nil {
		return nil, err
	}

	book, ok := result.(*proto.Book)
	if !ok {
		return nil, fmt.Errorf("could not fetch book with isbn(%v)", req.GetIsbn())
	}

	// Define a retryable function for getting reviews.
	retryableGetReviews := func() (interface{}, error) {
		return a.reviewServerClient.GetBookReviews(ctx, &proto.GetBookReviewsRequest{Isbn: req.GetIsbn()})
	}

	// Execute the retryableGetReviews function with retry logic.
	result, err = WithRetry(ctx, retryableGetReviews, 3, 3*time.Second)
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*proto.GetBookReviewsResponse)
	if !ok {
		return nil, fmt.Errorf("could not fetch reviews for book with isbn(%v)", req.GetIsbn())
	}

	return &proto.GetBookInfoResponse{
		Isbn:      req.GetIsbn(),
		Name:      book.Name,
		Publisher: book.Publisher,
		Reviews:   resp.GetReviews(),
	}, nil
}

// RetryableFunc represents a function that can be retried.
type RetryableFunc func() (interface{}, error)

// WithRetry executes a function with retry logic.
func WithRetry(ctx context.Context, retryFunc RetryableFunc, maxAttempts int, retryInterval time.Duration) (interface{}, error) {
	var result interface{}
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			// Operation was canceled.
			return nil, ctx.Err()
		default:
			// [No-Op] Continue with the retry.
		}

		result, err = retryFunc()

		if err == nil {
			// Operation succeeded, break the loop.
			break
		}

		// Log or handle the error (optional).
		fmt.Printf("Attempt %d failed: %v\n", attempt, err)

		// Wait before the next retry.
		time.Sleep(retryInterval)
	}

	return result, err
}
