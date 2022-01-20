package koofrclient_test

import (
	"fmt"
	k "github.com/koofr/go-koofrclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("ClientDevice", func() {
	testDeviceName := fmt.Sprintf("Test %s", time.Now())
	var device k.Device

	It("should list devices", func() {
		devices, err := client.Devices()
		Expect(err).NotTo(HaveOccurred())
		Expect(devices).NotTo(BeEmpty())
	})

	It("should create new device", func() {
		var err error
		device, err = client.DevicesCreate(testDeviceName, k.StorageHubProvider)
		Expect(err).NotTo(HaveOccurred())
		Expect(device.Name).NotTo(BeEmpty())
	})

	It("should get device details", func() {
		Expect(device).NotTo(BeNil())
		deviceDetails, err := client.DevicesDetails(device.Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(deviceDetails).NotTo(BeNil())
		Expect(deviceDetails).To(Equal(device))
	})

	It("should update device", func() {
		newName := "NewName"
		Expect(device).NotTo(BeNil())
		err := client.DevicesUpdate(device.Id, k.DeviceUpdate{Name: newName})
		Expect(err).NotTo(HaveOccurred())
		deviceDetails, err := client.DevicesDetails(device.Id)
		Expect(err).NotTo(HaveOccurred())
		Expect(deviceDetails.Name).To(Equal(newName))

	})

	It("should delete device", func() {
		Expect(device).NotTo(BeNil())
		err := client.DevicesDelete(device.Id)
		Expect(err).NotTo(HaveOccurred())
	})
})
