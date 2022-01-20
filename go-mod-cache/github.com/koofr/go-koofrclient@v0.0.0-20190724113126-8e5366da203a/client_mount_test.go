package koofrclient_test

import (
	k "github.com/koofr/go-koofrclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientMount", func() {
	var mount k.Mount

	It("should list mounts", func() {
		mounts, err := client.Mounts()
		Expect(err).NotTo(HaveOccurred())
		Expect(mounts).NotTo(BeEmpty())
		mount = mounts[0]
	})

	It("get mount details", func() {
		mount, err := client.MountsDetails(mount.Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(mount).NotTo(BeNil())
	})
})
