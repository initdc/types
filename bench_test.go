package bench_test

import (
	"testing"

	. "github.com/initdc/types/option"
	. "github.com/initdc/types/result"
)

func BenchmarkOptionSome(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var o = new(Option[int32])
		o.Some(0)
	}

	b.StopTimer()
}

func BenchmarkOptionNone(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var o = new(Option[int32])
		o.None()
	}

	b.StopTimer()
}

func BenchmarkOptionSomeArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32

	for i := 0; i < b.N; i++ {
		var o = new(Option[*[8192]int32])
		o.Some(&a)
	}

	b.StopTimer()
}

func BenchmarkOptionNoneArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var _ [8192]int32

	for i := 0; i < b.N; i++ {
		var o = new(Option[[8192]int32])
		o.None()
	}

	b.StopTimer()
}

func BenchmarkSome(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Some[int32](0)
	}

	b.StopTimer()
}

func BenchmarkNone(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		None[int32]()
	}

	b.StopTimer()
}

func BenchmarkSomeArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32

	for i := 0; i < b.N; i++ {
		Some[*[8192]int32](&a)
	}

	b.StopTimer()
}

func BenchmarkNoneArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var _ [8192]int32

	for i := 0; i < b.N; i++ {
		None[[8192]int32]()
	}

	b.StopTimer()
}

func BenchmarkResultOk(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		var r = new(Result[int32, int32])
		r.Ok(0)
	}

	b.StopTimer()
}

func BenchmarkResultErr(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		var r = new(Result[int32, int32])
		r.Err(0)
	}

	b.StopTimer()
}

func BenchmarkResultOkArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32

	for i := 0; i < b.N; i++ {
		var r = new(Result[*[8192]int32, *[8192]int32])
		r.Ok(&a)
	}

	b.StopTimer()
}

func BenchmarkResultErrArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32

	for i := 0; i < b.N; i++ {
		var r = new(Result[*[8192]int32, *[8192]int32])
		r.Err(&a)
	}

	b.StopTimer()
}

func BenchmarkOk(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Ok[int32, int32](0)
	}

	b.StopTimer()
}

func BenchmarkErr(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Err[int32, int32](0)
	}

	b.StopTimer()
}

func BenchmarkOkArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32
	for i := 0; i < b.N; i++ {
		Ok[*[8192]int32, *[8192]int32](&a)
	}

	b.StopTimer()
}

func BenchmarkErrArray(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()

	var a [8192]int32
	for i := 0; i < b.N; i++ {
		Err[*[8192]int32, *[8192]int32](&a)
	}

	b.StopTimer()
}
