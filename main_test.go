package golangDemos

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolangDemos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GolangDemos Suite")
}

var _ = Describe("my test", func() {
	var adder Adder

	BeforeEach(func() {
		adder = Adder{}
	})

	It("adds 5", func() {
		adder.Add(5)
		Expect(adder.Sum()).To(BeEquivalentTo(5))
	})

	It("adds 5", func() {
		actual := func() {
			adder.Add(7)
		}
		Expect(actual).To(Panic())
	})

})
