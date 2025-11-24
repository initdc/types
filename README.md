# types

Bring the Rust [Result Option] types to Golang

## Installation

  `go get github.com/initdc/types`

## Usage

```go
package e2e_test

import (
	"fmt"
	. "github.com/initdc/types/option"
	. "github.com/initdc/types/result"
	"testing"
)

func TestE2E(t *testing.T) {
	// Option
	s1 := Some(1)

	var s2 Option[int]
	s2.None()

	n1 := None[int]()


	// Result
	var r1 Result[int, string]
	r1.Ok(1)

	var e1 Result[int, string]
	e1.Err("error")

	r2 := Ok[int, string](1)
	e2 := Err[int, string]("error")

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)
	fmt.Printf("%#v\n", n1)

	fmt.Printf("%#v\n", r1)
	fmt.Printf("%#v\n", e1)
	fmt.Printf("%#v\n", r2)
	fmt.Printf("%#v\n", e2)
}
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/initdc/types.
