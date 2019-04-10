package gcounter

type Increment struct{}

func (inc Increment) handle(c *Counter) {
	c.value++
}
