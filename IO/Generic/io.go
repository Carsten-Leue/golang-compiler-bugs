package Generic

import (
	F "github.com/Carsten-Leue/golang-compiler-bugs/Function"
)

func MakeIO[GA ~func() A, A any](f func() A) GA {
	return f
}

func MonadOf[GA ~func() A, A any](a A) GA {
	return MakeIO[GA](F.Constant(a))
}
