package netutil_test

import "net"

// Common test IPs.  Do not mutate.
var (
	testIPv4 = net.IP{1, 2, 3, 4}
	testIPv6 = net.IP{
		0x12, 0x34, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xcd, 0xef,
	}
)
