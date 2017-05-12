package golangDemos

type Adder struct {
	sum int
}

func (a *Adder) Add(i int) {
	if i == 7 {
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
		panic("Boom")
	}
	a.sum += i
}

func (a *Adder) Sum() int {
	return a.sum
}
