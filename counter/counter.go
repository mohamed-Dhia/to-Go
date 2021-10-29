package counter

type Counter struct {
	count int
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) GetNextUniqueId() int {
	c.count++
	return c.count
}
