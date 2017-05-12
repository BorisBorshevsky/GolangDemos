package codecov

type board struct {
	player1 player
	player2 player
	turn    int
}

type player struct {
	name  string
	score int
}

type Game interface {
	Play(i int)
	Winner() string
}

func NewGame(p1, p2 string) Game {
	return &board{
		player1: player{name: p1},
		player2: player{name: p2},
	}
}

func (b *board) Winner() string {
	if b.player1.score >= b.player2.score {
		return b.player1.name
	}

	return b.player2.name
}

func (b *board) Play(i int) {
	if b.turn%2 == 0 {
		b.player1.score += i
	} else {
		b.player2.score += i

	}
	b.turn++
}
