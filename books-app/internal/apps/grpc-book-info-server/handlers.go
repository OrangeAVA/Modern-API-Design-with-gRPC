package grpcbooksserver

import (
	"context"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) GetBookInfoWithReviews(ctx context.Context, req *proto.GetBookInfoRequest) (*proto.GetBookInfoResponse, error) {
	log.Println("fetching book and reviews")

	book, err := a.bookServerClient.GetBook(ctx, &proto.GetBookRequest{Isbn: req.GetIsbn()})
	if err != nil {
		return nil, err
	}

	resp, err := a.reviewServerClient.GetBookReviews(ctx, &proto.GetBookReviewsRequest{Isbn: req.GetIsbn()})
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
