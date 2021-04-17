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

		Context("and it fails", func() {

			BeforeEach(func() {
				username = "non_existent_username"
				repo = storage.NewRepository()
				boxes, err = repo.GetTopLevelBoxesForUser(username)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should return an empty list of boxes", func() {
				Expect(boxes).To(BeNil())
			})

		})

		XContext("and it succeeds", func() {

		})

	})
})
