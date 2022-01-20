package koofrclient_test

import (
	"bytes"
	"io/ioutil"
	"strings"

	koofrclient "github.com/koofr/go-koofrclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientFiles", func() {
	BeforeEach(func() {
		client.FilesDelete(defaultMountId, rootPath)
		parts := strings.Split(rootPath, "/")
		created := "/"
		for _, part := range parts {
			if part != "" {
				client.FilesDelete(defaultMountId, created+"/"+part)
				err := client.FilesNewFolder(defaultMountId, created, part)
				Expect(err).NotTo(HaveOccurred())
				created += "/" + part
			}
		}
	})

	It("should get file info", func() {
		info, err := client.FilesInfo(defaultMountId, rootPath)
		Expect(err).NotTo(HaveOccurred())
		parts := strings.Split(rootPath, "/")
		expected := parts[len(parts)-1]
		Expect(info.Name).To(Equal(expected))
	})

	It("should list files", func() {
		files, err := client.FilesList(defaultMountId, rootPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(files).To(HaveLen(0))
	})

	It("should get files tree", func() {
		tree, err := client.FilesTree(defaultMountId, rootPath)
		Expect(err).NotTo(HaveOccurred())
		parts := strings.Split(rootPath, "/")
		expected := parts[len(parts)-1]
		Expect(tree.Name).To(Equal(expected))
		Expect(tree.Children).To(HaveLen(0))
	})

	It("should create new folder and delete it", func() {
		err := client.FilesNewFolder(defaultMountId, rootPath, "dir")
		Expect(err).NotTo(HaveOccurred())
		err = client.FilesDelete(defaultMountId, rootPath+"/dir")
		Expect(err).NotTo(HaveOccurred())
		_, err = client.FilesInfo(defaultMountId, rootPath+"/dir")
		Expect(err).To(HaveOccurred())
	})

	It("should copy a folder", func() {
		err := client.FilesNewFolder(defaultMountId, rootPath, "dir")
		Expect(err).NotTo(HaveOccurred())
		err = client.FilesCopy(defaultMountId, rootPath+"/dir", defaultMountId, rootPath+"/dircopy")
		Expect(err).NotTo(HaveOccurred())
		_, err = client.FilesInfo(defaultMountId, rootPath+"/dir")
		Expect(err).NotTo(HaveOccurred())
		_, err = client.FilesInfo(defaultMountId, rootPath+"/dircopy")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should move a folder", func() {
		err := client.FilesNewFolder(defaultMountId, rootPath, "dir")
		Expect(err).NotTo(HaveOccurred())
		err = client.FilesMove(defaultMountId, rootPath+"/dir", defaultMountId, rootPath+"/dirmoved")
		Expect(err).NotTo(HaveOccurred())
		_, err = client.FilesInfo(defaultMountId, rootPath+"/dir")
		Expect(err).To(HaveOccurred())
		_, err = client.FilesInfo(defaultMountId, rootPath+"/dirmoved")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should put file, get it and delete it", func() {
		newName, err := client.FilesPut(defaultMountId, rootPath, "file.txt", bytes.NewReader([]byte("content")))
		Expect(err).NotTo(HaveOccurred())
		Expect(newName).To(Equal("file.txt"))
		reader, err := client.FilesGet(defaultMountId, rootPath+"/file.txt")
		Expect(err).NotTo(HaveOccurred())
		content, err := ioutil.ReadAll(reader)
		Expect(err).NotTo(HaveOccurred())
		Expect(content).To(Equal([]byte("content")))
		reader, err = client.FilesGetRange(defaultMountId, rootPath+"/file.txt", &koofrclient.FileSpan{Start: 2, End: 3})
		Expect(err).NotTo(HaveOccurred())
		content, err = ioutil.ReadAll(reader)
		Expect(err).NotTo(HaveOccurred())
		Expect(content).To(Equal([]byte("nt")))
		err = client.FilesDelete(defaultMountId, rootPath+"/file.txt")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should put a file and set its modified time", func() {
		mtime := int64(1562663291000)
		putOptions := koofrclient.PutOptions{
			SetModified: &mtime,
		}
		_, err := client.FilesPutWithOptions(defaultMountId, rootPath, "file.txt", bytes.NewReader([]byte("content")), &putOptions)
		Expect(err).NotTo(HaveOccurred())
		reader, err := client.FilesGet(defaultMountId, rootPath+"/file.txt")
		Expect(err).NotTo(HaveOccurred())
		content, err := ioutil.ReadAll(reader)
		Expect(err).NotTo(HaveOccurred())
		Expect(content).To(Equal([]byte("content")))
		reader, err = client.FilesGetRange(defaultMountId, rootPath+"/file.txt", &koofrclient.FileSpan{Start: 2, End: 3})
		Expect(err).NotTo(HaveOccurred())
		content, err = ioutil.ReadAll(reader)
		Expect(err).NotTo(HaveOccurred())
		Expect(content).To(Equal([]byte("nt")))
		info, err := client.FilesInfo(defaultMountId, rootPath+"/file.txt")
		Expect(err).NotTo(HaveOccurred())
		Expect(info.Modified).To(Equal(int64(1562663291000)))
		err = client.FilesDelete(defaultMountId, rootPath+"/file.txt")
		Expect(err).NotTo(HaveOccurred())
	})
})
