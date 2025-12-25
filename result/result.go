package result

import (
	"fmt"
)

type Result[T, E any] struct {
	value    T
	err      E
	ok       bool
	assigned bool
}

func Ok[T, E any](v T) Result[T, E] {
	return Result[T, E]{
		value:    v,
		ok:       true,
		assigned: true,
	}
}

func Err[T, E any](e E) Result[T, E] {
	return Result[T, E]{
		err:      e,
		ok:       false,
		assigned: true,
	}
}

func (r *Result[T, E]) Ok(v T) {
	if r.assigned {
		panic("Result already assigned")
	}
	r.value = v
	r.ok = true
	r.assigned = true
}

func (r *Result[T, E]) Err(e E) {
	if r.assigned {
		panic("Result already assigned")
	}
	r.err = e
	r.ok = false
	r.assigned = true
}

func From[T any](v T, err error) Result[T, error] {
	if err != nil {
		return Err[T, error](err)
	}
	return Ok[T, error](v)
}

func (r Result[T, E]) Valid() bool {
	return r.ok
}

func (r Result[T, E]) Assigned() bool {
	return r.assigned
}

func (r Result[T, E]) IsOk() bool {
	return r.ok
}

func (r Result[T, E]) IsOkAnd(f func(T) bool) bool {
	if !r.ok {
		return false
	}
	return f(r.value)
}

func (r Result[T, E]) IsErr() bool {
	return !r.ok
}

func (r Result[T, E]) IsErrAnd(f func(E) bool) bool {
	if r.ok {
		return false
	}
	return f(r.err)
}

func Map[T, E, U any](r Result[T, E], f func(T) U) Result[U, E] {
	if r.ok {
		return Ok[U, E](f(r.value))
	}
	return Err[U, E](r.err)
}

func MapOr[T, E, U any](r Result[T, E], def U, f func(T) U) U {
	if r.ok {
		return f(r.value)
	}
	return def
}

func MapOrElse[T, E, U any](r Result[T, E], def func(E) U, f func(T) U) U {
	if r.ok {
		return f(r.value)
	}
	return def(r.err)
}

func MapOrDefault[T, E, U any](r Result[T, E], f func(T) U) U {
	if r.ok {
		return f(r.value)
	}
	var zero U
	return zero
}

func MapErr[T, E, F any](r Result[T, E], f func(E) F) Result[T, F] {
	if r.ok {
		return Ok[T, F](r.value)
	}
	return Err[T, F](f(r.err))
}

func (r Result[T, E]) Inspect(f func(T)) Result[T, E] {
	if r.ok {
		f(r.value)
	}
	return r
}

func (r Result[T, E]) InspectErr(f func(E)) Result[T, E] {
	if !r.ok {
		f(r.err)
	}
	return r
}

func (r Result[T, E]) Expect(msg string) T {
	if r.ok {
		return r.value
	}
	panic(fmt.Sprintf("%s: %v", msg, r.err))
}

func (r Result[T, E]) Unwrap() T {
	if r.ok {
		return r.value
	}
	panic(fmt.Sprintf("called Result.Unwrap() on an Err value: %v", r.err))
}

func (r Result[T, E]) UnwrapOrDefault() T {
	if r.ok {
		return r.value
	}
	var zero T
	return zero
}

func (r Result[T, E]) ExpectErr(msg string) E {
	if r.ok {
		panic(fmt.Sprintf("%s: %v", msg, r.value))
	}
	return r.err
}

func (r Result[T, E]) UnwrapErr() E {
	if r.ok {
		panic(fmt.Sprintf("called Result.UnwrapErr() on an Ok value: %v", r.value))
	}
	return r.err
}

func And[T, E, U any](r Result[T, E], res Result[U, E]) Result[U, E] {
	if r.ok {
		return res
	}
	return Err[U, E](r.err)
}

func AndThen[T, E, U any](r Result[T, E], op func(T) Result[U, E]) Result[U, E] {
	if r.ok {
		return op(r.value)
	}
	return Err[U, E](r.err)
}

func Or[T, E, F any](r Result[T, E], res Result[T, F]) Result[T, F] {
	if r.ok {
		return Ok[T, F](r.value)
	}
	return res
}

func OrElse[T, E, F any](r Result[T, E], op func(E) Result[T, F]) Result[T, F] {
	if r.ok {
		return Ok[T, F](r.value)
	}
	return op(r.err)
}

func (r Result[T, E]) UnwrapOr(def T) T {
	if r.ok {
		return r.value
	}
	return def
}

func (r Result[T, E]) UnwrapOrElse(f func(E) T) T {
	if r.ok {
		return r.value
	}
	return f(r.err)
}
