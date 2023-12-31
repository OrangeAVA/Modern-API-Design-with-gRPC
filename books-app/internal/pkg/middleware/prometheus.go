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

var Emp_get_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "get_employee_fail_ctr",
		Help:      "counts failure to fetch employee",
	})

var Emp_create_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "create_employee_fail_ctr",
		Help:      "counts failure to create employee",
	})

var Emp_update_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "update_employee_fail_ctr",
		Help:      "counts failure to update employee",
	})

var Emp_delete_fail_counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "delete_employee_fail_ctr",
		Help:      "counts failure to delete employee",
	})

func Register() {
	logger.Log.Info("registering of prometheus custom metrics starts")

	prometheus.MustRegister(Incoming_api_req_counter)
	prometheus.MustRegister(Emp_get_fail_counter)
	prometheus.MustRegister(Emp_create_fail_counter)
	prometheus.MustRegister(Emp_update_fail_counter)
	prometheus.MustRegister(Emp_delete_fail_counter)

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
