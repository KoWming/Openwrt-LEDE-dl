package dnsstamps

import (
	"encoding/hex"
	"strings"
	"testing"
)

var (
	pk1 []byte
)

func init() {
	var err error
	// generated with:
	// openssl x509 -noout -fingerprint -sha256 -inform pem -in /etc/ssl/certs/Go_Daddy_Class_2_CA.pem
	pkStr := "C3:84:6B:F2:4B:9E:93:CA:64:27:4C:0E:C6:7C:1E:CC:5E:02:4F:FC:AC:D2:D7:40:19:35:0E:81:FE:54:6A:E4"
	pk1, err = hex.DecodeString(strings.Replace(pkStr, ":", "", -1))
	if err != nil {
		panic(err)
	}
	if len(pk1) != 32 {
		panic("invalid public key fingerprint")
	}
}

func TestDNSCryptStampCreate(t *testing.T) {
	// same as exampleStamp in dnscrypt-stamper
	const expected = `sdns://AQcAAAAAAAAACTEyNy4wLjAuMSDDhGvyS56TymQnTA7GfB7MXgJP_KzS10AZNQ6B_lRq5BkyLmRuc2NyeXB0LWNlcnQubG9jYWxob3N0`

	var stamp ServerStamp
	stamp.Props |= ServerInformalPropertyDNSSEC
	stamp.Props |= ServerInformalPropertyNoLog
	stamp.Props |= ServerInformalPropertyNoFilter
	stamp.Proto = StampProtoTypeDNSCrypt
	stamp.ServerAddrStr = "127.0.0.1"

	stamp.ProviderName = "2.dnscrypt-cert.localhost"
	stamp.ServerPk = pk1
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDNSCryptStampParse(t *testing.T) {
	// AdGuard DNSCrypt
	stampStr := "sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20"
	stamp, err := NewServerStampFromString(stampStr)

	if err != nil || stamp.ProviderName == "" {
		t.Fatalf("Could not parse stamp %s: %s", stampStr, err)
	}

	if stamp.Proto != StampProtoTypeDNSCrypt ||
		stamp.ProviderName != "2.dnscrypt.default.ns1.adguard.com" ||
		stamp.ServerAddrStr != "176.103.130.130:5443" {
		t.Fatalf("Wrong stamp parse results: %v", stamp)
	}
}

func TestDOHStamp(t *testing.T) {
	const expected = `sdns://AgcAAAAAAAAACTEyNy4wLjAuMSDDhGvyS56TymQnTA7GfB7MXgJP_KzS10AZNQ6B_lRq5AtleGFtcGxlLmNvbQovZG5zLXF1ZXJ5`

	var stamp ServerStamp
	stamp.Props |= ServerInformalPropertyDNSSEC
	stamp.Props |= ServerInformalPropertyNoLog
	stamp.Props |= ServerInformalPropertyNoFilter
	stamp.ServerAddrStr = "127.0.0.1"

	stamp.Proto = StampProtoTypeDoH
	stamp.ProviderName = "example.com"
	stamp.Hashes = [][]uint8{pk1}
	stamp.Path = "/dns-query"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOHShortStamp(t *testing.T) {
	const expected = `sdns://AgAAAAAAAAAAAAALZXhhbXBsZS5jb20KL2Rucy1xdWVyeQ`

	var stamp ServerStamp
	stamp.Proto = StampProtoTypeDoH
	stamp.ProviderName = "example.com"
	stamp.Path = "/dns-query"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOHStampParse(t *testing.T) {
	// Google DoH
	stampStr := "sdns://AgUAAAAAAAAAACAe9iTP_15r07rd8_3b_epWVGfjdymdx-5mdRZvMAzBuQ5kbnMuZ29vZ2xlLmNvbQ0vZXhwZXJpbWVudGFs"
	stamp, err := NewServerStampFromString(stampStr)

	if err != nil || stamp.ProviderName == "" {
		t.Fatalf("Could not parse stamp %s: %s", stampStr, err)
	}

	if stamp.Proto != StampProtoTypeDoH ||
		stamp.ProviderName != "dns.google.com" ||
		stamp.Path != "/experimental" {
		t.Fatalf("Wrong stamp parse results: %v", stamp)
	}
}

func TestDOTStamp(t *testing.T) {
	const expected = `sdns://AwcAAAAAAAAACTEyNy4wLjAuMSDDhGvyS56TymQnTA7GfB7MXgJP_KzS10AZNQ6B_lRq5AtleGFtcGxlLmNvbQ`

	var stamp ServerStamp
	stamp.Props |= ServerInformalPropertyDNSSEC
	stamp.Props |= ServerInformalPropertyNoLog
	stamp.Props |= ServerInformalPropertyNoFilter
	stamp.ServerAddrStr = "127.0.0.1"

	stamp.Proto = StampProtoTypeTLS
	stamp.ProviderName = "example.com"
	stamp.Hashes = [][]uint8{pk1}
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOTShortStamp(t *testing.T) {
	const expected = "sdns://AwAAAAAAAAAAAAAPZG5zLmFkZ3VhcmQuY29t"

	var stamp ServerStamp
	stamp.Proto = StampProtoTypeTLS
	stamp.ProviderName = "dns.adguard.com"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOQStamp(t *testing.T) {
	const expected = `sdns://BAcAAAAAAAAACTEyNy4wLjAuMSDDhGvyS56TymQnTA7GfB7MXgJP_KzS10AZNQ6B_lRq5AtleGFtcGxlLmNvbQ`

	var stamp ServerStamp
	stamp.Props |= ServerInformalPropertyDNSSEC
	stamp.Props |= ServerInformalPropertyNoLog
	stamp.Props |= ServerInformalPropertyNoFilter
	stamp.ServerAddrStr = "127.0.0.1"

	stamp.Proto = StampProtoTypeDoQ
	stamp.ProviderName = "example.com"
	stamp.Hashes = [][]uint8{pk1}
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOQShortStamp(t *testing.T) {
	const expected = "sdns://BAAAAAAAAAAAAAAPZG5zLmFkZ3VhcmQuY29t"

	var stamp ServerStamp
	stamp.Proto = StampProtoTypeDoQ
	stamp.ProviderName = "dns.adguard.com"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestDOQOnlyPort(t *testing.T) {
	const expected = "sdns://BAAAAAAAAAAABTo3ODQ0AA9kbnMuYWRndWFyZC5jb20"

	var stamp ServerStamp
	stamp.ServerAddrStr = ":7844"
	stamp.Proto = StampProtoTypeDoQ
	stamp.ProviderName = "dns.adguard.com"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}

func TestPlainStamp(t *testing.T) {
	const expected = `sdns://AAcAAAAAAAAABzguOC44Ljg`

	var stamp ServerStamp
	stamp.Props |= ServerInformalPropertyDNSSEC
	stamp.Props |= ServerInformalPropertyNoLog
	stamp.Props |= ServerInformalPropertyNoFilter
	stamp.ServerAddrStr = "127.0.0.1"

	stamp.Proto = StampProtoTypePlain
	stamp.ServerAddrStr = "8.8.8.8"
	stampStr := stamp.String()

	if stampStr != expected {
		t.Errorf("expected stamp %q but got instead %q", expected, stampStr)
	}

	parsedStamp, err := NewServerStampFromString(stampStr)
	if err != nil {
		t.Fatal(err)
	}
	ps := parsedStamp.String()
	if ps != stampStr {
		t.Errorf("re-parsed stamp string is %q, but %q expected", ps, stampStr)
	}
}
