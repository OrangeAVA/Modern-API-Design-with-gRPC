package grpcbooksserver

import (
	"context"
	"fmt"
	"time"
)

type CircuitBreaker struct {
	threshold    int
	failureCount int
	open         bool
}

// NewCircuitBreaker creates a new CircuitBreaker.
func NewCircuitBreaker(threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		threshold: threshold,
		open:      false,
	}
}

// Execute executes a function using the circuit breaker pattern.
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	if cb.open {
		return nil, fmt.Errorf("circuit breaker is open")
	}

	result, err := fn()

	if err != nil {
		cb.failureCount++

		if cb.failureCount >= cb.threshold {
			cb.open = true
			go func() {
				// Automatically attempt to close the circuit after a timeout.
				<-time.After(5 * time.Second)
				cb.open = false
				cb.failureCount = 0
			}()
		}
	} else {
		// Reset the failure count on successful execution.
		cb.failureCount = 0
	}

	return result, err
}
