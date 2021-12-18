package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterContextTimeout(ctx context.Context) chan int {
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
				// tambahkan time sleep 1 detik
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()
	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	// kita menggunakan context WithTimeout
	// kalau misal kejadian atau timeoutnya kecil, sehingga akan cepat dalam melakukan cancelnya, maka tetap disarankan memanggil function cancel
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	// defer, setelah function goroutine selesai dieksekusi, maka dia akan tetap membatalkan contextnya (memastikan aja tidak terjadi goroutine leak)
	defer cancel()

	destination := CreateCounterContextTimeout(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	// ini looping tidak pernah berhenti
	for n := range destination {
		fmt.Println("Counter", n)
	}
	// sekarang tugas function cancel() diarahkan ke defer

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
