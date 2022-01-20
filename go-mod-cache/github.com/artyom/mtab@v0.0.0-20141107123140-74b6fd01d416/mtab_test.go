package mtab

import (
	"reflect"
	"testing"
)

func TestEntries(t *testing.T) {
	entries, err := Entries("mounts.testfile")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(entries, reference) {
		t.Log("entries read:")
		for _, e := range entries {
			t.Logf("%#v", e)
		}
		t.Log("entries expected:")
		for _, e := range reference {
			t.Logf("%#v", e)
		}
		t.Fatal("read entries do not match reference")
	}
}

var reference = []Entry{
	Entry{"rootfs", "/", "rootfs", "rw", 0, 0},
	Entry{"sysfs", "/sys", "sysfs", "rw,nosuid,nodev,noexec", 0, 0},
	Entry{"proc", "/proc", "proc", "rw,nosuid,nodev,noexec", 0, 0},
	Entry{"/dev/sdb", "/path with spaces", "xfs", "rw", 0, 0},
}
