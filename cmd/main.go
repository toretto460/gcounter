package main

import (
	"context"
	"fmt"
	"github.com/toretto460/gcounter"
)

func main() {
	ctx, done := context.WithCancel(context.Background())

	c := gcounter.Counter{}
	c.Run(ctx)

	c.Increment()
	c.Increment()
	c.Increment()
	c.Increment()
	c.Decrement()

	fmt.Print(c.GetValue())

	done()
}
