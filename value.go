package gcounter

type GetValue struct {
	valueOut chan<- int
}

func (value GetValue) handle(c *Counter) {
	value.valueOut <- c.value
}
