package concourse_test

import (
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/suhlig/concourse-resource-go"
)

var _ = Describe("Check Request Validation", func() {
	var (
		err     error
		request concourse.CheckRequest[Source, Version]
	)

	JustBeforeEach(func() {
		err = request.Validate()
	})

	Context("valid request", func() {
		BeforeEach(func() {
			request = concourse.CheckRequest[Source, Version]{
				Source:  Source{URL: "https://example.com"},
				Version: Version{Number: 9},
			}
		})

		It("succeeds", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("invalid source", func() {
		BeforeEach(func() {
			request = concourse.CheckRequest[Source, Version]{
				Source: Source{URL: "this is not a URL"},
			}
		})

		It("fails", func() {
			Expect(err).To(HaveOccurred())
		})

		Context("validation error", func() {
			var validationErrors validator.ValidationErrors

			JustBeforeEach(func() {
				validationErrors = err.(validator.ValidationErrors)
			})

			It("has the expected error", func() {
				Expect(validationErrors[0].Field()).To(Equal("URL"))
				Expect(validationErrors[0].Tag()).To(Equal("http_url"))
			})
		})
	})

	Context("invalid version", func() {
		BeforeEach(func() {
			request = concourse.CheckRequest[Source, Version]{
				Source:  Source{URL: "https://example.com"},
				Version: Version{Number: 12},
			}
		})

		It("fails", func() {
			Expect(err).To(HaveOccurred())
		})

		Context("validation error", func() {
			var validationErrors validator.ValidationErrors

			JustBeforeEach(func() {
				validationErrors = err.(validator.ValidationErrors)
			})

			It("has the expected error", func() {
				Expect(validationErrors[0].Field()).To(Equal("Number"))
				Expect(validationErrors[0].Tag()).To(Equal("max"))
			})
		})
	})
})

var _ = Describe("Response Validation", func() {
	var (
		err      error
		response concourse.Response[Version]
	)

	JustBeforeEach(func() {
		err = response.Validate()
	})

	Describe("Missing version", func() {
		BeforeEach(func() {
			response = concourse.Response[Version]{}
		})

		It("fails", func() {
			Expect(err).To(HaveOccurred())
		})

		Context("validation error", func() {
			var validationErrors validator.ValidationErrors

			JustBeforeEach(func() {
				validationErrors = err.(validator.ValidationErrors)
			})

			It("has the expected number of validation errors", func() {
				Expect(validationErrors).To(HaveLen(1))
			})

			It("has the expected error", func() {
				Expect(validationErrors[0].Field()).To(Equal("Version"))
				Expect(validationErrors[0].Tag()).To(Equal("required"))
			})
		})
	})

	Describe("Invalid metadata", func() {
		BeforeEach(func() {
			response = concourse.Response[Version]{
				Version: Version{Number: 1},
				Metadata: []concourse.NameValuePair{
					{Name: "", Value: ""},
				},
			}
		})

		It("fails", func() {
			Expect(err).To(HaveOccurred())
		})

		Context("validation error", func() {
			var validationErrors validator.ValidationErrors

			JustBeforeEach(func() {
				validationErrors = err.(validator.ValidationErrors)
			})

			It("has the expected number of validation errors", func() {
				Expect(validationErrors).To(HaveLen(1))
			})

			It("has the expected error", func() {
				Expect(validationErrors[0].Field()).To(Equal("Name"))
				Expect(validationErrors[0].Tag()).To(Equal("required"))
			})
		})
	})
})
