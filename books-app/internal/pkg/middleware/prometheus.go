package middleware

import (
	"fmt"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Incoming_api_req_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "request_counter_to_api",
		Help:      "counts incoming requests to api",
	})

var Book_get_pass_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "get_book_pass_ctr",
		Help:      "counts success to fetch book",
	})

var Book_get_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "get_book_fail_ctr",
		Help:      "counts fail to fetch book",
	})

var Book_create_pass_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "create_book_pass_ctr",
		Help:      "counts success to create book",
	})

var Book_update_pass_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "update_book_pass_ctr",
		Help:      "counts success to update book",
	})

var Book_delete_pass_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "delete_book_pass_ctr",
		Help:      "counts success to delete book",
	})

func Register() {
	logger.Log.Info("registering of prometheus custom metrics starts")

	prometheus.MustRegister(Incoming_api_req_counter)
	prometheus.MustRegister(Book_get_pass_counter)
	prometheus.MustRegister(Book_create_pass_counter)
	prometheus.MustRegister(Book_update_pass_counter)
	prometheus.MustRegister(Book_delete_pass_counter)
	prometheus.MustRegister(Book_get_fail_counter)

	logger.Log.Info("registering of prometheus custom metrics ends")
}

func RunPrometheusServer(promPort int) {
	Register()
	http.Handle("/metrics", promhttp.Handler())
	port := promPort
	go func() {
		logger.Log.Info("starting prometheus server....")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			logger.Log.Warn("Unable to start a http server for prometheus : " + err.Error())
		}
	}()
}

func AddPrometheus(uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	*uInterceptors = append(*uInterceptors, grpc_prometheus.UnaryServerInterceptor)
	*sInterceptors = append(*sInterceptors, grpc_prometheus.StreamServerInterceptor)
}
