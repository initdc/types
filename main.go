package main

import (
	"fmt"
	"os"

	. "github.com/initdc/types/result"
)

func Div(a, b int) Result[int, rune] {
	if b == 0 {
		return Err[int]('E')
	}
	return Ok[int, rune](a / b)
}

func closure(r rune) Result[float32, string] {
	goto clean

clean:
	err := Err[float32](string(r))
	fmt.Printf("%#v\n", err)
	os.Exit(255)
	return err
}

func main() {
	f1 := func(x int) float32 { return float32(x) * 2 }

	r := Div(10, 0)
	fmt.Printf("%#v\n", Try(r, f1, closure))

	e := Div(10, 2)
	fmt.Printf("%#v\n", Try(e, f1, closure))
}
