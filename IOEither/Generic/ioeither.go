package Generic

import (
	ET "github.com/Carsten-Leue/golang-compiler-bugs/Either"
	IO "github.com/Carsten-Leue/golang-compiler-bugs/IO/Generic"
	"github.com/Carsten-Leue/golang-compiler-bugs/internal/EitherT"
)

func MakeIO[GA ~func() ET.Either[E, A], E, A any](f GA) GA {
	return f
}

func Right[GA ~func() ET.Either[E, A], E, A any](r A) GA {
	return MakeIO(EitherT.Right(IO.MonadOf[GA], r))
}
