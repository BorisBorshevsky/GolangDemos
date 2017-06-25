package main

import (
	"testing"

	"github.com/BorisBorshevsky/GolangDemos/gomock/importable"
	"github.com/golang/mock/gomock"
	"github.com/k0kubun/pp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Mock")
}

var _ = Describe("", func() {
	var mockCtrl *gomock.Controller
	var swimmer *importable.MockSwimmer

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		swimmer = importable.NewMockSwimmer(mockCtrl)

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("", func() {

		call1 := swimmer.EXPECT().Say(7, "Boris ").Return("EFG", 800)
		call2 := swimmer.EXPECT().Say(8, "Boris ").Return("ABC", 700).After(call1)
		swimmer.EXPECT().Say(9, "Boris ").Return("ABC", 700).After(call2)
		abc, seven00 := swimmer.Say(7, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

		abc, seven00 = swimmer.Say(8, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

		abc, seven00 = swimmer.Say(9, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

	})


	It("", func() {

		call1 := swimmer.EXPECT().Say(7, "Boris ").Return("EFG", 800)
		call2 := swimmer.EXPECT().Say(8, "Boris ").Return("ABC", 700).After(call1)
		swimmer.EXPECT().Say(9, "Boris ").Return("ABC", 700).After(call2)
		abc, seven00 := swimmer.Say(7, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

		abc, seven00 = swimmer.Say(8, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

		abc, seven00 = swimmer.Say(9, "Boris ")
		pp.Println(abc)
		pp.Println(seven00)

	})

})
