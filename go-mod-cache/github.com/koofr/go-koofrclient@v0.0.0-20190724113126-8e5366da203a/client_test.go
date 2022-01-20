package koofrclient_test

import (
	k "github.com/koofr/go-koofrclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	It("should create client with token", func() {
		c := k.NewKoofrClient(apiBase, true)
		Expect(c).NotTo(BeNil())
		c.SetToken("eac6e8d7-edea-4af1-92db-9636795379d1")
		Expect(c).NotTo(BeNil())
	})

	It("should create client and authorize", func() {
		c := k.NewKoofrClient(apiBase, true)
		Expect(c).NotTo(BeNil())
		err := c.Authenticate(email, password)
		Expect(err).NotTo(HaveOccurred())
		Expect(c.GetToken()).To(HaveLen(36))
	})
})
