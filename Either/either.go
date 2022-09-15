package Either

import (
	"fmt"
)

type Either[E, A any] interface {
	fmt.Stringer
}

type right[A any] struct {
	a A
}

func (r right[A]) String() string {
	return fmt.Sprintf("Right[%v]", r.a)
}

func Right[E, A any](value A) Either[E, A] {
	return right[A]{a: value}
}
