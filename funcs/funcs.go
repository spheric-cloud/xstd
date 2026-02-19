// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package funcs

import (
	"iter"
	"slices"
)

// Identity returns the value unchanged. It is the identity function.
func Identity[V any](v V) V {
	return v
}

// Const returns a function that always returns the given value, ignoring its argument.
func Const[F, E any](e E) func(F) E {
	return func(f F) E {
		return e
	}
}

// CastIn0 wraps a function by casting its input parameter from NewIn to OldIn via type assertion.
func CastIn0[NewIn, OldIn any](f func(OldIn)) func(NewIn) {
	return func(v NewIn) {
		f(any(v).(OldIn))
	}
}

// CastIn wraps a function by casting its input parameter from NewIn to OldIn via type assertion.
func CastIn[NewIn, OldIn, Out any](f func(OldIn) Out) func(NewIn) Out {
	return func(v NewIn) Out {
		return f(any(v).(OldIn))
	}
}

// CastIn2 wraps a function by casting its input parameter from NewIn to OldIn via type assertion,
// for functions that return two values.
func CastIn2[NewIn, OldIn, Out1, Out2 any](f func(OldIn) (Out1, Out2)) func(NewIn) (Out1, Out2) {
	return func(v NewIn) (Out1, Out2) {
		return f(any(v).(OldIn))
	}
}

// CastOut0 wraps a zero-argument function by casting its output from OldOut to NewOut via type assertion.
func CastOut0[NewOut, OldOut any](f func() OldOut) func() NewOut {
	return func() NewOut {
		return any(f()).(NewOut)
	}
}

// CastOut wraps a function by casting its output from OldOut to NewOut via type assertion.
func CastOut[NewOut, OldOut, In any](f func(In) OldOut) func(In) NewOut {
	return func(v In) NewOut {
		return any(f(v)).(NewOut)
	}
}

// CastOut2 wraps a two-argument function by casting its output from OldOut to NewOut via type assertion.
func CastOut2[NewOut, OldOut, In1, In2 any](f func(In1, In2) OldOut) func(In1, In2) NewOut {
	return func(v1 In1, v2 In2) NewOut {
		return any(f(v1, v2)).(NewOut)
	}
}

// CastInOut wraps a function by casting both its input and output via type assertions.
func CastInOut[NewIn, NewOut, OldIn, OldOut any](f func(OldIn) OldOut) func(NewIn) NewOut {
	return func(v NewIn) NewOut {
		return any(f(any(v).(OldIn))).(NewOut)
	}
}

// Narrow0 narrows a function that takes any to a function that takes a specific type.
func Narrow0[In any](f func(any)) func(In) {
	return func(in In) {
		f(in)
	}
}

// Narrow narrows a function that takes any to a function that takes a specific type, preserving the output.
func Narrow[In, Out any](f func(any) Out) func(In) Out {
	return func(in In) Out {
		return f(in)
	}
}

// Narrow2 narrows a function that takes any to a function that takes a specific type,
// preserving two output values.
func Narrow2[In, Out1, Out2 any](f func(any) (Out1, Out2)) func(In) (Out1, Out2) {
	return func(in In) (Out1, Out2) {
		return f(in)
	}
}

// Chain returns a function that applies all given functions in sequence, passing the output of each
// as input to the next.
func Chain[V any](fs ...func(V) V) func(V) V {
	return ChainSeq(slices.Values(fs))
}

// ChainSeq returns a function that applies all functions from the iterator in sequence.
func ChainSeq[V any](fs iter.Seq[func(V) V]) func(V) V {
	return func(v V) V {
		for f := range fs {
			v = f(v)
		}
		return v
	}
}

// Chain2 returns a function that applies all given two-argument functions in sequence.
func Chain2[K, V any](fs ...func(K, V) (K, V)) func(K, V) (K, V) {
	return ChainSeq2(slices.Values(fs))
}

// ChainSeq2 returns a function that applies all two-argument functions from the iterator in sequence.
func ChainSeq2[K, V any](fs iter.Seq[func(K, V) (K, V)]) func(K, V) (K, V) {
	return func(k K, v V) (K, V) {
		for f := range fs {
			k, v = f(k, v)
		}
		return k, v
	}
}

// Compose composes two functions, applying f1 first and then f2 to the result.
func Compose[In, Mid, Out any](f1 func(In) Mid, f2 func(Mid) Out) func(In) Out {
	return func(in In) Out {
		return f2(f1(in))
	}
}

// Compose2 composes two functions that take and return pairs, applying f1 first and then f2.
func Compose2[In1, In2, Mid1, Mid2, Out1, Out2 any](f1 func(In1, In2) (Mid1, Mid2), f2 func(Mid1, Mid2) (Out1, Out2)) func(In1, In2) (Out1, Out2) {
	return func(in1 In1, in2 In2) (Out1, Out2) { return f2(f1(in1, in2)) }
}

// ComposeErr composes two functions that can return errors, short-circuiting on the first error.
func ComposeErr[In, Mid, Out any](f1 func(In) (Mid, error), f2 func(Mid) (Out, error)) func(In) (Out, error) {
	return func(in In) (Out, error) {
		mid, err := f1(in)
		if err != nil {
			var zero Out
			return zero, err
		}
		return f2(mid)
	}
}

// Uncurried converts a curried function into an uncurried two-argument function.
func Uncurried[In, Mid, Out any](f func(In) func(Mid) Out) func(In, Mid) Out {
	return func(in In, mid Mid) Out {
		return f(in)(mid)
	}
}

// Flip swaps the arguments of a two-argument function.
func Flip[In1, In2, Out any](f func(In1, In2) Out) func(In2, In1) Out {
	return func(in2 In2, in1 In1) Out {
		return f(in1, in2)
	}
}

// Merge combines two functions with the same input into a single function that returns both outputs.
func Merge[In, Out1, Out2 any](f1 func(In) Out1, f2 func(In) Out2) func(In) (Out1, Out2) {
	return func(in In) (Out1, Out2) {
		out1 := f1(in)
		out2 := f2(in)
		return out1, out2
	}
}

// Split splits a function that returns two values into two functions that each return one value.
func Split[In, Out1, Out2 any](f func(In) (Out1, Out2)) (func(In) Out1, func(In) Out2) {
	return DropValue(f), DropKey(f)
}

// LiftValueConst lifts a function to return a pair, where the second value is a constant.
func LiftValueConst[In, Out1, Out2 any](f func(In) Out1, out2 Out2) func(In) (Out1, Out2) {
	return Merge(f, Const[In, Out2](out2))
}

// LiftValueZero lifts a function to return a pair, where the second value is the zero value.
func LiftValueZero[Out2, In, Out1 any](f func(In) Out1) func(In) (Out1, Out2) {
	return Merge(f, func(In) Out2 {
		var zero Out2
		return zero
	})
}

// LiftSuccess lifts a function to return a (value, error) pair with a nil error.
func LiftSuccess[In, Out any](f func(In) Out) func(In) (Out, error) {
	return LiftValueZero[error](f)
}

// LiftKeyConst lifts a function to return a pair, where the first value is a constant.
func LiftKeyConst[In, Out1, Out2 any](f func(In) Out2, out1 Out1) func(In) (Out1, Out2) {
	return Merge(Const[In, Out1](out1), f)
}

// LiftKeyZero lifts a function to return a pair, where the first value is the zero value.
func LiftKeyZero[Out1, In, Out2 any](f func(In) Out2) func(In) (Out1, Out2) {
	return Merge(func(in In) Out1 {
		var zero Out1
		return zero
	}, f)
}

// DropKey wraps a function that returns (K, V) to return only V, discarding the first return value.
func DropKey[E, K, V any](f func(E) (K, V)) func(E) V {
	return func(e E) V {
		_, v := f(e)
		return v
	}
}

// DropValue wraps a function that returns (K, V) to return only K, discarding the second return value.
func DropValue[E, K, V any](f func(E) (K, V)) func(E) K {
	return func(e E) K {
		k, _ := f(e)
		return k
	}
}

// Bind partially applies the first argument of a two-argument function.
func Bind[B1, E, F any](f func(B1, E) F, b1 B1) func(E) F {
	return func(e E) F {
		return f(b1, e)
	}
}

// BindFunc partially applies the first argument of a two-argument function using a supplier function.
func BindFunc[B1, E, F any](f func(B1, E) F, supplier func() B1) func(E) F {
	return func(e E) F {
		return f(supplier(), e)
	}
}
