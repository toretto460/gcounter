package gcounter

type Decrement struct{}

func (dec Decrement) handle(c *Counter) {
	c.value--
}
