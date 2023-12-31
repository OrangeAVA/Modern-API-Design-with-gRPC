package grpcbooksserver

import (
	"context"
	"fmt"
	"log"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/proto"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (a *App) GetBookInfoWithReviews(ctx context.Context, req *proto.GetBookInfoRequest) (*proto.GetBookInfoResponse, error) {
	// Start a root span.
	rootSpan := tracer.StartSpan("get.bookInfo.withReviews")
	defer rootSpan.Finish()

	log.Println("fetching book and reviews")

	bookSpan := spawnChildSpan(rootSpan, "get.bookInfo", fmt.Sprintf("%d", req.GetIsbn()))

	book, err := a.bookServerClient.GetBook(ctx, &proto.GetBookRequest{Isbn: req.GetIsbn()})
	bookSpan.Finish(tracer.WithError(err))
	if err != nil {
		return nil, err
	}

	reviewSpan := spawnChildSpan(rootSpan, "get.bookReview", fmt.Sprintf("%d", req.GetIsbn()))

	resp, err := a.reviewServerClient.GetBookReviews(ctx, &proto.GetBookReviewsRequest{Isbn: req.GetIsbn()})
	reviewSpan.Finish(tracer.WithError(err))
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

func spawnChildSpan(parentSpan ddtrace.Span, spanName, resourceName string) ddtrace.Span {
	child := tracer.StartSpan(spanName, tracer.ChildOf(parentSpan.Context()))
	child.SetTag(ext.ResourceName, resourceName)

	// If you are using 128 bit trace ids and want to generate the high
	// order bits, cast the span's context to ddtrace.SpanContextW3C.
	if w3Cctx, ok := child.Context().(ddtrace.SpanContextW3C); ok {
		fmt.Printf("128 bit trace id = %s\n", w3Cctx.TraceID128())
	}

	return child
}
