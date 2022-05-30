package example

import "context"

type In struct{}
type Out struct{}

func DoStuff(_ context.Context, _ In) (Out, error) {
	// doing nothing, bro
	return Out{}, nil
}
