package restfibonacciserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type AsyncResponse struct {
	RequestId        string `json:"requestid"`
	FibonacciNumbers []int  `json:"fibonacciNumbers"`
	EndOfResponse    bool   `json:"endOfResponse"`
}

type AsyncStore struct {
	mu             sync.Mutex
	current        int
	requestedRange int
	numbers        []int
}

func NewAsyncStore(requestedRange int) *AsyncStore {
	return &AsyncStore{requestedRange: requestedRange}
}

func (ns *AsyncStore) Write(number, current int) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.current = current
	ns.numbers = append(ns.numbers, number)
}

func (ns *AsyncStore) Read() ([]int, int, int) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	readNumbers := make([]int, len(ns.numbers))
	copy(readNumbers, ns.numbers)
	ns.numbers = []int{}
	return readNumbers, ns.current, ns.requestedRange
}

func (a *App) fibonacciAsyncHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]

	numFibonacci, err := strconv.Atoi(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	headers := r.Header
	reqId := headers.Get("request-id")
	if strings.TrimSpace(reqId) == "" {
		http.Error(w, "no request id in request", http.StatusBadRequest)
		return
	}

	if _, ok := a.asyncStores[reqId]; !ok {
		fmt.Printf("creating new store for reqId %s\n", reqId)
		a.asyncStores[reqId] = NewAsyncStore(numFibonacci)
		go a.fibAsync(numFibonacci, reqId)
	}

	numbersNow, current, requested := a.asyncStores[reqId].Read()
	fmt.Printf("read fibs reqId %s till current %d and numbers are: %v\n", reqId, current, numbersNow)
	end := false
	if (current + 1) == requested {
		end = true
		delete(a.asyncStores, reqId)
	}
	response := AsyncResponse{
		RequestId:        reqId,
		FibonacciNumbers: numbersNow,
		EndOfResponse:    end,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (a *App) fibAsync(n int, reqId string) {
	for i := 0; i < n; i++ {
		fmt.Printf("for %s computing and writing fib of %d\n", reqId, i)
		a.asyncStores[reqId].Write(fib(i), i)
	}

}
