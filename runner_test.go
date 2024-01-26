package concourse_test

import (
	"errors"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/suhlig/concourse-resource-go"
)

// These tests ensure that validation is being _invoked_ on request and response.
// The actual validations are tested more exhaustively in validation_test.go.
var _ = Describe("Check Validation", func() {
	var (
		err            error
		resource       concourse.Resource[Source, Version, Params]
		stdin          io.Reader
		stdout, stderr io.Writer
	)

	BeforeEach(func() {
		stdout = &strings.Builder{}
		stderr = &strings.Builder{}
	})

	JustBeforeEach(func(ctx SpecContext) {
		err = concourse.CheckWithValidation[Source, Version, Params](ctx, resource, stdin, stdout, stderr)
	})

	Context("valid request", func() {
		BeforeEach(func() {
			resource = NullResource[Source, Version, Params]{}

			stdin = strings.NewReader(`{
				"source": {
					"url": "https://example.com",
					"version": {
						"number": 2
					}
				}
			}`)
		})

		It("works", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("invalid response", func() {
			BeforeEach(func() {
				resource = Troublemaker[Source, Version, Params]{}
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})

			It("provides a reasonable message", func() {
				Expect(err.Error()).To(ContainSubstring("Field validation for 'Number' failed on the 'max' tag"))
			})
		})
	})

	Context("invalid request", func() {
		BeforeEach(func() {
			stdin = strings.NewReader(`{"source": { "url": "null" }}`)
		})

		It("fails", func() {
			Expect(err).To(HaveOccurred())
		})

		Context("validation error", func() {
			var validationErrors validator.ValidationErrors

			JustBeforeEach(func() {
				validationErrors = errors.Unwrap(err).(validator.ValidationErrors)
			})

			It("has the expected error", func() {
				Expect(validationErrors[0].Field()).To(Equal("URL"))
				Expect(validationErrors[0].Tag()).To(Equal("http_url"))
			})
		})
	})
})
