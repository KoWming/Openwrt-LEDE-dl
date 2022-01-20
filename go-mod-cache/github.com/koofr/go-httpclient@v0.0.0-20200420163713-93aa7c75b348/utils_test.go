package httpclient_test

import (
	. "github.com/koofr/go-httpclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EscapePath", func() {
	It("should escape path", func() {
		Expect(EscapePath("foo+bar baz?&")).To(Equal("foo%2bbar%20baz%3F&"))
	})
})
