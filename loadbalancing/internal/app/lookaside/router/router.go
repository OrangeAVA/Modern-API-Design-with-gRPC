package router

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Router struct {
	Addresses                []string
	LastRefresh              time.Time
	RefreshIntervalInSeconds float64
	roundRobinPtr            int
}

func (r *Router) ResolveRandom() (string, error) {
	fmt.Println("random")

	if len(r.Addresses) == 0 {
		return "", errors.New("attempted to resolve address with empty array")
	}

	return r.Addresses[rand.Intn(len(r.Addresses))], nil

}

func (r *Router) ResolveRoundRobin() (string, error) {
	fmt.Println("roundrobin", r.roundRobinPtr)

	if len(r.Addresses) == 0 {
		return "", errors.New("attempted to resolve address with empty array")
	}

	if r.roundRobinPtr == len(r.Addresses) {
		r.roundRobinPtr = 0
	}

	address := r.Addresses[r.roundRobinPtr]
	r.roundRobinPtr++
	return address, nil

}

func (r *Router) NeedsRefresh() bool {
	return time.Since(r.LastRefresh).Seconds() >= r.RefreshIntervalInSeconds
}
