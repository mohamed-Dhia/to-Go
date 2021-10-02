package main

type Counter struct {
	count int
}

func (c *Counter) GetNextUniqueId() int {
	c.count++
	return c.count
}
