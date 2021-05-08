package storage_test

import (
	"github.com/Wastoids/boxesandthings-api/model"
	"github.com/Wastoids/boxesandthings-api/service"
	"github.com/Wastoids/boxesandthings-api/storage"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO: Make tests more exhaustive
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

	When("we need to SaveBox", func() {
		var (
			b = model.Box{ID: uuid.NewString(), Name: "my box"}
		)

		Context("and there is no error", func() {

			var (
				err error
			)

			BeforeEach(func() {
				repo := storage.NewRepository()
				err = repo.SaveBox(b)
			})

			// TODO: When you add a method to get a specific box detail
			// assert that the box actually got saved
			It("should not cause an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	When("we need to SaveThing", func() {

		Context("and there is no error", func() {
			var (
				t   model.Thing
				err error
			)
			BeforeEach(func() {
				t = model.Thing{ID: uuid.NewString(), Name: "my thing", Description: "my description"}
				err = storage.NewRepository().SaveThing(t, uuid.NewString())
			})

			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	When("we need to GetBoxContents", func() {

		var (
			err            error
			box            = model.Box{ID: uuid.NewString(), Name: "my box"}
			firstThing     = model.Thing{ID: uuid.NewString(), Name: "my first thing", Description: "my first description"}
			secondThing    = model.Thing{ID: uuid.NewString(), Name: "my second thing", Description: "my second description"}
			output         service.BoxContentResult
			expectedOutput = service.BoxContentResult{
				Boxes:  []model.Box{box},
				Things: []model.Thing{firstThing, secondThing},
			}
		)

		BeforeEach(func() {
			repo := storage.NewRepository()
			repo.SaveBox(box)
			repo.SaveThing(firstThing, box.ID)
			repo.SaveThing(secondThing, box.ID)
			output, err = repo.GetBoxContent(box.ID)
		})

		It("should not return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the expected result", func() {
			Expect(output.Equals(expectedOutput)).To(BeTrue())
		})
	})
})
