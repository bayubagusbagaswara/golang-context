package golangcontext

import (
	"context"
	"fmt"
	"testing"
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
