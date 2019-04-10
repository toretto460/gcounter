# Go Counter Process

The go counter process is a gen-server like process implementation; it's inspired by the CodeSync 2019 Talk [Go VS Elixir, a concurrency comparison](https://codesync.global/media/go-vs-elixir-a-concurrency-comparison) by [aneyzberg](https://github.com/aneyzberg) and [hannahhoward](https://github.com/hannahhoward)

The objective of that code is to model a very simple gen-server example in go without covering all the OTP features.

Just start from the main.go

```go
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
```