package IOEither

import (
	ET "github.com/Carsten-Leue/golang-compiler-bugs/Either"
	"github.com/Carsten-Leue/golang-compiler-bugs/IO"
	G "github.com/Carsten-Leue/golang-compiler-bugs/IOEither/Generic"
)

type IOEither[E, A any] IO.IO[ET.Either[E, A]]

func Right[E, A any](r A) IOEither[E, A] {
	return G.Right[IOEither[E, A]](r)
}
