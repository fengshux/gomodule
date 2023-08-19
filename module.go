package gomodule

import (
	"fmt"

	"github.com/fengshux/bitmap"
)

func Module() {
	fmt.Println("gomodule")
	b := bitmap.New(uint64(20))
	fmt.Println(b)
}
