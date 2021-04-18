package storage_test

import (
	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/Wastoids/boxesandthings-api/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {

	When("we need to GetTopLevelBoxesForUser", func() {

		var (
			username string
			err      error
			repo     storage.Repository
			boxes    []model.Box
		)

		Context("and there are no boxes", func() {

			BeforeEach(func() {
				username = "some_guy"
				repo = storage.NewRepository()
				boxes, err = repo.GetTopLevelBoxesForUser(username)
			})

			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return an empty list of boxes", func() {
				Expect(boxes).To(BeNil())
			})

		})

		Context("and there are a few boxes", func() {
			BeforeEach(func() {
				username = "another_guy"
				repo = storage.NewRepository()
				boxes, err = repo.GetTopLevelBoxesForUser(username)
			})

			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return a list of 3 boxes", func() {
				Expect(boxes).NotTo(BeNil())
				Expect(len(boxes)).To(Equal(3))
			})
		})

	})
})
