package request

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"

	"google.golang.org/grpc/metadata"
)

// Code is taken from: https://github.com/go-chi/chi/blob/master/middleware/request_id.go

// Key to use when setting the request ID.
type ctxKeyRequestID int

// RequestIDKey is the key that holds th unique request ID in a request context.
const RequestIDKey string = "request-id"

var (
	// prefix is const prefix for request ID
	prefix string

	// reqID is counter for request ID
	reqID uint64
)

// init Initializes constant part of request ID
func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	var buf [12]byte
	var b64 string
	for len(b64) < 10 {
		_, _ = rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

// RequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
func AddRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myid := atomic.AddUint64(&reqID, 1)
		ctx := r.Context()
		reqID := fmt.Sprintf("%s-%06d", prefix, myid)
		ctx = context.WithValue(ctx, RequestIDKey, reqID)
		// to pass the request id to grpc methods, we need to pass it as a metadata

		h.ServeHTTP(w, r.WithContext(addGrpcReqId(ctx, reqID)))
	})
}

func addGrpcReqId(ctx context.Context, reqID string) context.Context {
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs(RequestIDKey, reqID))
	return ctx
}

//func WithReqId(httpCtx context.Context) context.Context {
//	ctx := context.Background()
//	ctx = metadata.NewOutgoingContext(
//		ctx,
//		metadata.Pairs(RequestIDKey, GetReqID(httpCtx)))
//
//	return ctx
//}

func GetContextRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md[RequestIDKey]) > 0 {
		return md[RequestIDKey][0]
	}
	return GetReqID(ctx)
}

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
