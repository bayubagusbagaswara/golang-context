package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterContextDeadline(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()
	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	// kita menggunakan context WithDeadline
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

	defer cancel()

	destination := CreateCounterContextDeadline(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
