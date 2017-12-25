package importable

import "strings"

type Swimmer interface {
	Swim(speed float64) string
	Splash(times int) int
	Say(times int, what string) (string, int)
}

type swimmer struct{}

func (s *swimmer) Swim(speed float64) string {
	return "Done"
}

func (s *swimmer) Splash(times int) int {
	return times * 2
}

func (s *swimmer) Say(times int, what string) (string, int) {
	return strings.Repeat(what, times), 100
}
