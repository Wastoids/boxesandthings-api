package service_test

import (
	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/Wastoids/boxesandthings-api/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SaveBox", func() {

	// TODO: Keep this test up to date whenever new properties are added to the box model
	When("we need to get a box from the request body", func() {

		Context("and the body is invalid", func() {
			var (
				body        string = `{"someName": "my box"}`
				box         model.Box
				expectedBox model.Box
			)

			BeforeEach(func() {
				box = service.GetBoxFromRequest(body)
			})

			It("should return an empty box model", func() {
				Expect(box).To(Equal(expectedBox))
			})
		})

		Context("and the body is valid", func() {
			var (
				body        string = `{"name": "my box"}`
				box         model.Box
				expectedBox = model.Box{Name: "my box"}
			)

			BeforeEach(func() {
				box = service.GetBoxFromRequest(body)
			})

			It("should return the expected box", func() {
				Expect(box).To(Equal(expectedBox))
			})
		})

	})
})
