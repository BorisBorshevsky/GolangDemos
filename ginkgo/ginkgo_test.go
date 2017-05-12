package ginkgo

import (
	"testing"

	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Getter Suite")
}

var _ = BeforeSuite(func() {
	log.Println("before suit #1")
})

var _ = AfterSuite(func() {
	log.Println("after suit #1")
})

var _ = Describe("Describe 1", func() {

	BeforeEach(func() {
		log.Println("before each external #1")
	})

	Context("context #1.1", func() {
		log.Println("in Context #1.1")

		AfterEach(func() {
			log.Println("after each context #1.1")
		})

		BeforeEach(func() {
			log.Println("before each context #1.1")
		})

		JustBeforeEach(func() {
			log.Println("just before each context #1.1")
		})

		It("expect #1.1.1", func() {
			log.Println("test #1.1.1")
			By("test 1")
			Expect(1).To(Equal(1))
			By("test 2")
			Expect(2).To(Equal(2))
		})

		It("expect #1.1.2", func() {
			log.Println("test #1.1.2")
			Expect(1).To(Equal(1))

		})

	})

	Context("context #1.2", func() {
		log.Println("in Context #1.2")
		JustBeforeEach(func() {
			log.Println("just before each context #1.2")
		})

		BeforeEach(func() {
			log.Println("before each context #1.2")
		})

		It("expect #1.2.1", func() {
			log.Println("test #1.2.1")
			Expect(1).To(Equal(1))
		})

		It("expect #1.2.2", func() {
			log.Println("test #1.2.2")
			Expect(1).To(Equal(1))
		})

		Context("context #1.2 inner", func() {
			log.Println("in Context #1.2 inner")

			BeforeEach(func() {
				log.Println("before each context #1.2 inner")
			})

			It("expect #1.2.2 inner", func() {
				log.Println("test #1.2.2 inner")
				Expect(1).To(Equal(1))
			})
		})

	})

})

var _ = Describe("Describe2", func() {
	BeforeEach(func() {
		log.Println("before each external #2")
	})

	Context("context #2.1", func() {
		log.Println("in Context #2.1")

		BeforeEach(func() {
			log.Println("before each context #2.1")
		})

		JustBeforeEach(func() {
			log.Println("just before each context #2.1")
		})

		It("expect #2.1.1", func() {
			log.Println("test #2.1.1")
			Expect(1).To(Equal(1))
		})

		It("expect #2.1.2", func() {
			log.Println("test #2.1.2")
			Expect(1).To(Equal(1))
		})

	})
})
