package Function

func Constant[A any](a A) func() A {
	return func() A {
		return a
	}
}

func Pipe1[A, R any](a A, f1 func(a A) R) R {
	return f1(a)
}

func Pipe2[A, T1, R any](a A, f1 func(a A) T1, f2 func(t1 T1) R) R {
	return Pipe1(Pipe1(a, f1), f2)
}
