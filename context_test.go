package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	// ambil context.Background()
	background := context.Background()
	// lihat isi dari context background
	fmt.Println(background)

	// ambil context TODO
	todo := context.TODO()
	fmt.Println(todo)

}

// context with value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// parentnya adalah ContextA, keynya adalah a, valuenya adalah A
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// parentnya adalah ContextB
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	// parentnya adalah ContextC
	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Context Get Value

	fmt.Println(contextF.Value("f")) // dapat data valuenya, karena contextF isinya adalah f
	fmt.Println(contextF.Value("c")) // tetap dapat data valuenya, karena dia akan mencari ke parentnya, dan parent dari contextF adalah contextC
	fmt.Println(contextF.Value("b")) // tidak dapat data (nil), karena b berada di contextB, sedangkan contextB bukan parent dari contextF
	fmt.Println(contextA.Value("b")) // tidak dapat data (nil), meskipun b berada di contextB dan contextB adalah child dari contextA. Dikarenakan tidak bisa mengambil data child nya
}

// buat function CreateCounter dengan balikannya adalah channel yang tipenya integer
func CreateCounter() chan int {

	destination := make(chan int)

	// untuk isi dari channelnya kita menggunakan goroutine
	// jika goroutine selesai dikerjakan, maka kita close juga channelnya
	go func() {
		defer close(destination)
		counter := 1
		for {
			// data counter kita kirim ke channel destination
			// karena ini perulang tanpa henti, ini menyebabkan goroutine leak
			destination <- counter
			counter++
		}
	}()
	return destination
}

// mengatasi problem goroutine leak dengan sinyal Cancel
func CreateCounterContext(ctx context.Context) chan int {

	// kita cek contextnya, apakah dia sudah dibatalkan atau belum
	// kalo misal belum dibatalkan, maka kita lakukan perulangan (counter) terus-menerus
	// jika contextnya sudah dibatalkan, maka kita batalkan perulangan counternya

	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			// kita select dari context.done
			// jika context sudah done, maka langsung kita return
			// alias menghentikan goroutinenya
			// defaultnya dijalankan jika tidak terjadi Done
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	// ambil destination
	// ada parameter context, nah gimana cara ngirim contextnya?
	// kita buat dulu parent contextnya
	parent := context.Background()
	// bikin context.WithCancel, ada 2 balikannya,yakni Context dan function cancel
	ctx, cancel := context.WithCancel(parent)
	// kalo langsung dibatalin bisa langsung panggil cancel
	// kita ingin setelah proses goroutine, maka langsung dicancel

	destination := CreateCounterContext(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine()) // harusnya 3 goroutine

	// perulangan untuk mencetak destination
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
