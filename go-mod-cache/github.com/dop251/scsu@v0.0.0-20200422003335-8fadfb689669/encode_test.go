package scsu

import (
	"bytes"
	"errors"
	"testing"
)

const (
	referenceString = "　♪リンゴ可愛いや可愛いやリンゴ。半世紀も前に流行した「リンゴの歌」がぴったりするかもしれない。米アップルコンピュータ社のパソコン「マック（マッキントッシュ）」を、こよなく愛する人たちのことだ。「アップル信者」なんて言い方まである。"
)

func TestWriteString(t *testing.T) {
	var b bytes.Buffer
	e := NewWriter(&b)
	n, err := e.WriteString("Москва")
	if err != nil {
		t.Fatal(err)
	}
	if n != 7 {
		t.Fatalf("Unexpected len: %d", n)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x12, 0x9C, 0xBE, 0xC1, 0xBA, 0xB2, 0xB0}) {
		t.Fatalf("Content does not match: %v", b.Bytes())
	}
}

func TestWriteRune(t *testing.T) {
	var b bytes.Buffer
	e := NewWriter(&b)
	n, err := e.WriteRune('\u041C')
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("Unexpected len: %d", n)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x12, 0x9C}) {
		t.Fatalf("Content does not match: %v", b.Bytes())
	}
}

func TestEncodeRuneSlice(t *testing.T) {
	var b bytes.Buffer
	e := NewWriter(&b)
	n, err := e.WriteRunes(RuneSlice([]rune("Москва")))
	if err != nil {
		t.Fatal(err)
	}
	if n != 7 {
		t.Fatalf("Unexpected len: %d", n)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x12, 0x9C, 0xBE, 0xC1, 0xBA, 0xB2, 0xB0}) {
		t.Fatalf("Content does not match: %v", b.Bytes())
	}
}

func TestEncodeAppend(t *testing.T) {
	buf := make([]byte, 0, 8)
	buf = append(buf, "head"...)
	newBuf, err := Encode("body", buf)
	if err != nil {
		t.Fatal(err)
	}
	if string(newBuf) != "headbody" {
		t.Fatalf("Unexpected buf: %v", newBuf)
	}
	buf = buf[:cap(buf)]
	if cap(newBuf) != cap(buf) || &newBuf[cap(newBuf)-1] != &buf[cap(buf)-1] {
		t.Fatal("buffer was reallocated")
	}
}
func TestEncodeNilDst(t *testing.T) {
	buf, err := Encode("test", nil)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != "test" {
		t.Fatalf("Unexpected buf: %v", buf)
	}
}

func TestReferenceString(t *testing.T) {
	var b bytes.Buffer
	e := NewWriter(&b)
	n, err := e.WriteString(referenceString)
	if err != nil {
		t.Fatal(err)
	}
	if n != 178 {
		t.Fatalf("Unexpected len: %d", n)
	}
	if !bytes.Equal(b.Bytes(), []byte{
		0x08, 0x00, 0x1B, 0x4C, 0xEA, 0x16, 0xCA, 0xD3, 0x94, 0x0F, 0x53, 0xEF, 0x61, 0x1B, 0xE5, 0x84,
		0xC4, 0x0F, 0x53, 0xEF, 0x61, 0x1B, 0xE5, 0x84, 0xC4, 0x16, 0xCA, 0xD3, 0x94, 0x08, 0x02, 0x0F,
		0x53, 0x4A, 0x4E, 0x16, 0x7D, 0x00, 0x30, 0x82, 0x52, 0x4D, 0x30, 0x6B, 0x6D, 0x41, 0x88, 0x4C,
		0xE5, 0x97, 0x9F, 0x08, 0x0C, 0x16, 0xCA, 0xD3, 0x94, 0x15, 0xAE, 0x0E, 0x6B, 0x4C, 0x08, 0x0D,
		0x8C, 0xB4, 0xA3, 0x9F, 0xCA, 0x99, 0xCB, 0x8B, 0xC2, 0x97, 0xCC, 0xAA, 0x84, 0x08, 0x02, 0x0E,
		0x7C, 0x73, 0xE2, 0x16, 0xA3, 0xB7, 0xCB, 0x93, 0xD3, 0xB4, 0xC5, 0xDC, 0x9F, 0x0E, 0x79, 0x3E,
		0x06, 0xAE, 0xB1, 0x9D, 0x93, 0xD3, 0x08, 0x0C, 0xBE, 0xA3, 0x8F, 0x08, 0x88, 0xBE, 0xA3, 0x8D,
		0xD3, 0xA8, 0xA3, 0x97, 0xC5, 0x17, 0x89, 0x08, 0x0D, 0x15, 0xD2, 0x08, 0x01, 0x93, 0xC8, 0xAA,
		0x8F, 0x0E, 0x61, 0x1B, 0x99, 0xCB, 0x0E, 0x4E, 0xBA, 0x9F, 0xA1, 0xAE, 0x93, 0xA8, 0xA0, 0x08,
		0x02, 0x08, 0x0C, 0xE2, 0x16, 0xA3, 0xB7, 0xCB, 0x0F, 0x4F, 0xE1, 0x80, 0x05, 0xEC, 0x60, 0x8D,
		0xEA, 0x06, 0xD3, 0xE6, 0x0F, 0x8A, 0x00, 0x30, 0x44, 0x65, 0xB9, 0xE4, 0xFE, 0xE7, 0xC2, 0x06,
		0xCB, 0x82,
	}) {
		t.Fatalf("Content does not match: %v", b.Bytes())
	}

}

func verifyEncode(t *testing.T, s string, invalid bool) {
	b, err := EncodeStrict(s, nil)
	if err != nil {
		if invalid && errors.Is(err, ErrInvalidUTF8) {
			return
		}
		t.Fatal(err)
	} else {
		if invalid {
			t.Fatal("Expected an error")
		}
	}
	s1, err := Decode(b)
	if err != nil {
		t.Fatal(err)
	}
	if s1 != s {
		t.Fatalf("Strings dont match: Expected: '%s', actual: '%s'", s, s1)
	}
	if len(b)-len(s) > 1 {
		t.Fatalf("Size increased too much (was %d, compressed %d)", len(s), len(b))
	}
}

func TestEncodeDecode(t *testing.T) {
	for _, s := range []string{
		"🤷🏻‍♀😰😀",
		"𬀀𛀿\u007f",
		"翻😰😰",
		"😰",
		"00翻0",
		"😰😰Ж😰",
		"Тест可testТест",
		"المؤتمر الدولي العاشر ليونيكود (Unicode Conference)، الذي سيعقد في 10-",
		"סעיף א. כל בני אדם נולדו בני חורין ושווים בערכם ובזכויותיהם.",
		"山自作久筋出難具固馬記式点連類無書着",
		"\U0003f02c𬀀\U0002f03f𭀀\U0002f080\U0001403f𮀿",
		"翫�000",
	} {
		s := s
		t.Run("", func(t *testing.T) {
			t.Parallel()
			verifyEncode(t, s, false)
		})
	}

	for _, s := range []string{
		"�Ϳ͔\u0379Ϳ\xcd0\x8c",
		"翻翻\x025翫翿\x025翫\xe7",
	} {
		s := s
		t.Run("", func(t *testing.T) {
			t.Parallel()
			verifyEncode(t, s, true)
		})
	}
}

func TestEncoderReuse(t *testing.T) {
	var e Encoder
	s := StringRuneSource("á山тест")
	buf, err := e.Encode(s, nil)
	if err != nil {
		t.Fatal(err)
	}
	buf1, err := e.Encode(s, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf, buf1) {
		t.Fatalf("buffers are not equal: %v, %v", buf, buf1)
	}

	buf = buf[:cap(buf)]
	buf1 = buf1[:cap(buf1)]
	if &buf[cap(buf)-1] == &buf1[cap(buf1)-1] {
		t.Fatal("Buffers are the same")
	}
}

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf, _ = Encode(referenceString, buf)
		buf = buf[:0]
	}
}

func BenchmarkEncodeZeroAlloc(b *testing.B) {
	var e Encoder
	var buf []byte
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf, _ = e.Encode(StringRuneSource(referenceString), buf)
		buf = buf[:0]
	}
}
