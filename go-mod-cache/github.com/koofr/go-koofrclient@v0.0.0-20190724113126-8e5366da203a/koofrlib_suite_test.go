package koofrclient_test

import (
	"os"
	"strings"
	"testing"

	k "github.com/koofr/go-koofrclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	client         *k.KoofrClient
	apiBase        string
	rootPath       string
	email          string
	password       string
	defaultMountId string
)

func TestKoofrclient(t *testing.T) {
	RegisterFailHandler(Fail)

	apiBase = os.Getenv("KOOFR_APIBASE")
	if apiBase == "" {
		t.Fatal("Missing KOOFR_APIBASE")
	}

	rootPath = os.Getenv("KOOFR_ROOTPATH")
	if rootPath == "" {
		t.Fatal("Missing KOOFR_ROOTPATH")
	}
	if !strings.HasPrefix(rootPath, "/") {
		rootPath = "/" + rootPath
	}

	email = os.Getenv("KOOFR_EMAIL")
	if email == "" {
		t.Fatal("Missing KOOFR_EMAIL")
	}

	password = os.Getenv("KOOFR_PASSWORD")
	if password == "" {
		t.Fatal("Missing KOOFR_PASSWORD")
	}

	client = k.NewKoofrClient(apiBase, true)

	err := client.Authenticate(email, password)

	if err != nil {
		t.Fatal("Koofr authorization failed")
	}

	mounts, err := client.Mounts()

	if err != nil {
		t.Fatal("Koofr listing mounts failed")
	}

	if len(mounts) == 0 {
		t.Fatal("Koofr mounts must not be empty")
	}

	for _, m := range mounts {
		if m.IsPrimary {
			defaultMountId = m.Id
		}
	}

	if defaultMountId == "" {
		t.Fatal("Koofr primary mount not found")
	}

	RunSpecs(t, "Koofrclient Suite")
}
