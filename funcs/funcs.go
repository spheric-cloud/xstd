// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package funcs

import (
	"iter"

	"spheric.cloud/xstd/slices"
)

func Identity[V any](v V) V {
	return v
}

func Const[F, E any](e E) func(F) E {
	return func(f F) E {
		return e
	}
}

func Narrow0[In any](f func(any)) func(In) {
	return func(in In) {
		f(in)
	}
}

func Narrow[In, Out any](f func(any) Out) func(In) Out {
	return func(in In) Out {
		return f(in)
	}
}

func Narrow2[In, Out1, Out2 any](f func(any) (Out1, Out2)) func(In) (Out1, Out2) {
	return func(in In) (Out1, Out2) {
		return f(in)
	}
}

func Chain[V any](fs ...func(V) V) func(V) V {
	return ChainSeq(slices.Values(fs))
}

func ChainSeq[V any](fs iter.Seq[func(V) V]) func(V) V {
	return func(v V) V {
		for f := range fs {
			v = f(v)
		}
		return v
	}
}

func Chain2[K, V any](fs ...func(K, V) (K, V)) func(K, V) (K, V) {
	return ChainSeq2(slices.Values(fs))
}

func ChainSeq2[K, V any](fs iter.Seq[func(K, V) (K, V)]) func(K, V) (K, V) {
	return func(k K, v V) (K, V) {
		for f := range fs {
			k, v = f(k, v)
		}
		return k, v
	}
}

func Compose[In, Mid, Out any](f1 func(In) Mid, f2 func(Mid) Out) func(In) Out {
	return func(in In) Out {
		return f2(f1(in))
	}
}

func Compose2[In1, In2, Mid1, Mid2, Out1, Out2 any](f1 func(In1, In2) (Mid1, Mid2), f2 func(Mid1, Mid2) (Out1, Out2)) func(In1, In2) (Out1, Out2) {
	return func(in1 In1, in2 In2) (Out1, Out2) { return f2(f1(in1, in2)) }
}

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

func Uncurried[In, Mid, Out any](f func(In) func(Mid) Out) func(In, Mid) Out {
	return func(in In, mid Mid) Out {
		return f(in)(mid)
	}
}

func Flip[In1, In2, Out any](f func(In1, In2) Out) func(In2, In1) Out {
	return func(in2 In2, in1 In1) Out {
		return f(in1, in2)
	}
}

func Merge[In, Out1, Out2 any](f1 func(In) Out1, f2 func(In) Out2) func(In) (Out1, Out2) {
	return func(in In) (Out1, Out2) {
		out1 := f1(in)
		out2 := f2(in)
		return out1, out2
	}
}

func Split[In, Out1, Out2 any](f func(In) (Out1, Out2)) (func(In) Out1, func(In) Out2) {
	return DropValue(f), DropKey(f)
}

func LiftValueConst[In, Out1, Out2 any](f func(In) Out1, out2 Out2) func(In) (Out1, Out2) {
	return Merge(f, Const[In, Out2](out2))
}

func LiftValueZero[Out2, In, Out1 any](f func(In) Out1) func(In) (Out1, Out2) {
	return Merge(f, func(In) Out2 {
		var zero Out2
		return zero
	})
}

func LiftSuccess[In, Out any](f func(In) Out) func(In) (Out, error) {
	return LiftValueZero[error](f)
}

func LiftKeyConst[In, Out1, Out2 any](f func(In) Out2, out1 Out1) func(In) (Out1, Out2) {
	return Merge(Const[In, Out1](out1), f)
}

func LiftKeyZero[Out1, In, Out2 any](f func(In) Out2) func(In) (Out1, Out2) {
	return Merge(func(in In) Out1 {
		var zero Out1
		return zero
	}, f)
}

func DropKey[E, K, V any](f func(E) (K, V)) func(E) V {
	return func(e E) V {
		_, v := f(e)
		return v
	}
}

func DropValue[E, K, V any](f func(E) (K, V)) func(E) K {
	return func(e E) K {
		k, _ := f(e)
		return k
	}
}

func Bind[B1, E, F any](f func(B1, E) F, b1 B1) func(E) F {
	return func(e E) F {
		return f(b1, e)
	}
}

func BindFunc[B1, E, F any](f func(B1, E) F, supplier func() B1) func(E) F {
	return func(e E) F {
		return f(supplier(), e)
	}
}
