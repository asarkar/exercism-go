package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Given a function, a list, and an initial accumulator, Foldr
// reduces each item into the accumulator from the left.
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	return foldl(s, fn, initial)
}

// Given a function, a list, and an initial accumulator, Foldr
// reduces each item into the accumulator from the right.
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	return foldr(s, fn, initial)
}

// Generic FP-style foldl.
func foldl[T any, R any](xs []T, f func(R, T) R, z R) R {
	acc := z
	for i := range xs {
		acc = f(acc, xs[i])
	}
	return acc
}

// foldr(f, z, [1,2,3]):

// g3 --> g2 --> g1 --> g0 --> id(z)
//  |      |      |      |
//  f(1)  f(2)  f(3)   returns z=0

// Evaluation order (right fold):
// g3(z)
//   -> g2(f(1, z))
//       -> g1(f(2, f(1, z)))
//           -> g0(f(3, f(2, f(1, z))))
//               -> id(f(3, f(2, f(1, z)))) = result

// * Each gN wraps the previous g(N-1).
// * Applying g3(z) evaluates right-to-left, producing f(1, f(2, f(3, z))).
func foldr[T any, R any](xs []T, f func(T, R) R, z R) R {
	type acc func(R) R

	// Identity function.
	id := func(acc R) R { return acc }

	composed := foldl(xs, func(g acc, x T) acc {
		return func(acc R) R {
			return g(f(x, acc))
		}
	}, id)

	// Apply composed function to initial value.
	return composed(z)
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, 0, len(s))
	return foldl(s, func(acc IntList, i int) IntList {
		if fn(i) {
			return append(acc, i)
		}
		return acc
	}, res)
}

func (s IntList) Length() int {
	return s.Foldl(func(acc, _ int) int { return acc + 1 }, 0)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, 0, len(s))
	return foldl(s, func(acc IntList, i int) IntList {
		return append(acc, fn(i))
	}, res)
}

func (s IntList) Reverse() IntList {
	res := make(IntList, 0, len(s))
	return foldr(s, func(i int, acc IntList) IntList {
		return append(acc, i)
	}, res)
}

func (s IntList) Append(lst IntList) IntList {
	return foldl(lst, func(acc IntList, i int) IntList {
		return append(acc, i)
	}, s)
}

func (s IntList) Concat(lists []IntList) IntList {
	// 1: Compute total length.
	totalLen := len(s)
	for i := range lists {
		totalLen += len(lists[i])
	}

	// 2: Allocate final slice.
	result := make(IntList, 0, totalLen)
	result = append(result, s...)

	// 3: Copy all lists into result.
	for _, lst := range lists {
		result = append(result, lst...)
	}

	return result
}
