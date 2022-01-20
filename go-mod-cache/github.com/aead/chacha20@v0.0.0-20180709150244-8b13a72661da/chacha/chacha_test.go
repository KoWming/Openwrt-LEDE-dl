// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package chacha

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"testing"
)

func toHex(bits []byte) string {
	return hex.EncodeToString(bits)
}

func fromHex(bits string) []byte {
	b, err := hex.DecodeString(bits)
	if err != nil {
		panic(err)
	}
	return b
}

func TestHChaCha20(t *testing.T) {
	defer func(sse2, ssse3, avx, avx2 bool) {
		useSSE2, useSSSE3, useAVX, useAVX2 = sse2, ssse3, avx, avx2
	}(useSSE2, useSSSE3, useAVX, useAVX2)

	if useAVX2 {
		t.Log("AVX2 version")
		testHChaCha20(t)
		useAVX2 = false
	}
	if useAVX {
		t.Log("AVX version")
		testHChaCha20(t)
		useAVX = false
	}
	if useSSSE3 {
		t.Log("SSSE3 version")
		testHChaCha20(t)
		useSSSE3 = false
	}
	if useSSE2 {
		t.Log("SSE2 version")
		testHChaCha20(t)
		useSSE2 = false
	}
	t.Log("generic version")
	testHChaCha20(t)
}

func TestVectors(t *testing.T) {
	defer func(sse2, ssse3, avx, avx2 bool) {
		useSSE2, useSSSE3, useAVX, useAVX2 = sse2, ssse3, avx, avx2
	}(useSSE2, useSSSE3, useAVX, useAVX2)

	if useAVX2 {
		t.Log("AVX2 version")
		testVectors(t)
		useAVX2 = false
	}
	if useAVX {
		t.Log("AVX version")
		testVectors(t)
		useAVX = false
	}
	if useSSSE3 {
		t.Log("SSSE3 version")
		testVectors(t)
		useSSSE3 = false
	}
	if useSSE2 {
		t.Log("SSE2 version")
		testVectors(t)
		useSSE2 = false
	}
	t.Log("generic version")
	testVectors(t)
}

var overflowTests = []struct {
	NonceSize     int
	Counter       uint64
	PlaintextSize int
}{
	{NonceSize: NonceSize, Counter: ^uint64(0), PlaintextSize: 65},
	{NonceSize: NonceSize, Counter: ^uint64(1), PlaintextSize: 129},
	{NonceSize: INonceSize, Counter: uint64(^uint32(0)), PlaintextSize: 65},
	{NonceSize: INonceSize, Counter: uint64(^uint32(1)), PlaintextSize: 129},
	{NonceSize: XNonceSize, Counter: ^uint64(0), PlaintextSize: 65},
	{NonceSize: XNonceSize, Counter: ^uint64(1), PlaintextSize: 129},
}

func TestOverflow(t *testing.T) {
	var key [32]byte
	for i, test := range overflowTests {
		stream, err := NewCipher(make([]byte, test.NonceSize), key[:], 20)
		if err != nil {
			t.Errorf("Test %d: Failed to create cipher.Stream: %v", i, err)
			continue
		}
		stream.SetCounter(test.Counter)
		testOverflow(i, make([]byte, test.PlaintextSize), stream, t)
	}
}

func testOverflow(i int, plaintext []byte, stream cipher.Stream, t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Test %d: expected test to panic but it succeeded", i)
		}
	}()
	stream.XORKeyStream(plaintext, plaintext)
}

func TestIncremental(t *testing.T) {
	defer func(sse2, ssse3, avx, avx2 bool) {
		useSSE2, useSSSE3, useAVX, useAVX2 = sse2, ssse3, avx, avx2
	}(useSSE2, useSSSE3, useAVX, useAVX2)

	if useAVX2 {
		t.Log("AVX2 version")
		testIncremental(t, 5, 2049)
		useAVX2 = false
	}
	if useAVX {
		t.Log("AVX version")
		testIncremental(t, 5, 2049)
		useAVX = false
	}
	if useSSSE3 {
		t.Log("SSSE3 version")
		testIncremental(t, 5, 2049)
		useSSSE3 = false
	}
	if useSSE2 {
		t.Log("SSE2 version")
		testIncremental(t, 5, 2049)
	}
}

func testHChaCha20(t *testing.T) {
	for i, v := range hChaCha20Vectors {
		var key [32]byte
		var nonce [16]byte
		copy(key[:], v.key)
		copy(nonce[:], v.nonce)

		hChaCha20(&key, &nonce, &key)
		if !bytes.Equal(key[:], v.keystream) {
			t.Errorf("Test %d: keystream mismatch:\n \t got:  %s\n \t want: %s", i, toHex(key[:]), toHex(v.keystream))
		}
	}
}

func testVectors(t *testing.T) {
	for i, v := range vectors {
		if len(v.plaintext) == 0 {
			v.plaintext = make([]byte, len(v.ciphertext))
		}

		dst := make([]byte, len(v.ciphertext))

		XORKeyStream(dst, v.plaintext, v.nonce, v.key, v.rounds)
		if !bytes.Equal(dst, v.ciphertext) {
			t.Errorf("Test %d: ciphertext mismatch:\n \t got:  %s\n \t want: %s", i, toHex(dst), toHex(v.ciphertext))
		}

		c, err := NewCipher(v.nonce, v.key, v.rounds)
		if err != nil {
			t.Fatal(err)
		}
		c.XORKeyStream(dst[:1], v.plaintext[:1])
		c.XORKeyStream(dst[1:], v.plaintext[1:])
		if !bytes.Equal(dst, v.ciphertext) {
			t.Errorf("Test %d: ciphertext mismatch:\n \t got:  %s\n \t want: %s", i, toHex(dst), toHex(v.ciphertext))
		}
	}
}

func testIncremental(t *testing.T, iter int, size int) {
	sse2, ssse3, avx, avx2 := useSSE2, useSSSE3, useAVX, useAVX2
	msg, ref, stream := make([]byte, size), make([]byte, size), make([]byte, size)

	for i := 0; i < iter; i++ {
		var key [32]byte
		var nonce []byte
		switch i % 3 {
		case 0:
			nonce = make([]byte, 8)
		case 1:
			nonce = make([]byte, 12)
		case 2:
			nonce = make([]byte, 24)
		}

		for j := range key {
			key[j] = byte(len(nonce) + i)
		}
		for j := range nonce {
			nonce[j] = byte(i)
		}

		for j := 0; j <= len(msg); j++ {
			useSSE2, useSSSE3, useAVX, useAVX2 = false, false, false, false
			XORKeyStream(ref[:j], msg[:j], nonce, key[:], 20)

			useSSE2, useSSSE3, useAVX, useAVX2 = sse2, ssse3, avx, avx2
			XORKeyStream(stream[:j], msg[:j], nonce, key[:], 20)

			if !bytes.Equal(ref[:j], stream[:j]) {
				t.Fatalf("Iteration %d failed:\n Message length: %d\n\n got:  %s\nwant: %s", i, j, toHex(stream[:j]), toHex(ref[:j]))
			}

			useSSE2, useSSSE3, useAVX, useAVX2 = false, false, false, false
			c, _ := NewCipher(nonce, key[:], 20)
			c.XORKeyStream(stream[:j], msg[:j])

			useSSE2, useSSSE3, useAVX, useAVX2 = sse2, ssse3, avx, avx2
			c, _ = NewCipher(nonce, key[:], 20)
			c.XORKeyStream(stream[:j], msg[:j])

			if !bytes.Equal(ref[:j], stream[:j]) {
				t.Fatalf("Iteration %d failed:\n Message length: %d\n\n got:  %s\nwant: %s", i, j, toHex(stream[:j]), toHex(ref[:j]))
			}
		}
		copy(msg, stream)
	}
}

var hChaCha20Vectors = []struct {
	key, nonce, keystream []byte
}{
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000000000000000000000000000"),
		fromHex("1140704c328d1d5d0e30086cdf209dbd6a43b8f41518a11cc387b669b2ee6586"),
	},
	{
		fromHex("8000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000000000000000000000000000"),
		fromHex("7d266a7fd808cae4c02a0a70dcbfbcc250dae65ce3eae7fc210f54cc8f77df86"),
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000001"),
		fromHex("000000000000000000000000000000000000000000000002"),
		fromHex("e0c77ff931bb9163a5460c02ac281c2b53d792b1c43fea817e9ad275ae546963"),
	},
	{
		fromHex("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
		fromHex("000102030405060708090a0b0c0d0e0f1011121314151617"),
		fromHex("51e3ff45a895675c4b33b46c64f4a9ace110d34df6a2ceab486372bacbd3eff6"),
	},
}

var vectors = []struct {
	key, nonce, plaintext, ciphertext []byte
	rounds                            int
}{
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("76b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770dc7da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee6586"),
		20,
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000001"),
		fromHex("000000000000000000000002"),
		fromHex("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"416e79207375626d697373696f6e20746f20746865204945544620696e74656e6465642062792074686520436f6e7472696275746f7220666f72207075626c69" +
			"636174696f6e20617320616c6c206f722070617274206f6620616e204945544620496e7465726e65742d4472616674206f722052464320616e6420616e792073" +
			"746174656d656e74206d6164652077697468696e2074686520636f6e74657874206f6620616e204945544620616374697669747920697320636f6e7369646572" +
			"656420616e20224945544620436f6e747269627574696f6e222e20537563682073746174656d656e747320696e636c756465206f72616c2073746174656d656e" +
			"747320696e20494554462073657373696f6e732c2061732077656c6c206173207772697474656e20616e6420656c656374726f6e696320636f6d6d756e696361" +
			"74696f6e73206d61646520617420616e792074696d65206f7220706c6163652c207768696368206172652061646472657373656420746f"),
		fromHex("ecfa254f845f647473d3cb140da9e87606cb33066c447b87bc2666dde3fbb739a371c9ec7abcb4cfa9211f7d90f64c2d07f89e5cf9b93e330a6e4c08af5ba6d5" +
			"a3fbf07df3fa2fde4f376ca23e82737041605d9f4f4f57bd8cff2c1d4b7955ec2a97948bd3722915c8f3d337f7d370050e9e96d647b7c39f56e031ca5eb6250d" +
			"4042e02785ececfa4b4bb5e8ead0440e20b6e8db09d881a7c6132f420e52795042bdfa7773d8a9051447b3291ce1411c680465552aa6c405b7764d5e87bea85a" +
			"d00f8449ed8f72d0d662ab052691ca66424bc86d2df80ea41f43abf937d3259dc4b2d0dfb48a6c9139ddd7f76966e928e635553ba76c5c879d7b35d49eb2e62b" +
			"0871cdac638939e25e8a1e0ef9d5280fa8ca328b351c3c765989cbcf3daa8b6ccc3aaf9f3979c92b3720fc88dc95ed84a1be059c6499b9fda236e7e818b04b0b" +
			"c39c1e876b193bfe5569753f88128cc08aaa9b63d1a16f80ef2554d7189c411f5869ca52c5b83fa36ff216b9c1d30062bebcfd2dc5bce0911934fda79a86f6e6" +
			"98ced759c3ff9b6477338f3da4f9cd8514ea9982ccafb341b2384dd902f3d1ab7ac61dd29c6f21ba5b862f3730e37cfdc4fd806c22f221"),
		20,
	},
	{
		fromHex("8000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("e29edae0466dea17f2576ce95025dd2db2d34fc81b5153f1b70a87f315a35286"),
		20,
	},
	{
		fromHex("8000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("e29edae0466dea17f2576ce95025dd2db2d34fc81b5153f1b70a87f315a35286fb56db91e8dbf0a93faaa25777aad63450dae65ce3eae7fc210f54cc8f77df8662f8" +
			"955228b2358d61d8c5ccf63a6c40203be5fb4541c39c52861de70b8a1416ddd3fe9a818bae8f0e8ff2288cede0459fbb00032fd85fef972fcb586c228d"),
		20,
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("0000000000000000"),
		nil,
		fromHex("76b8e0ada0f13d90405d6ae55386bd28bdd219b8a08ded1aa836efcc8b770dc7da41597c5157488d7724e03fb8d84a376a43b8f41518a11cc387b669b2ee65869f07" +
			"e7be5551387a98ba977c732d080dcb0f29a048e3656912c6533e32ee7aed29b721769ce64e43d57133b074d839d531ed1f28510afb45ace10a1f4b794d6f2d09a0e663266ce1ae7ed1081968a0758e7" +
			"18e997bd362c6b0c34634a9a0b35d012737681f7b5d0f281e3afde458bc1e73d2d313c9cf94c05ff3716240a248f21320a058d7b3566bd520daaa3ed2bf0ac5b8b120fb852773c3639734b45c91a42d" +
			"d4cb83f8840d2eedb158131062ac3f1f2cf8ff6dcd1856e86a1e6c3167167ee5a688742b47c5adfb59d4df76fd1db1e51ee03b1ca9f82aca173edb8b7293474ebe980f904d10c916442b4783a0e9848" +
			"60cb6c957b39c38ed8f51cffaa68a4de01025a39c504546b9dc1406a7eb28151e5150d7b204baa719d4f091021217db5cf1b5c84c4fa71a879610a1a695ac527c5b56774a6b8a21aae88685868e094c" +
			"f29ef4090af7a90cc07e8817aa528763797d3c332b67ca4bc110642c2151ec47ee84cb8c42d85f10e2a8cb18c3b7335f26e8c39a12b1bcc1707177b7613873"),
		20,
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("0100000000000000"),
		nil,
		fromHex("ef3fdfd6c61578fbf5cf35bd3dd33b8009631634d21e42ac33960bd138e50d32111e4caf237ee53ca8ad6426194a88545ddc497a0b466e7d6bbdb0041b2f586b5305" +
			"e5e44aff19b235936144675efbe4409eb7e8e5f1430f5f5836aeb49bb5328b017c4b9dc11f8a03863fa803dc71d5726b2b6b31aa32708afe5af1d6b690584d58792b271e5fdb92c486051c48b79a4d4" +
			"8a109bb2d0477956e74c25e93c3c2db34bf779470464a033b8394517a5cf3576a6618c8551a456628b253ef0117c90cd46d8177a2a06d16e20e05c05f889bf87e95d6ee8a03807d1cd53d586872b125" +
			"9d0647da7b7aae80af9b3aad41ad5a8141d2e156c9dd52a3bd2ae165bd7d6a2a4e2cf6938b8b390828ff20dc8fd60e2cd17fe368e35b467a70654ba93cfa62760a9d2f26da7818d4d863808e1add5ff" +
			"db76d41efd524ded4246e03caa008950c91dedfc9a8e68173fe481c4d3d3c215fdf3af22aeab0097b835a84faabbbce094c6181a193ffeda067271ff7c10cce76542241116283842e31e922430211dc" +
			"b38e556158fc2daaec367b705b75f782f8bc2c2c5e33a375390c3052f7e3446feb105fb47820f1d2539811c5b49bb76dc15f2d20a7e2c200b573db9f653ed7"),
		20,
	},
	{
		fromHex("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
		fromHex("0001020304050607"),
		nil,
		fromHex("f798a189f195e66982105ffb640bb7757f579da31602fc93ec01ac56f85ac3c134a4547b733b46413042c9440049176905d3be59ea1c53f15916155c2be8241a3800" +
			"8b9a26bc35941e2444177c8ade6689de95264986d95889fb60e84629c9bd9a5acb1cc118be563eb9b3a4a472f82e09a7e778492b562ef7130e88dfe031c79db9d4f7c7a899151b9a475032b63fc3852" +
			"45fe054e3dd5a97a5f576fe064025d3ce042c566ab2c507b138db853e3d6959660996546cc9c4a6eafdc777c040d70eaf46f76dad3979e5c5360c3317166a1c894c94a371876a94df7628fe4eaaf2cc" +
			"b27d5aaae0ad7ad0f9d4b6ad3b54098746d4524d38407a6deb3ab78fab78c94213668bbbd394c5de93b853178addd6b97f9fa1ec3e56c00c9ddff0a44a204241175a4cab0f961ba53ede9bdf960b94f" +
			"9829b1f3414726429b362c5b538e391520f489b7ed8d20ae3fd49e9e259e44397514d618c96c4846be3c680bdc11c71dcbbe29ccf80d62a0938fa549391e6ea57ecbe2606790ec15d2224ae307c1442" +
			"26b7c4e8c2f97d2a1d67852d29beba110edd445197012062a393a9c92803ad3b4f31d7bc6033ccf7932cfed3f019044d25905916777286f82f9a4cc1ffe430"),
		20,
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("9bf49a6a0755f953811fce125f2683d50429c3bb49e074147e0089a52eae155f0564f879d27ae3c02ce82834acfa8c793a629f2ca0de6919610be82f411326be0bd588" +
			"41203e74fe86fc71338ce0173dc628ebb719bdcbcc151585214cc089b442258dcda14cf111c602b8971b8cc843e91e46ca905151c02744a6b017e69316b20cd67c4bdecc538e8be990c1b6425d68bfd3a" +
			"6fe97693e4846351596cca8abf59fddd0b7f52dcc0c60a448cbf9511610b0a742f1e4d238a7a45cae054ec2"),
		12,
	},
	{
		fromHex("8000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("789cc357f0b6cda5395f08c8538f1226d08eb3e16ebd6b6db6cc9ca77d81d900bb9d21f6ef0b720550d161f1a80fab0468e48c086daad356edce3a3f988d8e"),
		12,
	},
	{
		fromHex("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
		fromHex("0001020304050607"),
		nil,
		fromHex("6898eb04f3d151985e28e882f35daf28d2a1689f79081ffb08cdc48edbbd3dcd683c764f3dd7302293928ca3d4ef4194e6e22f41a72204a14b89115d06ca29fb0b9f6e" +
			"ba3da6793a928afe76cdf62a5d5b0898bb9bb2348612189fdb825e5aa7559c9ec79ff80d05079fad81e9bc2521b2ebcb179cebeade91f20ff3e13192d60de2ee983ec07047e7827594773c28448d89e9b" +
			"96bb0f8665b1a56f85abebd584a446e17d5a6fb847a1dbf341ece5124ff5f80d4a57fb7edf65a2907939b2f3c9654ccbfa2e5225edc8d799bf7ce296d6c8f9234cec0bd7b91b3d2ddc27f93ff8591ddb3" +
			"62b54fab111a7da9d5b4187661ed0e691f7aa5959fb83112427a95bbeb"),
		12,
	},
	{
		fromHex("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
		fromHex("0001020304050607"),
		nil,
		fromHex("40e1aaea1c843baa28b18eb728fec05dce47b0e824bf9a5d3f1bb1aad13b37fbbf0b0e146732c16380efeab70a1b6edff9acedc876b70d98b61f192290537973"),
		8,
	},
	{
		fromHex("0000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000000000000000000000000000"),
		nil,
		fromHex("bcd02a18bf3f01d19292de30a7a8fdaca4b65e50a6002cc72cd6d2f7c91ac3d5728f83e0aad2bfcf9abd2d2db58faedd65015dd83fc09b131e271043019e8e0f789e96" +
			"89e5208d7fd9e1f3c5b5341f48ef18a13e418998addadd97a3693a987f8e82ecd5c1433bfed1af49750c0f1ff29c4174a05b119aa3a9e8333812e0c0fea49e1ee0134a70a9d49c24e0cbd8fc3ba27e97c" +
			"3322ad487f778f8dc6a122fa59cbe33e7"),
		20,
	},
	{
		fromHex("8000000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000000000000000000000000000"),
		nil,
		fromHex("ccfe8a9e93431bd582f07b3eb0f4a7afc22ef39337ddd84f0d3545b318a315a32b3abb96de0fc6acde48b248fe8a80e6fa72bfcdf9d8d2656b991676476f052d937308" +
			"0e30d8c0e217126a3c64402e1d9404ba9d6b8ce4ad5ac9693f3660638c26ea2cd1b4a8d3348c1e179ead353ee72fee558e9994c51a27195e287d00ec2f8cfef8866d1f98714f40cbe4e18cebabf3cd1fd" +
			"3bb65506e5dce1ad09f438bffe2c96d7f2f0827c8c3f2ca59dbaa393785c6b8da7c69c8a4a63ffd113dcc93de8f52dbcfaed5e4cbcc1dc310b1352868fab7b14d930a9f7a7d47bed0eaf5b151f6dac8bd" +
			"45510698bdc205d70b944ea5450888dd3ec753da9708bf06c0714822dda74f285c361abd0cd1071324c253dc421905edca36e8808bffef091e7dbdecebdad98cf70b7cede72e9c3c4108e5b32ffae0f42" +
			"151a8196939d8e3b8384be1"),
		20,
	},
	{
		fromHex("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
		fromHex("000102030405060708090a0b0c0d0e0f1011121314151617"),
		nil,
		fromHex("e53a61cef151e81401067de33adfc02e90ab205361b49b539fda7f0e63b1bc7d68fbee56c9c20c39960e595f3ea76c979804d08cfa728e66cb5f766b840ec61f9ec20f" +
			"7f90d28dae334426cecb52a8e84b4728a5fdd61deb7f1a3fb63dadf5595e06b6e441670964d595ae59cf21536271bae2594774fb19079b933d8fe744f4"),
		20,
	},
	{
		fromHex("FF00000000000000000000000000000000000000000000000000000000000000"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("4fe0956ef81829ff96ef093f03c15dc0eaf4e6905eff9777a5db78348915689ed64204e8fce664cb71ea4016185d15e05be4329e02fcd472707508ef62fd89565ffa632effdb" +
			"bf08394aa437d8ff093e6cea49b61672cf294474927a8150e06cec9fdec0f5cf26f257fe335a8d7dd6d208e6df6f0a83bb1b0b5c574edc2c9a604e4310acb970815a9819c91a5137794d1ee71ede3e5d59f27e76" +
			"84d287d704fe3945de0a9b66be3d86e66980263602aeb600efaef243b1adf4c701dbf8f57427dee71dacd703d25317ffc7a67e7881ad13f0bf096d3b0486eec71fef5e0efb5964d14eb2cea0336e34ed4444cc2b" +
			"bdbd8ef5ba89a0a5e9e35a2e23b38d3f9136f42aefb25c2e7eae0b42c1d1ada5618c5299aedd469ce4f9353ccbae3f89110922b669b8d1b62e72aaf893b83ca264707efbefdcf22ef2333b01f18a849653b52925" +
			"63c37314bf34289b0636a2f8c24bc97fec554a9c31ec2cb4e30ba70fa965a17561e56739be138d86a4777f866ca24ba24f70913230e1b3ea34a9a90eea1b6a3a81b93286bb582a53e78557845a654775a18efb77" +
			"eee098d2680bc4ceb866874f31c7fadd70262cca6039833522de03cb2527dc5cfc7072db48b6011b852d705c7b24ffedf52facf352ab2512c625811db7965edc87d08f7f27e02665c9a6a42968e4c58cd86aa847" +
			"69658153b62f208b2dcfbcb364d63e6671cf60698640"),
		20,
	},
	{
		fromHex("0120000000000000000000000000007000000000000000000000000000000DEF"),
		fromHex("000000000000000000000000"),
		nil,
		fromHex("ba6bce79c4f79c815b7fec53840ff0549ff5496378aa1f6ba481a48a5b9b8dbea8b820eccbc4eca37e1050fc53510a746037d2707f81e9683ec3f495b02ad0f848d7f9bf67bc" +
			"6299be525d1bf3bfd9953caa12cc4e1d5a6969e6fcd5d3c3e3d9f2e735cd7808755ddda7b22a3ae6040e7f8d05d62661a97d84dad694c69637aea3ae0af9f73303ffce3ae6161281d7a3c7e50a5706d766b34ddd" +
			"eab6974fdab10b3f48fb31f26df72e54c616edf1afc019f240c059a7c003677008227f49b021bc23c9c51d6f85ad136a4aa4950d9692f7094d344d88c05868691eb620d39bd8154986c971a8c9552ff0015fd78a" +
			"6bdd33df94b0056786a1e0ceb9cc9a38a31fbba224c1fb82bf6af376f67e94337a730301a6365d49b0dd56328e0269cbdfb5bcbccf1c7c3f4922ec1310aa2ef8136be788a55190453d3d3153b1b960a16f79365a" +
			"0bc7d6d2d5cda9f0993dbb815ee72f83b9d2ed296598fb21d91c29d1acf4ff0a549784a1d6a4f0935ee18efbf41fdc98d81c449544e9701d92648c06e5f416833b90d15fd4c04fc720a5ec6c6fc8b3d85a66826a" +
			"5e6817e21c4c4c0d7151b128236c41397ad4c6549e827c42269659973c153db70ffc33951b19ff21428091cea3836f72f88082508bae1839b59fa9c2556bdf373419d3cf29a8fad4d1787d829ad884f9927228fc" +
			"0b8bb7f1a067e7bdbf06c3885154f76f5be0cde8c7c59442b72b0e3f0341afe644e7eb4c29a467288aebc893e17b446c63da7551b8b59ebdd0cbcd65bc79a969bd3397f83d149840de731df4c09a833d5bd9feda" +
			"e1cd78a09b233b020de86ab71b9fd425adf84e502cef7c62015eade66ca91b0a90306894b53c7c5147e524d7b919ccdd0731e4eef8fe476b6eed38c91b611cd1777b9acf6eee0a11eaff16ae872db92a5d133fe7" +
			"bed999882da283893dd1e96f530be3cd36bf38c16deed2cd77651b6e0d3628de3cb86a78f1d07f6fc79434da5f73888be617b84595acef154f66b95ade1a3e120421a9dac6eec1e5b60139da3d604a03d4a9b7a3" +
			"0810a9c7d551aa8df08e11544486ad33000bfe410e8e6f35cb9d22806a5fcacefc6a1257d373d426243576fad9b20ad5ba84befc1a47c79d7bd2923b5776d3df86c8ed98b700d317502849ec8c02ecb8513a7a32" +
			"e2db15e75a814f12cfc20429ae06cae2021406b4f174ce56dca65f7994a3b2722e764520a52f87d0a887fc771dbfbf381b4f750dc074fedec1a43a4df37a5a2c148f89d9630ebbd1be1858bed10207cdacae9a0a" +
			"b92df58de53de4718f929a83474fbcf9969f1d28a5b257cacd56f0ff0bc425c93d8c91ac833c2cfefb97d82fe6236f3ec3c29e0112a6cac5abfec733db41265f8ff486e7d7fa0b3d9766357377f089056c9408d8" +
			"2f09f18700236cc1058ea1c273e287d07d521fdbb5e28d41cc1d95999eccee"),
		20,
	},
}
