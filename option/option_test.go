package option

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/initdc/types/result"
	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	var a = Option[int]{
		value:    1,
		some:     true,
		assigned: true,
	}

	var b Option[int]
	b.Some(1)

	c := Some[int](1)

	assert.Equal(t, a, b)
	assert.Equal(t, a, *c)
	assert.Equal(t, b, *c)
}

func TestNone(t *testing.T) {
	var a = Option[int]{
		some:     false,
		assigned: true,
	}

	var b Option[int]
	b.None()

	c := None[int]()

	assert.Equal(t, a, b)
	assert.Equal(t, a, *c)
	assert.Equal(t, b, *c)
}

func TestIsSome(t *testing.T) {
	s := Some(1)
	assert.True(t, s.IsSome())

	n := None[int]()
	assert.False(t, n.IsSome())
}

func TestIsSomeAnd(t *testing.T) {
	s := Some(1)
	assert.True(t, s.IsSomeAnd(func(int) bool { return true }))
	assert.False(t, s.IsSomeAnd(func(int) bool { return false }))

	n := None[int]()
	assert.False(t, n.IsSomeAnd(func(int) bool { return true }))
}

func TestIsNone(t *testing.T) {
	s := Some(1)
	assert.False(t, s.IsNone())

	n := None[int]()
	assert.True(t, n.IsNone())
}

func TestIsNoneOr(t *testing.T) {
	s := Some(1)
	assert.False(t, s.IsNoneOr(func(int) bool { return false }))
	assert.True(t, s.IsNoneOr(func(int) bool { return true }))

	n := None[int]()
	assert.True(t, n.IsNoneOr(func(int) bool { return false }))
}

func TestExpect(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.Expect("error"), 1)

	n := None[int]()
	assert.Panics(t, func() { n.Expect("error") })
}

func TestUnwrap(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.Unwrap(), 1)

	n := None[int]()
	assert.Panics(t, func() { n.Unwrap() })
}

func TestUnwrapOr(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.UnwrapOr(99), 1)

	n := None[int]()
	assert.Equal(t, n.UnwrapOr(99), 99)
}

func TestUnwrapOrElse(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.UnwrapOrElse(func() int { return 99 }), 1)

	n := None[int]()
	assert.Equal(t, n.UnwrapOrElse(func() int { return 99 }), 99)
}

func TestUnwrapOrDefault(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.UnwrapOrDefault(), 1)

	n := None[int]()
	assert.Equal(t, n.UnwrapOrDefault(), 0)
}

func TestOptionMap(t *testing.T) {
	f1 := func(s string) int { return len(s) }
	f2 := func(x int) int { return x }

	s := Some("Hello, World!")
	assert.Equal(t, OptionMap(s, f1), Some[int](13))
	assert.Equal(t, To(s, f1).Map(f2), Some[int](13))
	assert.Equal(t, To(s, f1), Some[int](13))

	n := None[string]()
	assert.Equal(t, OptionMap(n, f1), None[int]())
	assert.Equal(t, To(n, f1).Map(f2), None[int]())
	assert.Equal(t, To(n, f1), None[int]())
}

func TestInspect(t *testing.T) {
	s := Some(1)
	assert.Equal(t, s.Inspect(func(x int) {}), s)

	n := None[int]()
	assert.Equal(t, n.Inspect(func(x int) {}), n)
}

func TestOptionMapOr(t *testing.T) {
	f1 := func(s string) int { return len(s) }
	f2 := func(x int) int { return x }

	s := Some("foo")
	assert.Equal(t, OptionMapOr(s, 42, f1), 3)
	assert.Equal(t, To(s, f1).MapOr(42, f2), 3)

	n := None[string]()
	assert.Equal(t, OptionMapOr(n, 42, f1), 42)
	assert.Equal(t, To(n, f1).MapOr(42, f2), 42)
}

func TestOptionMapOrElse(t *testing.T) {
	k := 21
	def := func() int { return 2 * k }
	f1 := func(s string) int { return len(s) }
	f2 := func(x int) int { return x }

	s := Some("foo")
	assert.Equal(t, OptionMapOrElse(s, def, f1), 3)
	assert.Equal(t, To(s, f1).MapOrElse(def, f2), 3)

	n := None[string]()
	assert.Equal(t, OptionMapOrElse(n, def, f1), 42)
	assert.Equal(t, To(n, f1).MapOrElse(def, f2), 42)
}

func TestOptionMapOrDefault(t *testing.T) {
	f1 := func(s string) int { return len(s) }
	f2 := func(x int) int { return x }

	s := Some("hi")
	assert.Equal(t, OptionMapOrDefault(s, f1), 2)
	assert.Equal(t, To(s, f1).MapOrDefault(f2), 2)

	n := None[string]()
	assert.Equal(t, OptionMapOrDefault(n, f1), 0)
	assert.Equal(t, To(n, f1).MapOrDefault(f2), 0)
}

func TestOkOr(t *testing.T) {
	x := Some("foo")
	assert.Equal(t, OkOr(x, 0), result.Ok[string, int]("foo"))

	y := None[string]()
	assert.Equal(t, OkOr(y, 0), result.Err[string, int](0))
}

func TestOkOrElse(t *testing.T) {
	x := Some("foo")
	assert.Equal(t, OkOrElse(x, func() int { return 0 }), result.Ok[string, int]("foo"))

	y := None[string]()
	assert.Equal(t, OkOrElse(y, func() int { return 0 }), result.Err[string, int](0))
}

func TestOptionAnd(t *testing.T) {
	f := func(x int) string { return fmt.Sprintf("%d", x) }

	x := Some(2)
	y := None[string]()
	assert.Equal(t, OptionAnd(x, y), None[string]())
	assert.Equal(t, To(x, f).And(y), None[string]())

	x2 := None[int]()
	y2 := Some("foo")
	assert.Equal(t, OptionAnd(x2, y2), None[string]())
	assert.Equal(t, To(x2, f).And(y2), None[string]())

	x3 := Some(2)
	y3 := Some("foo")
	assert.Equal(t, OptionAnd(x3, y3), Some("foo"))
	assert.Equal(t, To(x3, f).And(y3), Some("foo"))

	x4 := None[int]()
	y4 := None[string]()
	assert.Equal(t, OptionAnd(x4, y4), None[string]())
	assert.Equal(t, To(x4, f).And(y4), None[string]())
}

func TestOptionAndThen(t *testing.T) {
	sqThenToString := func(x int) *Option[string] {
		if x > 100 || x < 0 {
			return None[string]()
		}
		return Some(strconv.Itoa(x))
	}

	assert.Equal(t, OptionAndThen(Some(2), sqThenToString), Some[string]("2"))
	assert.Equal(t, OptionAndThen(Some(1000), sqThenToString), None[string]())
	assert.Equal(t, OptionAndThen(None[int](), sqThenToString), None[string]())
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	assert.Equal(t, None[int]().Filter(isEven), None[int]())
	assert.Equal(t, Some(3).Filter(isEven), None[int]())
	assert.Equal(t, Some(4).Filter(isEven), Some(4))
}

func TestOr(t *testing.T) {
	x := Some(2)
	y := None[int]()

	z := Some(2)
	z.Unwrap()
	assert.Equal(t, x.Or(y), z)

	x2 := None[int]()
	y2 := Some(100)
	assert.Equal(t, x2.Or(y2), Some(100))

	x3 := Some(2)
	y3 := Some(100)

	assert.Equal(t, x3.Or(y3), z)

	x4 := None[int]()
	y4 := None[int]()
	assert.Equal(t, x4.Or(y4), None[int]())
}

func TestOrElse(t *testing.T) {
	nobody := func() *Option[int] { return None[int]() }
	vikings := func() *Option[int] { return Some(42) }

	z := Some(10)
	z.Unwrap()

	assert.Equal(t, Some(10).OrElse(vikings), z)
	assert.Equal(t, None[int]().OrElse(vikings), Some(42))
	assert.Equal(t, None[int]().OrElse(nobody), None[int]())
}

func TestTryOr(t *testing.T) {
	f := func(x int) float32 { return float32(x) * 2 }
	x := Some(2)
	y := None[int]()

	assert.Equal(t, TryOr(x, f, "bad"), result.Ok[float32, string](4.0))
	assert.Equal(t, TryOr(y, f, "bad"), result.Err[float32, string]("bad"))
}

func TestTryOrElse(t *testing.T) {
	f := func(x int) float32 { return float32(x) * 2 }
	closure := func() *result.Result[float32, string] { return result.Err[float32, string]("bad") }
	x := Some(2)
	y := None[int]()

	assert.Equal(t, TryOrElse(x, f, closure), result.Ok[float32, string](4.0))
	assert.Equal(t, TryOrElse(y, f, closure), result.Err[float32, string]("bad"))
}

func TestXor(t *testing.T) {
	x := Some(2)
	y := None[int]()

	z := Some(2)
	z.Unwrap()
	assert.Equal(t, x.Xor(y), z)

	x2 := None[int]()
	y2 := Some(2)
	assert.Equal(t, x2.Xor(y2), Some(2))

	x3 := Some(2)
	y3 := Some(2)
	assert.Equal(t, x3.Xor(y3), None[int]())

	x4 := None[int]()
	y4 := None[int]()
	assert.Equal(t, x4.Xor(y4), None[int]())
}

func TestInsert(t *testing.T) {
	var opt Option[int]
	val := opt.Insert(1)
	assert.Equal(t, *val, 1)
	assert.Equal(t, opt.Unwrap(), 1)

	val2 := opt.Insert(2)
	assert.Equal(t, *val2, 2)
	*val2 = 3
	assert.Equal(t, opt.Unwrap(), 3)
}

func TestGetOrInsert(t *testing.T) {
	var x Option[int]

	y := x.GetOrInsert(5)
	assert.Equal(t, *y, 5)

	*y = 7
	assert.Equal(t, &x, Some[int](7))
}

func TestGetOrInsertDefault(t *testing.T) {
	var x Option[int]

	y := x.GetOrInsertDefault()
	assert.Equal(t, *y, 0)

	*y = 7
	assert.Equal(t, &x, Some[int](7))
}

func TestGetOrInsertWith(t *testing.T) {
	var x Option[int]

	y := x.GetOrInsertWith(func() int { return 5 })
	assert.Equal(t, *y, 5)

	*y = 7
	assert.Equal(t, &x, Some[int](7))
}

func TestTake(t *testing.T) {
	var x Option[int]
	x.Some(2)
	y := x.Take()
	assert.Equal(t, &x, None[int]())
	assert.Equal(t, y, Some(2))

	var x2 Option[int]
	x2.None()
	y2 := x2.Take()
	assert.Equal(t, &x2, None[int]())
	assert.Equal(t, y2, None[int]())
}

func TestReplace(t *testing.T) {
	var x Option[int]
	x.Some(2)
	old := x.Replace(5)
	assert.Equal(t, &x, Some(5))
	assert.Equal(t, old, Some(2))

	var x2 Option[int]
	x2.None()
	old2 := x2.Replace(3)
	assert.Equal(t, &x2, Some(3))
	assert.Equal(t, old2, None[int]())
}
