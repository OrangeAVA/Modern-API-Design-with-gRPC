package grpcreviewserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/model"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
)

func (a *App) GetBookReviews(_ context.Context, req *proto.GetBookReviewsRequest) (*proto.GetBookReviewsResponse, error) {
	log.Println("fethcing book reviews")

	// Simulate a potentially failing operation.
	if time.Now().Second()%2 == 0 {
		return nil, fmt.Errorf("failed to fetch book reviews")
	}

	reviews := a.bookRepo.GetAllReviews(int(req.GetIsbn()))

	protoReviews := make([]*proto.Review, 0)

	for _, r := range reviews {
		protoReview := &proto.Review{Reviewer: r.Reviewer, Comment: r.Comment, Rating: int32(r.Rating)}
		protoReviews = append(protoReviews, protoReview)
	}

	return &proto.GetBookReviewsResponse{Reviews: protoReviews}, nil
}

func (a *App) SubmitReviews(_ context.Context, req *proto.SubmitReviewRequest) (*proto.SubmitReviewResponse, error) {
	log.Println("submitting book review")

	review := &model.DBReview{
		Isbn:     int(req.Isbn),
		Comment:  req.GetComment(),
		Reviewer: req.GetReviewer(),
		Rating:   int(req.GetRating()),
	}

	a.bookRepo.AddReview(review)

	return &proto.SubmitReviewResponse{Status: fmt.Sprintf("review for book(%v) submitted successfully", req.GetIsbn())}, nil
}
