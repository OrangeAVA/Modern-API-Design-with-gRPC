package grpcbooksserver

import (
	"context"
	"log"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) GetBookInfoWithReviews(ctx context.Context, req *proto.GetBookInfoRequest) (*proto.GetBookInfoResponse, error) {
	log.Println("fetching book and reviews")

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	book, err := a.bookServerClient.GetBook(newCtx, &proto.GetBookRequest{Isbn: req.GetIsbn()})
	if err != nil {
		return nil, err
	}

	resp, err := a.reviewServerClient.GetBookReviews(newCtx, &proto.GetBookReviewsRequest{Isbn: req.GetIsbn()})
	if err != nil {
		return nil, err
	}

	return &proto.GetBookInfoResponse{
		Isbn:      req.GetIsbn(),
		Name:      book.Name,
		Publisher: book.Publisher,
		Reviews:   resp.GetReviews(),
	}, nil
}
