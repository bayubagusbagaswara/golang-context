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
