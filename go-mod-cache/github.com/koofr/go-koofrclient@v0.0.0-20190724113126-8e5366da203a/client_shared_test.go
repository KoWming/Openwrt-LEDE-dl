package koofrclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientShared", func() {
	It("should list shared", func() {
		shared, err := client.Shared()
		Expect(err).NotTo(HaveOccurred())
		Expect(shared).NotTo(BeEmpty())
	})
})
