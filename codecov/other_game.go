package codecov

type differenBoard struct {
	player1 otherPlayer
	player2 otherPlayer
	turn    int
}

type otherPlayer struct {
	name  string
	score int
}

type Match interface {
	Play(i int)
	Winner() string
}

func NewMatch(p1, p2 string) Match {
	return &differenBoard{
		player1: otherPlayer{name: p1},
		player2: otherPlayer{name: p2},
	}
}

func (b *differenBoard) Winner() string {
	if b.player1.score >= b.player2.score {
		return b.player1.name
	}

	return b.player2.name
}

func (b *differenBoard) Play(i int) {
	if b.turn%2 == 0 {
		b.player1.score += i
	} else {
		b.player2.score += i

	}
	b.turn++
}
