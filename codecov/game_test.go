package codecov

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCov(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cov suite")
}

var _ = Describe("Game", func() {

	var myBoard *board

	BeforeEach(func() {
		myBoard = &board{
			player1: player{name: "A"},
			player2: player{name: "B"},
		}
	})

	It("XX", func() {
		Expect(myBoard.Winner()).To(BeEquivalentTo("A"))
	})

	It("YY", func() {
		myBoard.Play(0)
		myBoard.Play(1)
		Expect(myBoard.Winner()).To(BeEquivalentTo("B"))
	})

})
