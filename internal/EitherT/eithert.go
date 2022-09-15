package EitherT

import (
	ET "github.com/Carsten-Leue/golang-compiler-bugs/Either"
	F "github.com/Carsten-Leue/golang-compiler-bugs/Function"
)

func Right[E, A, HKTA any](fof func(ET.Either[E, A]) HKTA, a A) HKTA {
	return F.Pipe2(a, ET.Right[E, A], fof)
}
