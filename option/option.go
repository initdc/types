package option

import (
	"github.com/initdc/types/result"
)

type Option[T any] struct {
	value    T
	some     bool
	assigned bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		value:    v,
		some:     true,
		assigned: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some:     false,
		assigned: true,
	}
}


func (o *Option[T]) Some(v T) {
	if o.assigned {
		panic("Option already assigned")
	}
	o.value = v
	o.some = true
	o.assigned = true
}

func (o *Option[T]) None() {
	if o.assigned {
		panic("Option already assigned")
	}
	o.some = false
	o.assigned = true
}

func (o Option[T]) Valid() bool {
	return o.some
}

func (o Option[T]) Assigned() bool {
	return o.assigned
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsSomeAnd(f func(T) bool) bool {
	if !o.some {
		return false
	}
	return f(o.value)
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) IsNoneOr(f func(T) bool) bool {
	if !o.some {
		return true
	}
	return f(o.value)
}

func (o Option[T]) Expect(msg string) T {
	if o.some {
		return o.value
	}
	panic(msg)
}

func (o Option[T]) Unwrap() T {
	if o.some {
		return o.value
	}
	panic("called Option.Unwrap() on a None value")
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.value
	}
	return def
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.some {
		return o.value
	}
	return f()
}

func (o Option[T]) UnwrapOrDefault() T {
	if o.some {
		return o.value
	}
	var zero T
	return zero
}

func Map[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.some {
		return Some[U](f(o.value))
	}
	return None[U]()
}

func (o Option[T]) Inspect(f func(T)) Option[T] {
	if o.some {
		f(o.value)
	}
	return o
}

func MapOr[T, U any](o Option[T], def U, f func(T) U) U {
	if o.some {
		return f(o.value)
	}
	return def
}

func MapOrElse[T, U any](o Option[T], def func() U, f func(T) U) U {
	if o.some {
		return f(o.value)
	}
	return def()
}

func MapOrDefault[T, U any](o Option[T], f func(T) U) U {
	if o.some {
		return f(o.value)
	}
	var zero U
	return zero
}

func OkOr[T, E any](o Option[T], err E) result.Result[T, E] {
	if o.some {
		return result.Ok[T, E](o.value)
	}
	return result.Err[T, E](err)
}

func OkOrElse[T, E any](o Option[T], err func() E) result.Result[T, E] {
	if o.some {
		return result.Ok[T, E](o.value)
	}
	return result.Err[T, E](err())
}

func And[T, U any](o Option[T], optb Option[U]) Option[U] {
	if o.some {
		return optb
	}
	return None[U]()
}

func AndThen[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.some {
		return f(o.value)
	}
	return None[U]()
}

func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.some && predicate(o.value) {
		return Some[T](o.value)
	}
	return None[T]()
}

func (o Option[T]) Or(optb Option[T]) Option[T] {
	if o.some {
		return o
	}
	return optb
}

func (o Option[T]) OrElse(f func() Option[T]) Option[T] {
	if o.some {
		return o
	}
	return f()
}

func (o Option[T]) Xor(optb Option[T]) Option[T] {
	if o.some && !optb.some {
		return o
	}
	if !o.some && optb.some {
		return optb
	}
	return None[T]()
}

func (o *Option[T]) Insert(value T) *T {
	o.value = value
	o.some = true
	o.assigned = true
	return &o.value
}

func (o *Option[T]) GetOrInsert(value T) *T {
	if !o.some {
		o.value = value
		o.some = true
		o.assigned = true
	}
	return &o.value
}

func (o *Option[T]) GetOrInsertDefault() *T {
	if !o.some {
		var zero T
		o.value = zero
		o.some = true
		o.assigned = true
	}
	return &o.value
}

func (o *Option[T]) GetOrInsertWith(f func() T) *T {
	if !o.some {
		o.value = f()
		o.some = true
		o.assigned = true
	}
	return &o.value
}

func (o *Option[T]) Take() Option[T] {
	var cp = *o
	*o = None[T]()
	return cp
}

func (o *Option[T]) Replace(value T) Option[T] {
	var cp = *o
	*o = Some[T](value)
	return cp
}
