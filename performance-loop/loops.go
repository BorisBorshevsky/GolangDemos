package performance_loop

type someStruct struct {
	ID    int
	Score float64
}

func loopOutside(n int) (res []someStruct) {

	s := someStruct{}
	for i := 0; i < n; i++ {
		s.ID = 1
		s.Score = 100
		res = append(res, s)
	}

	return
}

func loopInside(n int) (res []someStruct) {

	for i := 0; i < n; i++ {
		s := someStruct{}
		s.ID = 1
		s.Score = 100
		res = append(res, s)
	}

	return
}

func loopInline(n int) (res []someStruct) {

	for i := 0; i < n; i++ {
		res = append(res, someStruct{ID: 1, Score: 100})
	}

	return
}

func loopPoiner(n int) (res []*someStruct) {

	for i := 0; i < n; i++ {
		res = append(res, &someStruct{ID: 1, Score: 100})
	}

	return
}
