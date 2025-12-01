package result

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	var a = Result[int, string]{
		value:    1,
		ok:       true,
		assigned: true,
	}

	var b Result[int, string]
	b.Ok(1)

	c := Ok[int, string](1)

	assert.Equal(t, a, b)
	assert.Equal(t, a, *c)
	assert.Equal(t, b, *c)
}

func TestErr(t *testing.T) {
	var a = Result[int, string]{
		err:      "error",
		ok:       false,
		assigned: true,
	}

	var b Result[int, string]
	b.Err("error")

	c := Err[int, string]("error")

	assert.Equal(t, a, b)
	assert.Equal(t, a, *c)
	assert.Equal(t, b, *c)
}

func TestIsOk(t *testing.T) {
	r := Ok[int, string](-3)
	assert.True(t, r.IsOk())

	e := Err[int, string]("error")
	assert.False(t, e.IsOk())
}

func TestIsOkAnd(t *testing.T) {
	f1 := func(x int) bool { return x > 1 }
	f2 := func(x string) bool { return len(x) > 1 }

	r1 := Ok[int, string](2)
	assert.True(t, r1.IsOkAnd(f1))

	r2 := Ok[int, string](0)
	assert.False(t, r2.IsOkAnd(f1))

	e := Err[int, string]("error")
	assert.False(t, e.IsOkAnd(f1))

	r3 := Ok[string, string]("string")
	assert.True(t, r3.IsOkAnd(f2))
}

func TestIsErr(t *testing.T) {
	r := Ok[int, string](-3)

	assert.False(t, r.IsErr())

	e := Err[int, string]("error")
	assert.True(t, e.IsErr())
}

func TestIsErrAnd(t *testing.T) {
	var (
		ErrNotFound         = errors.New("not found")
		ErrPermissionDenied = errors.New("permission denied")
	)

	f := func(err error) bool { return errors.Is(err, ErrNotFound) }

	e1 := Err[int, error](ErrNotFound)
	assert.True(t, e1.IsErrAnd(f))

	e2 := Err[int, error](ErrPermissionDenied)
	assert.False(t, e2.IsErrAnd(f))

	r := Ok[int, error](123)
	assert.False(t, r.IsErrAnd(f))

	e3 := Err[int, string]("error")
	assert.True(t, e3.IsErrAnd(func(err string) bool { return len(err) > 1 }))
}

func TestMap(t *testing.T) {
	r := Ok[int, string](1)
	assert.Equal(t, Map(r, func(i int) string { return fmt.Sprintf("%d", i) }), Ok[string, string]("1"))
}

func TestMapOr(t *testing.T) {
	f := func(v string) int { return len(v) }

	r := Ok[string, string]("foo")
	assert.Equal(t, MapOr(r, 42, f), 3)

	e := Err[string, string]("bar")
	assert.Equal(t, MapOr(e, 42, f), 42)
}

func TestMapOrElse(t *testing.T) {
	k := 21
	f1 := func(e string) int { return k * 2 }
	f2 := func(v string) int { return len(v) }

	r := Ok[string, string]("foo")
	assert.Equal(t, MapOrElse(r, f1, f2), 3)

	e := Err[string, string]("bar")
	assert.Equal(t, MapOrElse(e, f1, f2), 42)
}

func TestMapOrDefault(t *testing.T) {
	f := func(v string) int { return len(v) }

	r := Ok[string, string]("foo")
	assert.Equal(t, MapOrDefault(r, f), 3)

	e := Err[string, string]("bar")
	assert.Equal(t, MapOrDefault(e, f), 0)
}

func TestMapErr(t *testing.T) {
	stringify := func(x int) string { return fmt.Sprintf("%d", x) }

	r := Ok[int, int](2)
	assert.Equal(t, MapErr(r, stringify), Ok[int, string](2))

	e := Err[int, int](13)
	assert.Equal(t, MapErr(e, stringify), Err[int, string]("13"))
}

func TestInspect(t *testing.T) {
	f := func(x int) { fmt.Println(x) }

	r := Ok[int, string](1)
	assert.Equal(t, r.Inspect(f), r)
}

func TestInspectErr(t *testing.T) {
	f := func(x string) { fmt.Println(x) }

	e := Err[int, string]("error")
	assert.Equal(t, e.InspectErr(f), e)
}

func TestExpect(t *testing.T) {
	e := Err[int, string]("error")
	assert.Panics(t, func() { e.Expect("error") })
}

func TestUnwrap(t *testing.T) {
	r := Ok[int, string](2)
	assert.Equal(t, r.Unwrap(), 2)

	e := Err[int, string]("error")
	assert.Panics(t, func() { e.Unwrap() })
}

func TestUnwrapOrDefault(t *testing.T) {
	parseYear := func(s string) *Result[int, string] {
		year, err := strconv.Atoi(s)
		if err != nil {
			return Err[int, string]("parseYear failed")
		}
		return Ok[int, string](year)
	}

	good := parseYear("1909")
	bad := parseYear("190blarg")

	assert.Equal(t, good.UnwrapOrDefault(), 1909)
	assert.Equal(t, bad.UnwrapOrDefault(), 0)
	assert.Equal(t, bad.UnwrapErr(), "parseYear failed")
}

func TestExpectErr(t *testing.T) {
	r := Ok[int, string](10)
	assert.Panics(t, func() { r.ExpectErr("error") })
}

func TestUnwrapErr(t *testing.T) {
	e := Err[int, string]("error")
	assert.Equal(t, e.UnwrapErr(), "error")
}

func TestAnd(t *testing.T) {
	x1 := Ok[int, string](2)
	y1 := Err[string, string]("late error")
	assert.Equal(t, And(x1, y1), y1)

	x2 := Err[int, string]("early error")
	y2 := Ok[string, string]("foo")
	assert.Equal(t, And(x2, y2), Err[string, string]("early error"))

	x3 := Err[int, string]("not a 2")
	y3 := Err[string, string]("late error")
	assert.Equal(t, And(x3, y3), Err[string, string]("not a 2"))

	x4 := Ok[int, string](2)
	y4 := Err[string, string]("different result type")
	assert.Equal(t, And(x4, y4), y4)
}

func TestAndThen(t *testing.T) {
	sqThenToString := func(x int) *Result[string, string] {
		if x > 100 || x < 0 {
			return Err[string, string]("overflowed")
		}
		return Ok[string, string](strconv.Itoa(x))
	}

	r1 := Ok[int, string](2)
	assert.Equal(t, AndThen(r1, sqThenToString), Ok[string, string]("2"))

	r2 := Ok[int, string](1000)
	assert.Equal(t, AndThen(r2, sqThenToString), Err[string, string]("overflowed"))

	r3 := Err[int, string]("error")
	assert.Equal(t, AndThen(r3, sqThenToString), Err[string, string]("error"))
}

func TestOr(t *testing.T) {
	x1 := Ok[int, int](2)
	y1 := Err[int, string]("late error")
	assert.Equal(t, Or(x1, y1), Ok[int, string](2))

	x2 := Err[int, string]("early error")
	y2 := Ok[int, int](2)

	assert.Equal(t, Or(x2, y2), y2)

	x3 := Err[int, string]("not a 2")
	y3 := Err[int, string]("late error")
	assert.Equal(t, Or(x3, y3), y3)

	x4 := Ok[int, string](2)
	y4 := Ok[int, int](100)
	assert.Equal(t, Or(x4, y4), Ok[int, int](2))
}

func TestOrElse(t *testing.T) {
	sq := func(x int) *Result[int, int] { return Ok[int, int](x * x) }
	err := func(x int) *Result[int, int] { return Err[int, int](x) }

	r := Ok[int, int](2)
	e := Err[int, int](3)

	assert.Equal(t, OrElse(OrElse(r, sq), sq), Ok[int, int](2))
	assert.Equal(t, OrElse(OrElse(r, err), sq), Ok[int, int](2))
	assert.Equal(t, OrElse(OrElse(e, sq), err), Ok[int, int](9))
	assert.Equal(t, OrElse(OrElse(e, err), err), Err[int, int](3))
}

func TestTry(t *testing.T) {
	f1 := func(x int) float32 { return float32(x) * 2 }
	closure := func(e rune) *Result[float32, string] { return Err[float32, string](string(e)) }

	r := Ok[int, rune](1)
	assert.Equal(t, Try(r, f1, closure), Ok[float32, string](2.0))

	e := Err[int, rune]('E')
	assert.Equal(t, Try(e, f1, closure), Err[float32, string]("E"))
}

func TestUnwrapOr(t *testing.T) {
	r := Ok[int, string](9)
	assert.Equal(t, r.UnwrapOr(2), 9)

	e := Err[int, string]("error")
	assert.Equal(t, e.UnwrapOr(2), 2)
}

func TestUnwrapOrElse(t *testing.T) {
	f := func(x string) int { return len(x) }

	r := Ok[int, string](2)
	assert.Equal(t, r.UnwrapOrElse(f), 2)

	e := Err[int, string]("foo")
	assert.Equal(t, e.UnwrapOrElse(f), 3)
}

// 新增的测试用例，提高覆盖率

func TestResultWithComplexTypes(t *testing.T) {
	// 测试复杂类型
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	r := Ok[Person, string](p)
	assert.True(t, r.IsOk())
	assert.Equal(t, r.Unwrap().Name, "Alice")
	assert.Equal(t, r.Unwrap().Age, 30)

	e := Err[Person, string]("person not found")
	assert.True(t, e.IsErr())
	assert.Equal(t, e.UnwrapErr(), "person not found")
}

func TestChainedOperations(t *testing.T) {
	// 测试链式操作
	parse := func(s string) *Result[int, string] {
		i, err := strconv.Atoi(s)
		if err != nil {
			return Err[int, string]("invalid number")
		}
		return Ok[int, string](i)
	}

	double := func(x int) *Result[int, string] {
		if x > 100 {
			return Err[int, string]("too large")
		}
		return Ok[int, string](x * 2)
	}

	// 成功链
	result := AndThen(parse("42"), double)
	assert.True(t, result.IsOk())
	assert.Equal(t, result.Unwrap(), 84)

	// 失败链
	result2 := AndThen(parse("200"), double)
	assert.True(t, result2.IsErr())
	assert.Equal(t, result2.UnwrapErr(), "too large")
}

func TestResultWithNilValues(t *testing.T) {
	// 测试 nil 值处理
	var s *string
	r := Ok[*string, string](s)
	assert.True(t, r.IsOk())
	assert.Nil(t, r.Unwrap())

	var p *string = new(string)
	*p = "test"
	r2 := Ok[*string, string](p)
	assert.True(t, r2.IsOk())
	assert.Equal(t, *r2.Unwrap(), "test")
}

func TestResultWithFunctions(t *testing.T) {
	// 测试函数作为值
	f := func() int { return 42 }
	r := Ok[func() int, string](f)
	assert.True(t, r.IsOk())
	assert.Equal(t, r.Unwrap()(), 42)

	r2 := Err[func() int, string]("function error")
	assert.True(t, r2.IsErr())
	assert.Equal(t, r2.UnwrapErr(), "function error")
}

func TestResultWithSlices(t *testing.T) {
	// 测试切片
	s := []int{1, 2, 3}
	r := Ok[[]int, string](s)
	assert.True(t, r.IsOk())
	assert.Equal(t, len(r.Unwrap()), 3)

	empty := []int{}
	r2 := Ok[[]int, string](empty)
	assert.True(t, r2.IsOk())
	assert.Equal(t, len(r2.Unwrap()), 0)

	r3 := Err[[]int, string]("slice error")
	assert.True(t, r3.IsErr())
	assert.Equal(t, r3.UnwrapErr(), "slice error")
}

func TestResultWithMaps(t *testing.T) {
	// 测试 map
	m := map[string]int{"a": 1, "b": 2}
	r := Ok[map[string]int, string](m)
	assert.True(t, r.IsOk())
	assert.Equal(t, r.Unwrap()["a"], 1)
	assert.Equal(t, r.Unwrap()["b"], 2)

	r2 := Err[map[string]int, string]("map error")
	assert.True(t, r2.IsErr())
	assert.Equal(t, r2.UnwrapErr(), "map error")
}

func TestResultWithChannels(t *testing.T) {
	// 测试通道
	ch := make(chan int)
	r := Ok[chan int, string](ch)
	assert.True(t, r.IsOk())
	assert.NotNil(t, r.Unwrap())

	r2 := Err[chan int, string]("channel error")
	assert.True(t, r2.IsErr())
	assert.Equal(t, r2.UnwrapErr(), "channel error")
}

func TestResultWithInterfaces(t *testing.T) {
	// 测试接口
	var r Result[any, string]
	r.Ok("test value")
	assert.True(t, r.IsOk())
	assert.Equal(t, r.Unwrap(), "test value")

	var r2 Result[any, string]
	r2.Err("interface error")
	assert.True(t, r2.IsErr())
	assert.Equal(t, r2.UnwrapErr(), "interface error")
}
