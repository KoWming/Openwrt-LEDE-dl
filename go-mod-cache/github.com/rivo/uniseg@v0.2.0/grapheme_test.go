package uniseg

import (
	"fmt"
	"testing"
)

// Type example.
func ExampleGraphemes() {
	gr := NewGraphemes("👍🏼!")
	for gr.Next() {
		fmt.Printf("%x ", gr.Runes())
	}
	// Output: [1f44d 1f3fc] [21]
}

// The test cases for the simple test function.
var testCases = []struct {
	original string
	expected [][]rune
}{
	{original: "", expected: [][]rune{}},
	{original: "x", expected: [][]rune{{0x78}}},
	{original: "basic", expected: [][]rune{{0x62}, {0x61}, {0x73}, {0x69}, {0x63}}},
	{original: "möp", expected: [][]rune{{0x6d}, {0x6f, 0x308}, {0x70}}},
	{original: "\r\n", expected: [][]rune{{0xd, 0xa}}},
	{original: "\n\n", expected: [][]rune{{0xa}, {0xa}}},
	{original: "\t*", expected: [][]rune{{0x9}, {0x2a}}},
	{original: "뢴", expected: [][]rune{{0x1105, 0x116c, 0x11ab}}},
	{original: "ܐ܏ܒܓܕ", expected: [][]rune{{0x710}, {0x70f, 0x712}, {0x713}, {0x715}}},
	{original: "ำ", expected: [][]rune{{0xe33}}},
	{original: "ำำ", expected: [][]rune{{0xe33, 0xe33}}},
	{original: "สระอำ", expected: [][]rune{{0xe2a}, {0xe23}, {0xe30}, {0xe2d, 0xe33}}},
	{original: "*뢴*", expected: [][]rune{{0x2a}, {0x1105, 0x116c, 0x11ab}, {0x2a}}},
	{original: "*👩‍❤️‍💋‍👩*", expected: [][]rune{{0x2a}, {0x1f469, 0x200d, 0x2764, 0xfe0f, 0x200d, 0x1f48b, 0x200d, 0x1f469}, {0x2a}}},
	{original: "👩‍❤️‍💋‍👩", expected: [][]rune{{0x1f469, 0x200d, 0x2764, 0xfe0f, 0x200d, 0x1f48b, 0x200d, 0x1f469}}},
	{original: "🏋🏽‍♀️", expected: [][]rune{{0x1f3cb, 0x1f3fd, 0x200d, 0x2640, 0xfe0f}}},
	{original: "🙂", expected: [][]rune{{0x1f642}}},
	{original: "🙂🙂", expected: [][]rune{{0x1f642}, {0x1f642}}},
	{original: "🇩🇪", expected: [][]rune{{0x1f1e9, 0x1f1ea}}},
	{original: "🏳️‍🌈", expected: [][]rune{{0x1f3f3, 0xfe0f, 0x200d, 0x1f308}}},

	// The following tests are taken from
	// http://www.unicode.org/Public/12.0.0/ucd/auxiliary/GraphemeBreakTest.txt,
	// see https://www.unicode.org/license.html for the Unicode license agreement.
	{original: "\u0020\u0020", expected: [][]rune{{0x0020}, {0x0020}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0020\u0308\u0020", expected: [][]rune{{0x0020, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0020\u000D", expected: [][]rune{{0x0020}, {0x000D}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0020\u0308\u000D", expected: [][]rune{{0x0020, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0020\u000A", expected: [][]rune{{0x0020}, {0x000A}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0020\u0308\u000A", expected: [][]rune{{0x0020, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0020\u0001", expected: [][]rune{{0x0020}, {0x0001}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0020\u0308\u0001", expected: [][]rune{{0x0020, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0020\u034F", expected: [][]rune{{0x0020, 0x034F}}},                                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0020\u0308\u034F", expected: [][]rune{{0x0020, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0020\U0001F1E6", expected: [][]rune{{0x0020}, {0x1F1E6}}},                                                                            // ÷ [0.2] SPACE (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0020\u0308\U0001F1E6", expected: [][]rune{{0x0020, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0020\u0600", expected: [][]rune{{0x0020}, {0x0600}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0020\u0308\u0600", expected: [][]rune{{0x0020, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0020\u0903", expected: [][]rune{{0x0020, 0x0903}}},                                                                                   // ÷ [0.2] SPACE (Other) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0020\u0308\u0903", expected: [][]rune{{0x0020, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0020\u1100", expected: [][]rune{{0x0020}, {0x1100}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0020\u0308\u1100", expected: [][]rune{{0x0020, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0020\u1160", expected: [][]rune{{0x0020}, {0x1160}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0020\u0308\u1160", expected: [][]rune{{0x0020, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0020\u11A8", expected: [][]rune{{0x0020}, {0x11A8}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0020\u0308\u11A8", expected: [][]rune{{0x0020, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0020\uAC00", expected: [][]rune{{0x0020}, {0xAC00}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0020\u0308\uAC00", expected: [][]rune{{0x0020, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0020\uAC01", expected: [][]rune{{0x0020}, {0xAC01}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0020\u0308\uAC01", expected: [][]rune{{0x0020, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0020\u231A", expected: [][]rune{{0x0020}, {0x231A}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0020\u0308\u231A", expected: [][]rune{{0x0020, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0020\u0300", expected: [][]rune{{0x0020, 0x0300}}},                                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0020\u0308\u0300", expected: [][]rune{{0x0020, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0020\u200D", expected: [][]rune{{0x0020, 0x200D}}},                                                                                   // ÷ [0.2] SPACE (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0020\u0308\u200D", expected: [][]rune{{0x0020, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0020\u0378", expected: [][]rune{{0x0020}, {0x0378}}},                                                                                 // ÷ [0.2] SPACE (Other) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0020\u0308\u0378", expected: [][]rune{{0x0020, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u000D\u0020", expected: [][]rune{{0x000D}, {0x0020}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] SPACE (Other) ÷ [0.3]
	{original: "\u000D\u0308\u0020", expected: [][]rune{{0x000D}, {0x0308}, {0x0020}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u000D\u000D", expected: [][]rune{{0x000D}, {0x000D}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u000D\u0308\u000D", expected: [][]rune{{0x000D}, {0x0308}, {0x000D}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u000D\u000A", expected: [][]rune{{0x000D, 0x000A}}},                                                                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) × [3.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u000D\u0308\u000A", expected: [][]rune{{0x000D}, {0x0308}, {0x000A}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u000D\u0001", expected: [][]rune{{0x000D}, {0x0001}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u000D\u0308\u0001", expected: [][]rune{{0x000D}, {0x0308}, {0x0001}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u000D\u034F", expected: [][]rune{{0x000D}, {0x034F}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u000D\u0308\u034F", expected: [][]rune{{0x000D}, {0x0308, 0x034F}}},                                                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u000D\U0001F1E6", expected: [][]rune{{0x000D}, {0x1F1E6}}},                                                                            // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u000D\u0308\U0001F1E6", expected: [][]rune{{0x000D}, {0x0308}, {0x1F1E6}}},                                                            // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u000D\u0600", expected: [][]rune{{0x000D}, {0x0600}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u000D\u0308\u0600", expected: [][]rune{{0x000D}, {0x0308}, {0x0600}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u000D\u0903", expected: [][]rune{{0x000D}, {0x0903}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u000D\u0308\u0903", expected: [][]rune{{0x000D}, {0x0308, 0x0903}}},                                                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u000D\u1100", expected: [][]rune{{0x000D}, {0x1100}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u000D\u0308\u1100", expected: [][]rune{{0x000D}, {0x0308}, {0x1100}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u000D\u1160", expected: [][]rune{{0x000D}, {0x1160}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u000D\u0308\u1160", expected: [][]rune{{0x000D}, {0x0308}, {0x1160}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u000D\u11A8", expected: [][]rune{{0x000D}, {0x11A8}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u000D\u0308\u11A8", expected: [][]rune{{0x000D}, {0x0308}, {0x11A8}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u000D\uAC00", expected: [][]rune{{0x000D}, {0xAC00}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u000D\u0308\uAC00", expected: [][]rune{{0x000D}, {0x0308}, {0xAC00}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u000D\uAC01", expected: [][]rune{{0x000D}, {0xAC01}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u000D\u0308\uAC01", expected: [][]rune{{0x000D}, {0x0308}, {0xAC01}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u000D\u231A", expected: [][]rune{{0x000D}, {0x231A}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u000D\u0308\u231A", expected: [][]rune{{0x000D}, {0x0308}, {0x231A}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u000D\u0300", expected: [][]rune{{0x000D}, {0x0300}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u000D\u0308\u0300", expected: [][]rune{{0x000D}, {0x0308, 0x0300}}},                                                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u000D\u200D", expected: [][]rune{{0x000D}, {0x200D}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u000D\u0308\u200D", expected: [][]rune{{0x000D}, {0x0308, 0x200D}}},                                                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u000D\u0378", expected: [][]rune{{0x000D}, {0x0378}}},                                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u000D\u0308\u0378", expected: [][]rune{{0x000D}, {0x0308}, {0x0378}}},                                                                 // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u000A\u0020", expected: [][]rune{{0x000A}, {0x0020}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] SPACE (Other) ÷ [0.3]
	{original: "\u000A\u0308\u0020", expected: [][]rune{{0x000A}, {0x0308}, {0x0020}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u000A\u000D", expected: [][]rune{{0x000A}, {0x000D}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u000A\u0308\u000D", expected: [][]rune{{0x000A}, {0x0308}, {0x000D}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u000A\u000A", expected: [][]rune{{0x000A}, {0x000A}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u000A\u0308\u000A", expected: [][]rune{{0x000A}, {0x0308}, {0x000A}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u000A\u0001", expected: [][]rune{{0x000A}, {0x0001}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u000A\u0308\u0001", expected: [][]rune{{0x000A}, {0x0308}, {0x0001}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u000A\u034F", expected: [][]rune{{0x000A}, {0x034F}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u000A\u0308\u034F", expected: [][]rune{{0x000A}, {0x0308, 0x034F}}},                                                                   // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u000A\U0001F1E6", expected: [][]rune{{0x000A}, {0x1F1E6}}},                                                                            // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u000A\u0308\U0001F1E6", expected: [][]rune{{0x000A}, {0x0308}, {0x1F1E6}}},                                                            // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u000A\u0600", expected: [][]rune{{0x000A}, {0x0600}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u000A\u0308\u0600", expected: [][]rune{{0x000A}, {0x0308}, {0x0600}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u000A\u0903", expected: [][]rune{{0x000A}, {0x0903}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u000A\u0308\u0903", expected: [][]rune{{0x000A}, {0x0308, 0x0903}}},                                                                   // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u000A\u1100", expected: [][]rune{{0x000A}, {0x1100}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u000A\u0308\u1100", expected: [][]rune{{0x000A}, {0x0308}, {0x1100}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u000A\u1160", expected: [][]rune{{0x000A}, {0x1160}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u000A\u0308\u1160", expected: [][]rune{{0x000A}, {0x0308}, {0x1160}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u000A\u11A8", expected: [][]rune{{0x000A}, {0x11A8}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u000A\u0308\u11A8", expected: [][]rune{{0x000A}, {0x0308}, {0x11A8}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u000A\uAC00", expected: [][]rune{{0x000A}, {0xAC00}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u000A\u0308\uAC00", expected: [][]rune{{0x000A}, {0x0308}, {0xAC00}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u000A\uAC01", expected: [][]rune{{0x000A}, {0xAC01}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u000A\u0308\uAC01", expected: [][]rune{{0x000A}, {0x0308}, {0xAC01}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u000A\u231A", expected: [][]rune{{0x000A}, {0x231A}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u000A\u0308\u231A", expected: [][]rune{{0x000A}, {0x0308}, {0x231A}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u000A\u0300", expected: [][]rune{{0x000A}, {0x0300}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u000A\u0308\u0300", expected: [][]rune{{0x000A}, {0x0308, 0x0300}}},                                                                   // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u000A\u200D", expected: [][]rune{{0x000A}, {0x200D}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u000A\u0308\u200D", expected: [][]rune{{0x000A}, {0x0308, 0x200D}}},                                                                   // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u000A\u0378", expected: [][]rune{{0x000A}, {0x0378}}},                                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u000A\u0308\u0378", expected: [][]rune{{0x000A}, {0x0308}, {0x0378}}},                                                                 // ÷ [0.2] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0001\u0020", expected: [][]rune{{0x0001}, {0x0020}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] SPACE (Other) ÷ [0.3]
	{original: "\u0001\u0308\u0020", expected: [][]rune{{0x0001}, {0x0308}, {0x0020}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0001\u000D", expected: [][]rune{{0x0001}, {0x000D}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0001\u0308\u000D", expected: [][]rune{{0x0001}, {0x0308}, {0x000D}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0001\u000A", expected: [][]rune{{0x0001}, {0x000A}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0001\u0308\u000A", expected: [][]rune{{0x0001}, {0x0308}, {0x000A}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0001\u0001", expected: [][]rune{{0x0001}, {0x0001}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0001\u0308\u0001", expected: [][]rune{{0x0001}, {0x0308}, {0x0001}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0001\u034F", expected: [][]rune{{0x0001}, {0x034F}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0001\u0308\u034F", expected: [][]rune{{0x0001}, {0x0308, 0x034F}}},                                                                   // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0001\U0001F1E6", expected: [][]rune{{0x0001}, {0x1F1E6}}},                                                                            // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0001\u0308\U0001F1E6", expected: [][]rune{{0x0001}, {0x0308}, {0x1F1E6}}},                                                            // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0001\u0600", expected: [][]rune{{0x0001}, {0x0600}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0001\u0308\u0600", expected: [][]rune{{0x0001}, {0x0308}, {0x0600}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0001\u0903", expected: [][]rune{{0x0001}, {0x0903}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0001\u0308\u0903", expected: [][]rune{{0x0001}, {0x0308, 0x0903}}},                                                                   // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0001\u1100", expected: [][]rune{{0x0001}, {0x1100}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0001\u0308\u1100", expected: [][]rune{{0x0001}, {0x0308}, {0x1100}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0001\u1160", expected: [][]rune{{0x0001}, {0x1160}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0001\u0308\u1160", expected: [][]rune{{0x0001}, {0x0308}, {0x1160}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0001\u11A8", expected: [][]rune{{0x0001}, {0x11A8}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0001\u0308\u11A8", expected: [][]rune{{0x0001}, {0x0308}, {0x11A8}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0001\uAC00", expected: [][]rune{{0x0001}, {0xAC00}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0001\u0308\uAC00", expected: [][]rune{{0x0001}, {0x0308}, {0xAC00}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0001\uAC01", expected: [][]rune{{0x0001}, {0xAC01}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0001\u0308\uAC01", expected: [][]rune{{0x0001}, {0x0308}, {0xAC01}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0001\u231A", expected: [][]rune{{0x0001}, {0x231A}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0001\u0308\u231A", expected: [][]rune{{0x0001}, {0x0308}, {0x231A}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0001\u0300", expected: [][]rune{{0x0001}, {0x0300}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0001\u0308\u0300", expected: [][]rune{{0x0001}, {0x0308, 0x0300}}},                                                                   // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0001\u200D", expected: [][]rune{{0x0001}, {0x200D}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0001\u0308\u200D", expected: [][]rune{{0x0001}, {0x0308, 0x200D}}},                                                                   // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0001\u0378", expected: [][]rune{{0x0001}, {0x0378}}},                                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0001\u0308\u0378", expected: [][]rune{{0x0001}, {0x0308}, {0x0378}}},                                                                 // ÷ [0.2] <START OF HEADING> (Control) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u034F\u0020", expected: [][]rune{{0x034F}, {0x0020}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u034F\u0308\u0020", expected: [][]rune{{0x034F, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u034F\u000D", expected: [][]rune{{0x034F}, {0x000D}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u034F\u0308\u000D", expected: [][]rune{{0x034F, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u034F\u000A", expected: [][]rune{{0x034F}, {0x000A}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u034F\u0308\u000A", expected: [][]rune{{0x034F, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u034F\u0001", expected: [][]rune{{0x034F}, {0x0001}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u034F\u0308\u0001", expected: [][]rune{{0x034F, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u034F\u034F", expected: [][]rune{{0x034F, 0x034F}}},                                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u034F\u0308\u034F", expected: [][]rune{{0x034F, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u034F\U0001F1E6", expected: [][]rune{{0x034F}, {0x1F1E6}}},                                                                            // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u034F\u0308\U0001F1E6", expected: [][]rune{{0x034F, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u034F\u0600", expected: [][]rune{{0x034F}, {0x0600}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u034F\u0308\u0600", expected: [][]rune{{0x034F, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u034F\u0903", expected: [][]rune{{0x034F, 0x0903}}},                                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u034F\u0308\u0903", expected: [][]rune{{0x034F, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u034F\u1100", expected: [][]rune{{0x034F}, {0x1100}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u034F\u0308\u1100", expected: [][]rune{{0x034F, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u034F\u1160", expected: [][]rune{{0x034F}, {0x1160}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u034F\u0308\u1160", expected: [][]rune{{0x034F, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u034F\u11A8", expected: [][]rune{{0x034F}, {0x11A8}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u034F\u0308\u11A8", expected: [][]rune{{0x034F, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u034F\uAC00", expected: [][]rune{{0x034F}, {0xAC00}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u034F\u0308\uAC00", expected: [][]rune{{0x034F, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u034F\uAC01", expected: [][]rune{{0x034F}, {0xAC01}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u034F\u0308\uAC01", expected: [][]rune{{0x034F, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u034F\u231A", expected: [][]rune{{0x034F}, {0x231A}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u034F\u0308\u231A", expected: [][]rune{{0x034F, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u034F\u0300", expected: [][]rune{{0x034F, 0x0300}}},                                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u034F\u0308\u0300", expected: [][]rune{{0x034F, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u034F\u200D", expected: [][]rune{{0x034F, 0x200D}}},                                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u034F\u0308\u200D", expected: [][]rune{{0x034F, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u034F\u0378", expected: [][]rune{{0x034F}, {0x0378}}},                                                                                 // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u034F\u0308\u0378", expected: [][]rune{{0x034F, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] COMBINING GRAPHEME JOINER (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\U0001F1E6\u0020", expected: [][]rune{{0x1F1E6}, {0x0020}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0020", expected: [][]rune{{0x1F1E6, 0x0308}, {0x0020}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\U0001F1E6\u000D", expected: [][]rune{{0x1F1E6}, {0x000D}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u000D", expected: [][]rune{{0x1F1E6, 0x0308}, {0x000D}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\U0001F1E6\u000A", expected: [][]rune{{0x1F1E6}, {0x000A}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u000A", expected: [][]rune{{0x1F1E6, 0x0308}, {0x000A}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\U0001F1E6\u0001", expected: [][]rune{{0x1F1E6}, {0x0001}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0001", expected: [][]rune{{0x1F1E6, 0x0308}, {0x0001}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\U0001F1E6\u034F", expected: [][]rune{{0x1F1E6, 0x034F}}},                                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u034F", expected: [][]rune{{0x1F1E6, 0x0308, 0x034F}}},                                                                // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\U0001F1E6\U0001F1E6", expected: [][]rune{{0x1F1E6, 0x1F1E6}}},                                                                         // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [12.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\U0001F1E6\u0308\U0001F1E6", expected: [][]rune{{0x1F1E6, 0x0308}, {0x1F1E6}}},                                                         // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\U0001F1E6\u0600", expected: [][]rune{{0x1F1E6}, {0x0600}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0600", expected: [][]rune{{0x1F1E6, 0x0308}, {0x0600}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\U0001F1E6\u0903", expected: [][]rune{{0x1F1E6, 0x0903}}},                                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0903", expected: [][]rune{{0x1F1E6, 0x0308, 0x0903}}},                                                                // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\U0001F1E6\u1100", expected: [][]rune{{0x1F1E6}, {0x1100}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u1100", expected: [][]rune{{0x1F1E6, 0x0308}, {0x1100}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\U0001F1E6\u1160", expected: [][]rune{{0x1F1E6}, {0x1160}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u1160", expected: [][]rune{{0x1F1E6, 0x0308}, {0x1160}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\U0001F1E6\u11A8", expected: [][]rune{{0x1F1E6}, {0x11A8}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u11A8", expected: [][]rune{{0x1F1E6, 0x0308}, {0x11A8}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\U0001F1E6\uAC00", expected: [][]rune{{0x1F1E6}, {0xAC00}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\U0001F1E6\u0308\uAC00", expected: [][]rune{{0x1F1E6, 0x0308}, {0xAC00}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\U0001F1E6\uAC01", expected: [][]rune{{0x1F1E6}, {0xAC01}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\U0001F1E6\u0308\uAC01", expected: [][]rune{{0x1F1E6, 0x0308}, {0xAC01}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\U0001F1E6\u231A", expected: [][]rune{{0x1F1E6}, {0x231A}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u231A", expected: [][]rune{{0x1F1E6, 0x0308}, {0x231A}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\U0001F1E6\u0300", expected: [][]rune{{0x1F1E6, 0x0300}}},                                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0300", expected: [][]rune{{0x1F1E6, 0x0308, 0x0300}}},                                                                // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\U0001F1E6\u200D", expected: [][]rune{{0x1F1E6, 0x200D}}},                                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u200D", expected: [][]rune{{0x1F1E6, 0x0308, 0x200D}}},                                                                // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\U0001F1E6\u0378", expected: [][]rune{{0x1F1E6}, {0x0378}}},                                                                            // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\U0001F1E6\u0308\u0378", expected: [][]rune{{0x1F1E6, 0x0308}, {0x0378}}},                                                              // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0600\u0020", expected: [][]rune{{0x0600, 0x0020}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] SPACE (Other) ÷ [0.3]
	{original: "\u0600\u0308\u0020", expected: [][]rune{{0x0600, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0600\u000D", expected: [][]rune{{0x0600}, {0x000D}}},                                                                                 // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0600\u0308\u000D", expected: [][]rune{{0x0600, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0600\u000A", expected: [][]rune{{0x0600}, {0x000A}}},                                                                                 // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0600\u0308\u000A", expected: [][]rune{{0x0600, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0600\u0001", expected: [][]rune{{0x0600}, {0x0001}}},                                                                                 // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0600\u0308\u0001", expected: [][]rune{{0x0600, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0600\u034F", expected: [][]rune{{0x0600, 0x034F}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0600\u0308\u034F", expected: [][]rune{{0x0600, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0600\U0001F1E6", expected: [][]rune{{0x0600, 0x1F1E6}}},                                                                              // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0600\u0308\U0001F1E6", expected: [][]rune{{0x0600, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0600\u0600", expected: [][]rune{{0x0600, 0x0600}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0600\u0308\u0600", expected: [][]rune{{0x0600, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0600\u0903", expected: [][]rune{{0x0600, 0x0903}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0600\u0308\u0903", expected: [][]rune{{0x0600, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0600\u1100", expected: [][]rune{{0x0600, 0x1100}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0600\u0308\u1100", expected: [][]rune{{0x0600, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0600\u1160", expected: [][]rune{{0x0600, 0x1160}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0600\u0308\u1160", expected: [][]rune{{0x0600, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0600\u11A8", expected: [][]rune{{0x0600, 0x11A8}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0600\u0308\u11A8", expected: [][]rune{{0x0600, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0600\uAC00", expected: [][]rune{{0x0600, 0xAC00}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0600\u0308\uAC00", expected: [][]rune{{0x0600, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0600\uAC01", expected: [][]rune{{0x0600, 0xAC01}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0600\u0308\uAC01", expected: [][]rune{{0x0600, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0600\u231A", expected: [][]rune{{0x0600, 0x231A}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0600\u0308\u231A", expected: [][]rune{{0x0600, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0600\u0300", expected: [][]rune{{0x0600, 0x0300}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0600\u0308\u0300", expected: [][]rune{{0x0600, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0600\u200D", expected: [][]rune{{0x0600, 0x200D}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0600\u0308\u200D", expected: [][]rune{{0x0600, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0600\u0378", expected: [][]rune{{0x0600, 0x0378}}},                                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.2] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0600\u0308\u0378", expected: [][]rune{{0x0600, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] ARABIC NUMBER SIGN (Prepend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0903\u0020", expected: [][]rune{{0x0903}, {0x0020}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0903\u0308\u0020", expected: [][]rune{{0x0903, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0903\u000D", expected: [][]rune{{0x0903}, {0x000D}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0903\u0308\u000D", expected: [][]rune{{0x0903, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0903\u000A", expected: [][]rune{{0x0903}, {0x000A}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0903\u0308\u000A", expected: [][]rune{{0x0903, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0903\u0001", expected: [][]rune{{0x0903}, {0x0001}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0903\u0308\u0001", expected: [][]rune{{0x0903, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0903\u034F", expected: [][]rune{{0x0903, 0x034F}}},                                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0903\u0308\u034F", expected: [][]rune{{0x0903, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0903\U0001F1E6", expected: [][]rune{{0x0903}, {0x1F1E6}}},                                                                            // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0903\u0308\U0001F1E6", expected: [][]rune{{0x0903, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0903\u0600", expected: [][]rune{{0x0903}, {0x0600}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0903\u0308\u0600", expected: [][]rune{{0x0903, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0903\u0903", expected: [][]rune{{0x0903, 0x0903}}},                                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0903\u0308\u0903", expected: [][]rune{{0x0903, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0903\u1100", expected: [][]rune{{0x0903}, {0x1100}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0903\u0308\u1100", expected: [][]rune{{0x0903, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0903\u1160", expected: [][]rune{{0x0903}, {0x1160}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0903\u0308\u1160", expected: [][]rune{{0x0903, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0903\u11A8", expected: [][]rune{{0x0903}, {0x11A8}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0903\u0308\u11A8", expected: [][]rune{{0x0903, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0903\uAC00", expected: [][]rune{{0x0903}, {0xAC00}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0903\u0308\uAC00", expected: [][]rune{{0x0903, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0903\uAC01", expected: [][]rune{{0x0903}, {0xAC01}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0903\u0308\uAC01", expected: [][]rune{{0x0903, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0903\u231A", expected: [][]rune{{0x0903}, {0x231A}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0903\u0308\u231A", expected: [][]rune{{0x0903, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0903\u0300", expected: [][]rune{{0x0903, 0x0300}}},                                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0903\u0308\u0300", expected: [][]rune{{0x0903, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0903\u200D", expected: [][]rune{{0x0903, 0x200D}}},                                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0903\u0308\u200D", expected: [][]rune{{0x0903, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0903\u0378", expected: [][]rune{{0x0903}, {0x0378}}},                                                                                 // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0903\u0308\u0378", expected: [][]rune{{0x0903, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] DEVANAGARI SIGN VISARGA (SpacingMark) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u1100\u0020", expected: [][]rune{{0x1100}, {0x0020}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u1100\u0308\u0020", expected: [][]rune{{0x1100, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u1100\u000D", expected: [][]rune{{0x1100}, {0x000D}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u1100\u0308\u000D", expected: [][]rune{{0x1100, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u1100\u000A", expected: [][]rune{{0x1100}, {0x000A}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u1100\u0308\u000A", expected: [][]rune{{0x1100, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u1100\u0001", expected: [][]rune{{0x1100}, {0x0001}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u1100\u0308\u0001", expected: [][]rune{{0x1100, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u1100\u034F", expected: [][]rune{{0x1100, 0x034F}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u1100\u0308\u034F", expected: [][]rune{{0x1100, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u1100\U0001F1E6", expected: [][]rune{{0x1100}, {0x1F1E6}}},                                                                            // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u1100\u0308\U0001F1E6", expected: [][]rune{{0x1100, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u1100\u0600", expected: [][]rune{{0x1100}, {0x0600}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u1100\u0308\u0600", expected: [][]rune{{0x1100, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u1100\u0903", expected: [][]rune{{0x1100, 0x0903}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u1100\u0308\u0903", expected: [][]rune{{0x1100, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u1100\u1100", expected: [][]rune{{0x1100, 0x1100}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [6.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u1100\u0308\u1100", expected: [][]rune{{0x1100, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u1100\u1160", expected: [][]rune{{0x1100, 0x1160}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [6.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u1100\u0308\u1160", expected: [][]rune{{0x1100, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u1100\u11A8", expected: [][]rune{{0x1100}, {0x11A8}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u1100\u0308\u11A8", expected: [][]rune{{0x1100, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u1100\uAC00", expected: [][]rune{{0x1100, 0xAC00}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [6.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u1100\u0308\uAC00", expected: [][]rune{{0x1100, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u1100\uAC01", expected: [][]rune{{0x1100, 0xAC01}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [6.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u1100\u0308\uAC01", expected: [][]rune{{0x1100, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u1100\u231A", expected: [][]rune{{0x1100}, {0x231A}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u1100\u0308\u231A", expected: [][]rune{{0x1100, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u1100\u0300", expected: [][]rune{{0x1100, 0x0300}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u1100\u0308\u0300", expected: [][]rune{{0x1100, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u1100\u200D", expected: [][]rune{{0x1100, 0x200D}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u1100\u0308\u200D", expected: [][]rune{{0x1100, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u1100\u0378", expected: [][]rune{{0x1100}, {0x0378}}},                                                                                 // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u1100\u0308\u0378", expected: [][]rune{{0x1100, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u1160\u0020", expected: [][]rune{{0x1160}, {0x0020}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u1160\u0308\u0020", expected: [][]rune{{0x1160, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u1160\u000D", expected: [][]rune{{0x1160}, {0x000D}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u1160\u0308\u000D", expected: [][]rune{{0x1160, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u1160\u000A", expected: [][]rune{{0x1160}, {0x000A}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u1160\u0308\u000A", expected: [][]rune{{0x1160, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u1160\u0001", expected: [][]rune{{0x1160}, {0x0001}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u1160\u0308\u0001", expected: [][]rune{{0x1160, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u1160\u034F", expected: [][]rune{{0x1160, 0x034F}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u1160\u0308\u034F", expected: [][]rune{{0x1160, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u1160\U0001F1E6", expected: [][]rune{{0x1160}, {0x1F1E6}}},                                                                            // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u1160\u0308\U0001F1E6", expected: [][]rune{{0x1160, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u1160\u0600", expected: [][]rune{{0x1160}, {0x0600}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u1160\u0308\u0600", expected: [][]rune{{0x1160, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u1160\u0903", expected: [][]rune{{0x1160, 0x0903}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u1160\u0308\u0903", expected: [][]rune{{0x1160, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u1160\u1100", expected: [][]rune{{0x1160}, {0x1100}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u1160\u0308\u1100", expected: [][]rune{{0x1160, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u1160\u1160", expected: [][]rune{{0x1160, 0x1160}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [7.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u1160\u0308\u1160", expected: [][]rune{{0x1160, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u1160\u11A8", expected: [][]rune{{0x1160, 0x11A8}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [7.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u1160\u0308\u11A8", expected: [][]rune{{0x1160, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u1160\uAC00", expected: [][]rune{{0x1160}, {0xAC00}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u1160\u0308\uAC00", expected: [][]rune{{0x1160, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u1160\uAC01", expected: [][]rune{{0x1160}, {0xAC01}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u1160\u0308\uAC01", expected: [][]rune{{0x1160, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u1160\u231A", expected: [][]rune{{0x1160}, {0x231A}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u1160\u0308\u231A", expected: [][]rune{{0x1160, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u1160\u0300", expected: [][]rune{{0x1160, 0x0300}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u1160\u0308\u0300", expected: [][]rune{{0x1160, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u1160\u200D", expected: [][]rune{{0x1160, 0x200D}}},                                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u1160\u0308\u200D", expected: [][]rune{{0x1160, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u1160\u0378", expected: [][]rune{{0x1160}, {0x0378}}},                                                                                 // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u1160\u0308\u0378", expected: [][]rune{{0x1160, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] HANGUL JUNGSEONG FILLER (V) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u11A8\u0020", expected: [][]rune{{0x11A8}, {0x0020}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u11A8\u0308\u0020", expected: [][]rune{{0x11A8, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u11A8\u000D", expected: [][]rune{{0x11A8}, {0x000D}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u11A8\u0308\u000D", expected: [][]rune{{0x11A8, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u11A8\u000A", expected: [][]rune{{0x11A8}, {0x000A}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u11A8\u0308\u000A", expected: [][]rune{{0x11A8, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u11A8\u0001", expected: [][]rune{{0x11A8}, {0x0001}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u11A8\u0308\u0001", expected: [][]rune{{0x11A8, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u11A8\u034F", expected: [][]rune{{0x11A8, 0x034F}}},                                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u11A8\u0308\u034F", expected: [][]rune{{0x11A8, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u11A8\U0001F1E6", expected: [][]rune{{0x11A8}, {0x1F1E6}}},                                                                            // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u11A8\u0308\U0001F1E6", expected: [][]rune{{0x11A8, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u11A8\u0600", expected: [][]rune{{0x11A8}, {0x0600}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u11A8\u0308\u0600", expected: [][]rune{{0x11A8, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u11A8\u0903", expected: [][]rune{{0x11A8, 0x0903}}},                                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u11A8\u0308\u0903", expected: [][]rune{{0x11A8, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u11A8\u1100", expected: [][]rune{{0x11A8}, {0x1100}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u11A8\u0308\u1100", expected: [][]rune{{0x11A8, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u11A8\u1160", expected: [][]rune{{0x11A8}, {0x1160}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u11A8\u0308\u1160", expected: [][]rune{{0x11A8, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u11A8\u11A8", expected: [][]rune{{0x11A8, 0x11A8}}},                                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [8.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u11A8\u0308\u11A8", expected: [][]rune{{0x11A8, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u11A8\uAC00", expected: [][]rune{{0x11A8}, {0xAC00}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u11A8\u0308\uAC00", expected: [][]rune{{0x11A8, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u11A8\uAC01", expected: [][]rune{{0x11A8}, {0xAC01}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u11A8\u0308\uAC01", expected: [][]rune{{0x11A8, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u11A8\u231A", expected: [][]rune{{0x11A8}, {0x231A}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u11A8\u0308\u231A", expected: [][]rune{{0x11A8, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u11A8\u0300", expected: [][]rune{{0x11A8, 0x0300}}},                                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u11A8\u0308\u0300", expected: [][]rune{{0x11A8, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u11A8\u200D", expected: [][]rune{{0x11A8, 0x200D}}},                                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u11A8\u0308\u200D", expected: [][]rune{{0x11A8, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u11A8\u0378", expected: [][]rune{{0x11A8}, {0x0378}}},                                                                                 // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u11A8\u0308\u0378", expected: [][]rune{{0x11A8, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] HANGUL JONGSEONG KIYEOK (T) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\uAC00\u0020", expected: [][]rune{{0xAC00}, {0x0020}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\uAC00\u0308\u0020", expected: [][]rune{{0xAC00, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\uAC00\u000D", expected: [][]rune{{0xAC00}, {0x000D}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\uAC00\u0308\u000D", expected: [][]rune{{0xAC00, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\uAC00\u000A", expected: [][]rune{{0xAC00}, {0x000A}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\uAC00\u0308\u000A", expected: [][]rune{{0xAC00, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\uAC00\u0001", expected: [][]rune{{0xAC00}, {0x0001}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\uAC00\u0308\u0001", expected: [][]rune{{0xAC00, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\uAC00\u034F", expected: [][]rune{{0xAC00, 0x034F}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\uAC00\u0308\u034F", expected: [][]rune{{0xAC00, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\uAC00\U0001F1E6", expected: [][]rune{{0xAC00}, {0x1F1E6}}},                                                                            // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\uAC00\u0308\U0001F1E6", expected: [][]rune{{0xAC00, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\uAC00\u0600", expected: [][]rune{{0xAC00}, {0x0600}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\uAC00\u0308\u0600", expected: [][]rune{{0xAC00, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\uAC00\u0903", expected: [][]rune{{0xAC00, 0x0903}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\uAC00\u0308\u0903", expected: [][]rune{{0xAC00, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\uAC00\u1100", expected: [][]rune{{0xAC00}, {0x1100}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC00\u0308\u1100", expected: [][]rune{{0xAC00, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC00\u1160", expected: [][]rune{{0xAC00, 0x1160}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [7.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\uAC00\u0308\u1160", expected: [][]rune{{0xAC00, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\uAC00\u11A8", expected: [][]rune{{0xAC00, 0x11A8}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [7.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\uAC00\u0308\u11A8", expected: [][]rune{{0xAC00, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\uAC00\uAC00", expected: [][]rune{{0xAC00}, {0xAC00}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\uAC00\u0308\uAC00", expected: [][]rune{{0xAC00, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\uAC00\uAC01", expected: [][]rune{{0xAC00}, {0xAC01}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\uAC00\u0308\uAC01", expected: [][]rune{{0xAC00, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\uAC00\u231A", expected: [][]rune{{0xAC00}, {0x231A}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\uAC00\u0308\u231A", expected: [][]rune{{0xAC00, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\uAC00\u0300", expected: [][]rune{{0xAC00, 0x0300}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\uAC00\u0308\u0300", expected: [][]rune{{0xAC00, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\uAC00\u200D", expected: [][]rune{{0xAC00, 0x200D}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\uAC00\u0308\u200D", expected: [][]rune{{0xAC00, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\uAC00\u0378", expected: [][]rune{{0xAC00}, {0x0378}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GA (LV) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\uAC00\u0308\u0378", expected: [][]rune{{0xAC00, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\uAC01\u0020", expected: [][]rune{{0xAC01}, {0x0020}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\uAC01\u0308\u0020", expected: [][]rune{{0xAC01, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\uAC01\u000D", expected: [][]rune{{0xAC01}, {0x000D}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\uAC01\u0308\u000D", expected: [][]rune{{0xAC01, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\uAC01\u000A", expected: [][]rune{{0xAC01}, {0x000A}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\uAC01\u0308\u000A", expected: [][]rune{{0xAC01, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\uAC01\u0001", expected: [][]rune{{0xAC01}, {0x0001}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\uAC01\u0308\u0001", expected: [][]rune{{0xAC01, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\uAC01\u034F", expected: [][]rune{{0xAC01, 0x034F}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\uAC01\u0308\u034F", expected: [][]rune{{0xAC01, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\uAC01\U0001F1E6", expected: [][]rune{{0xAC01}, {0x1F1E6}}},                                                                            // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\uAC01\u0308\U0001F1E6", expected: [][]rune{{0xAC01, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\uAC01\u0600", expected: [][]rune{{0xAC01}, {0x0600}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\uAC01\u0308\u0600", expected: [][]rune{{0xAC01, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\uAC01\u0903", expected: [][]rune{{0xAC01, 0x0903}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\uAC01\u0308\u0903", expected: [][]rune{{0xAC01, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\uAC01\u1100", expected: [][]rune{{0xAC01}, {0x1100}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC01\u0308\u1100", expected: [][]rune{{0xAC01, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC01\u1160", expected: [][]rune{{0xAC01}, {0x1160}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\uAC01\u0308\u1160", expected: [][]rune{{0xAC01, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\uAC01\u11A8", expected: [][]rune{{0xAC01, 0x11A8}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [8.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\uAC01\u0308\u11A8", expected: [][]rune{{0xAC01, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\uAC01\uAC00", expected: [][]rune{{0xAC01}, {0xAC00}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\uAC01\u0308\uAC00", expected: [][]rune{{0xAC01, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\uAC01\uAC01", expected: [][]rune{{0xAC01}, {0xAC01}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\uAC01\u0308\uAC01", expected: [][]rune{{0xAC01, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\uAC01\u231A", expected: [][]rune{{0xAC01}, {0x231A}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\uAC01\u0308\u231A", expected: [][]rune{{0xAC01, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\uAC01\u0300", expected: [][]rune{{0xAC01, 0x0300}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\uAC01\u0308\u0300", expected: [][]rune{{0xAC01, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\uAC01\u200D", expected: [][]rune{{0xAC01, 0x200D}}},                                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\uAC01\u0308\u200D", expected: [][]rune{{0xAC01, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\uAC01\u0378", expected: [][]rune{{0xAC01}, {0x0378}}},                                                                                 // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\uAC01\u0308\u0378", expected: [][]rune{{0xAC01, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u231A\u0020", expected: [][]rune{{0x231A}, {0x0020}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u231A\u0308\u0020", expected: [][]rune{{0x231A, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u231A\u000D", expected: [][]rune{{0x231A}, {0x000D}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u231A\u0308\u000D", expected: [][]rune{{0x231A, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u231A\u000A", expected: [][]rune{{0x231A}, {0x000A}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u231A\u0308\u000A", expected: [][]rune{{0x231A, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u231A\u0001", expected: [][]rune{{0x231A}, {0x0001}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u231A\u0308\u0001", expected: [][]rune{{0x231A, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u231A\u034F", expected: [][]rune{{0x231A, 0x034F}}},                                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u231A\u0308\u034F", expected: [][]rune{{0x231A, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u231A\U0001F1E6", expected: [][]rune{{0x231A}, {0x1F1E6}}},                                                                            // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u231A\u0308\U0001F1E6", expected: [][]rune{{0x231A, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u231A\u0600", expected: [][]rune{{0x231A}, {0x0600}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u231A\u0308\u0600", expected: [][]rune{{0x231A, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u231A\u0903", expected: [][]rune{{0x231A, 0x0903}}},                                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u231A\u0308\u0903", expected: [][]rune{{0x231A, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u231A\u1100", expected: [][]rune{{0x231A}, {0x1100}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u231A\u0308\u1100", expected: [][]rune{{0x231A, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u231A\u1160", expected: [][]rune{{0x231A}, {0x1160}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u231A\u0308\u1160", expected: [][]rune{{0x231A, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u231A\u11A8", expected: [][]rune{{0x231A}, {0x11A8}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u231A\u0308\u11A8", expected: [][]rune{{0x231A, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u231A\uAC00", expected: [][]rune{{0x231A}, {0xAC00}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u231A\u0308\uAC00", expected: [][]rune{{0x231A, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u231A\uAC01", expected: [][]rune{{0x231A}, {0xAC01}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u231A\u0308\uAC01", expected: [][]rune{{0x231A, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u231A\u231A", expected: [][]rune{{0x231A}, {0x231A}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u231A\u0308\u231A", expected: [][]rune{{0x231A, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u231A\u0300", expected: [][]rune{{0x231A, 0x0300}}},                                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u231A\u0308\u0300", expected: [][]rune{{0x231A, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u231A\u200D", expected: [][]rune{{0x231A, 0x200D}}},                                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u231A\u0308\u200D", expected: [][]rune{{0x231A, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u231A\u0378", expected: [][]rune{{0x231A}, {0x0378}}},                                                                                 // ÷ [0.2] WATCH (ExtPict) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u231A\u0308\u0378", expected: [][]rune{{0x231A, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] WATCH (ExtPict) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0300\u0020", expected: [][]rune{{0x0300}, {0x0020}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0300\u0308\u0020", expected: [][]rune{{0x0300, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0300\u000D", expected: [][]rune{{0x0300}, {0x000D}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0300\u0308\u000D", expected: [][]rune{{0x0300, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0300\u000A", expected: [][]rune{{0x0300}, {0x000A}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0300\u0308\u000A", expected: [][]rune{{0x0300, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0300\u0001", expected: [][]rune{{0x0300}, {0x0001}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0300\u0308\u0001", expected: [][]rune{{0x0300, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0300\u034F", expected: [][]rune{{0x0300, 0x034F}}},                                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0300\u0308\u034F", expected: [][]rune{{0x0300, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0300\U0001F1E6", expected: [][]rune{{0x0300}, {0x1F1E6}}},                                                                            // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0300\u0308\U0001F1E6", expected: [][]rune{{0x0300, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0300\u0600", expected: [][]rune{{0x0300}, {0x0600}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0300\u0308\u0600", expected: [][]rune{{0x0300, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0300\u0903", expected: [][]rune{{0x0300, 0x0903}}},                                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0300\u0308\u0903", expected: [][]rune{{0x0300, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0300\u1100", expected: [][]rune{{0x0300}, {0x1100}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0300\u0308\u1100", expected: [][]rune{{0x0300, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0300\u1160", expected: [][]rune{{0x0300}, {0x1160}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0300\u0308\u1160", expected: [][]rune{{0x0300, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0300\u11A8", expected: [][]rune{{0x0300}, {0x11A8}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0300\u0308\u11A8", expected: [][]rune{{0x0300, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0300\uAC00", expected: [][]rune{{0x0300}, {0xAC00}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0300\u0308\uAC00", expected: [][]rune{{0x0300, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0300\uAC01", expected: [][]rune{{0x0300}, {0xAC01}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0300\u0308\uAC01", expected: [][]rune{{0x0300, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0300\u231A", expected: [][]rune{{0x0300}, {0x231A}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0300\u0308\u231A", expected: [][]rune{{0x0300, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0300\u0300", expected: [][]rune{{0x0300, 0x0300}}},                                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0300\u0308\u0300", expected: [][]rune{{0x0300, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0300\u200D", expected: [][]rune{{0x0300, 0x200D}}},                                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0300\u0308\u200D", expected: [][]rune{{0x0300, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0300\u0378", expected: [][]rune{{0x0300}, {0x0378}}},                                                                                 // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0300\u0308\u0378", expected: [][]rune{{0x0300, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u200D\u0020", expected: [][]rune{{0x200D}, {0x0020}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u200D\u0308\u0020", expected: [][]rune{{0x200D, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u200D\u000D", expected: [][]rune{{0x200D}, {0x000D}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u200D\u0308\u000D", expected: [][]rune{{0x200D, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u200D\u000A", expected: [][]rune{{0x200D}, {0x000A}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u200D\u0308\u000A", expected: [][]rune{{0x200D, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u200D\u0001", expected: [][]rune{{0x200D}, {0x0001}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u200D\u0308\u0001", expected: [][]rune{{0x200D, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u200D\u034F", expected: [][]rune{{0x200D, 0x034F}}},                                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u200D\u0308\u034F", expected: [][]rune{{0x200D, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u200D\U0001F1E6", expected: [][]rune{{0x200D}, {0x1F1E6}}},                                                                            // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u200D\u0308\U0001F1E6", expected: [][]rune{{0x200D, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u200D\u0600", expected: [][]rune{{0x200D}, {0x0600}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u200D\u0308\u0600", expected: [][]rune{{0x200D, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u200D\u0903", expected: [][]rune{{0x200D, 0x0903}}},                                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u200D\u0308\u0903", expected: [][]rune{{0x200D, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u200D\u1100", expected: [][]rune{{0x200D}, {0x1100}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u200D\u0308\u1100", expected: [][]rune{{0x200D, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u200D\u1160", expected: [][]rune{{0x200D}, {0x1160}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u200D\u0308\u1160", expected: [][]rune{{0x200D, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u200D\u11A8", expected: [][]rune{{0x200D}, {0x11A8}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u200D\u0308\u11A8", expected: [][]rune{{0x200D, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u200D\uAC00", expected: [][]rune{{0x200D}, {0xAC00}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u200D\u0308\uAC00", expected: [][]rune{{0x200D, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u200D\uAC01", expected: [][]rune{{0x200D}, {0xAC01}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u200D\u0308\uAC01", expected: [][]rune{{0x200D, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u200D\u231A", expected: [][]rune{{0x200D}, {0x231A}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u200D\u0308\u231A", expected: [][]rune{{0x200D, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u200D\u0300", expected: [][]rune{{0x200D, 0x0300}}},                                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u200D\u0308\u0300", expected: [][]rune{{0x200D, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u200D\u200D", expected: [][]rune{{0x200D, 0x200D}}},                                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u200D\u0308\u200D", expected: [][]rune{{0x200D, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u200D\u0378", expected: [][]rune{{0x200D}, {0x0378}}},                                                                                 // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u200D\u0308\u0378", expected: [][]rune{{0x200D, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0378\u0020", expected: [][]rune{{0x0378}, {0x0020}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0378\u0308\u0020", expected: [][]rune{{0x0378, 0x0308}, {0x0020}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u0378\u000D", expected: [][]rune{{0x0378}, {0x000D}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0378\u0308\u000D", expected: [][]rune{{0x0378, 0x0308}, {0x000D}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <CARRIAGE RETURN (CR)> (CR) ÷ [0.3]
	{original: "\u0378\u000A", expected: [][]rune{{0x0378}, {0x000A}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0378\u0308\u000A", expected: [][]rune{{0x0378, 0x0308}, {0x000A}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [0.3]
	{original: "\u0378\u0001", expected: [][]rune{{0x0378}, {0x0001}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0378\u0308\u0001", expected: [][]rune{{0x0378, 0x0308}, {0x0001}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [5.0] <START OF HEADING> (Control) ÷ [0.3]
	{original: "\u0378\u034F", expected: [][]rune{{0x0378, 0x034F}}},                                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0378\u0308\u034F", expected: [][]rune{{0x0378, 0x0308, 0x034F}}},                                                                     // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAPHEME JOINER (Extend) ÷ [0.3]
	{original: "\u0378\U0001F1E6", expected: [][]rune{{0x0378}, {0x1F1E6}}},                                                                            // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0378\u0308\U0001F1E6", expected: [][]rune{{0x0378, 0x0308}, {0x1F1E6}}},                                                              // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) ÷ [0.3]
	{original: "\u0378\u0600", expected: [][]rune{{0x0378}, {0x0600}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0378\u0308\u0600", expected: [][]rune{{0x0378, 0x0308}, {0x0600}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) ÷ [0.3]
	{original: "\u0378\u0903", expected: [][]rune{{0x0378, 0x0903}}},                                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0378\u0308\u0903", expected: [][]rune{{0x0378, 0x0308, 0x0903}}},                                                                     // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [0.3]
	{original: "\u0378\u1100", expected: [][]rune{{0x0378}, {0x1100}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0378\u0308\u1100", expected: [][]rune{{0x0378, 0x0308}, {0x1100}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\u0378\u1160", expected: [][]rune{{0x0378}, {0x1160}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0378\u0308\u1160", expected: [][]rune{{0x0378, 0x0308}, {0x1160}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JUNGSEONG FILLER (V) ÷ [0.3]
	{original: "\u0378\u11A8", expected: [][]rune{{0x0378}, {0x11A8}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0378\u0308\u11A8", expected: [][]rune{{0x0378, 0x0308}, {0x11A8}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL JONGSEONG KIYEOK (T) ÷ [0.3]
	{original: "\u0378\uAC00", expected: [][]rune{{0x0378}, {0xAC00}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0378\u0308\uAC00", expected: [][]rune{{0x0378, 0x0308}, {0xAC00}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GA (LV) ÷ [0.3]
	{original: "\u0378\uAC01", expected: [][]rune{{0x0378}, {0xAC01}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0378\u0308\uAC01", expected: [][]rune{{0x0378, 0x0308}, {0xAC01}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] HANGUL SYLLABLE GAG (LVT) ÷ [0.3]
	{original: "\u0378\u231A", expected: [][]rune{{0x0378}, {0x231A}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0378\u0308\u231A", expected: [][]rune{{0x0378, 0x0308}, {0x231A}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] WATCH (ExtPict) ÷ [0.3]
	{original: "\u0378\u0300", expected: [][]rune{{0x0378, 0x0300}}},                                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0378\u0308\u0300", expected: [][]rune{{0x0378, 0x0308, 0x0300}}},                                                                     // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] COMBINING GRAVE ACCENT (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0378\u200D", expected: [][]rune{{0x0378, 0x200D}}},                                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0378\u0308\u200D", expected: [][]rune{{0x0378, 0x0308, 0x200D}}},                                                                     // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0378\u0378", expected: [][]rune{{0x0378}, {0x0378}}},                                                                                 // ÷ [0.2] <reserved-0378> (Other) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u0378\u0308\u0378", expected: [][]rune{{0x0378, 0x0308}, {0x0378}}},                                                                   // ÷ [0.2] <reserved-0378> (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] <reserved-0378> (Other) ÷ [0.3]
	{original: "\u000D\u000A\u0061\u000A\u0308", expected: [][]rune{{0x000D, 0x000A}, {0x0061}, {0x000A}, {0x0308}}},                                   // ÷ [0.2] <CARRIAGE RETURN (CR)> (CR) × [3.0] <LINE FEED (LF)> (LF) ÷ [4.0] LATIN SMALL LETTER A (Other) ÷ [5.0] <LINE FEED (LF)> (LF) ÷ [4.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0061\u0308", expected: [][]rune{{0x0061, 0x0308}}},                                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [0.3]
	{original: "\u0020\u200D\u0646", expected: [][]rune{{0x0020, 0x200D}, {0x0646}}},                                                                   // ÷ [0.2] SPACE (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] ARABIC LETTER NOON (Other) ÷ [0.3]
	{original: "\u0646\u200D\u0020", expected: [][]rune{{0x0646, 0x200D}, {0x0020}}},                                                                   // ÷ [0.2] ARABIC LETTER NOON (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] SPACE (Other) ÷ [0.3]
	{original: "\u1100\u1100", expected: [][]rune{{0x1100, 0x1100}}},                                                                                   // ÷ [0.2] HANGUL CHOSEONG KIYEOK (L) × [6.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC00\u11A8\u1100", expected: [][]rune{{0xAC00, 0x11A8}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GA (LV) × [7.0] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\uAC01\u11A8\u1100", expected: [][]rune{{0xAC01, 0x11A8}, {0x1100}}},                                                                   // ÷ [0.2] HANGUL SYLLABLE GAG (LVT) × [8.0] HANGUL JONGSEONG KIYEOK (T) ÷ [999.0] HANGUL CHOSEONG KIYEOK (L) ÷ [0.3]
	{original: "\U0001F1E6\U0001F1E7\U0001F1E8\u0062", expected: [][]rune{{0x1F1E6, 0x1F1E7}, {0x1F1E8}, {0x0062}}},                                    // ÷ [0.2] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [12.0] REGIONAL INDICATOR SYMBOL LETTER B (RI) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER C (RI) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\U0001F1E6\U0001F1E7\U0001F1E8\u0062", expected: [][]rune{{0x0061}, {0x1F1E6, 0x1F1E7}, {0x1F1E8}, {0x0062}}},                    // ÷ [0.2] LATIN SMALL LETTER A (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [13.0] REGIONAL INDICATOR SYMBOL LETTER B (RI) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER C (RI) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\U0001F1E6\U0001F1E7\u200D\U0001F1E8\u0062", expected: [][]rune{{0x0061}, {0x1F1E6, 0x1F1E7, 0x200D}, {0x1F1E8}, {0x0062}}},      // ÷ [0.2] LATIN SMALL LETTER A (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [13.0] REGIONAL INDICATOR SYMBOL LETTER B (RI) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER C (RI) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\U0001F1E6\u200D\U0001F1E7\U0001F1E8\u0062", expected: [][]rune{{0x0061}, {0x1F1E6, 0x200D}, {0x1F1E7, 0x1F1E8}, {0x0062}}},      // ÷ [0.2] LATIN SMALL LETTER A (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER B (RI) × [13.0] REGIONAL INDICATOR SYMBOL LETTER C (RI) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\U0001F1E6\U0001F1E7\U0001F1E8\U0001F1E9\u0062", expected: [][]rune{{0x0061}, {0x1F1E6, 0x1F1E7}, {0x1F1E8, 0x1F1E9}, {0x0062}}}, // ÷ [0.2] LATIN SMALL LETTER A (Other) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER A (RI) × [13.0] REGIONAL INDICATOR SYMBOL LETTER B (RI) ÷ [999.0] REGIONAL INDICATOR SYMBOL LETTER C (RI) × [13.0] REGIONAL INDICATOR SYMBOL LETTER D (RI) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\u200D", expected: [][]rune{{0x0061, 0x200D}}},                                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [0.3]
	{original: "\u0061\u0308\u0062", expected: [][]rune{{0x0061, 0x0308}, {0x0062}}},                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\u0903\u0062", expected: [][]rune{{0x0061, 0x0903}, {0x0062}}},                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.1] DEVANAGARI SIGN VISARGA (SpacingMark) ÷ [999.0] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\u0061\u0600\u0062", expected: [][]rune{{0x0061}, {0x0600, 0x0062}}},                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) ÷ [999.0] ARABIC NUMBER SIGN (Prepend) × [9.2] LATIN SMALL LETTER B (Other) ÷ [0.3]
	{original: "\U0001F476\U0001F3FF\U0001F476", expected: [][]rune{{0x1F476, 0x1F3FF}, {0x1F476}}},                                                    // ÷ [0.2] BABY (ExtPict) × [9.0] EMOJI MODIFIER FITZPATRICK TYPE-6 (Extend) ÷ [999.0] BABY (ExtPict) ÷ [0.3]
	{original: "\u0061\U0001F3FF\U0001F476", expected: [][]rune{{0x0061, 0x1F3FF}, {0x1F476}}},                                                         // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] EMOJI MODIFIER FITZPATRICK TYPE-6 (Extend) ÷ [999.0] BABY (ExtPict) ÷ [0.3]
	{original: "\u0061\U0001F3FF\U0001F476\u200D\U0001F6D1", expected: [][]rune{{0x0061, 0x1F3FF}, {0x1F476, 0x200D, 0x1F6D1}}},                        // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] EMOJI MODIFIER FITZPATRICK TYPE-6 (Extend) ÷ [999.0] BABY (ExtPict) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [11.0] OCTAGONAL SIGN (ExtPict) ÷ [0.3]
	{original: "\U0001F476\U0001F3FF\u0308\u200D\U0001F476\U0001F3FF", expected: [][]rune{{0x1F476, 0x1F3FF, 0x0308, 0x200D, 0x1F476, 0x1F3FF}}},       // ÷ [0.2] BABY (ExtPict) × [9.0] EMOJI MODIFIER FITZPATRICK TYPE-6 (Extend) × [9.0] COMBINING DIAERESIS (Extend_ExtCccZwj) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [11.0] BABY (ExtPict) × [9.0] EMOJI MODIFIER FITZPATRICK TYPE-6 (Extend) ÷ [0.3]
	{original: "\U0001F6D1\u200D\U0001F6D1", expected: [][]rune{{0x1F6D1, 0x200D, 0x1F6D1}}},                                                           // ÷ [0.2] OCTAGONAL SIGN (ExtPict) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [11.0] OCTAGONAL SIGN (ExtPict) ÷ [0.3]
	{original: "\u0061\u200D\U0001F6D1", expected: [][]rune{{0x0061, 0x200D}, {0x1F6D1}}},                                                              // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] OCTAGONAL SIGN (ExtPict) ÷ [0.3]
	{original: "\u2701\u200D\u2701", expected: [][]rune{{0x2701, 0x200D, 0x2701}}},                                                                     // ÷ [0.2] UPPER BLADE SCISSORS (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) × [11.0] UPPER BLADE SCISSORS (Other) ÷ [0.3]
	{original: "\u0061\u200D\u2701", expected: [][]rune{{0x0061, 0x200D}, {0x2701}}},                                                                   // ÷ [0.2] LATIN SMALL LETTER A (Other) × [9.0] ZERO WIDTH JOINER (ZWJ_ExtCccZwj) ÷ [999.0] UPPER BLADE SCISSORS (Other) ÷ [0.3]
}

// decomposed returns a grapheme cluster decomposition.
func decomposed(s string) (runes [][]rune) {
	gr := NewGraphemes(s)
	for gr.Next() {
		runes = append(runes, gr.Runes())
	}
	return
}

// Run the testCases slice above.
func TestSimple(t *testing.T) {
	for testNum, testCase := range testCases {
		/*t.Logf(`Test case %d "%s": Expecting %x, getting %x, code points %x"`,
		testNum,
		strings.TrimSpace(testCase.original),
		testCase.expected,
		decomposed(testCase.original),
		[]rune(testCase.original))*/
		gr := NewGraphemes(testCase.original)
		var index int
	GraphemeLoop:
		for index = 0; gr.Next(); index++ {
			if index >= len(testCase.expected) {
				t.Errorf(`Test case %d "%s" failed: More grapheme clusters returned than expected %d`,
					testNum,
					testCase.original,
					len(testCase.expected))
				break
			}
			cluster := gr.Runes()
			if len(cluster) != len(testCase.expected[index]) {
				t.Errorf(`Test case %d "%s" failed: Grapheme cluster at index %d has %d codepoints %x, %d expected %x`,
					testNum,
					testCase.original,
					index,
					len(cluster),
					cluster,
					len(testCase.expected[index]),
					testCase.expected[index])
				break
			}
			for i, r := range cluster {
				if r != testCase.expected[index][i] {
					t.Errorf(`Test case %d "%s" failed: Grapheme cluster at index %d is %x, expected %x`,
						testNum,
						testCase.original,
						index,
						cluster,
						testCase.expected[index])
					break GraphemeLoop
				}
			}
		}
		if index < len(testCase.expected) {
			t.Errorf(`Test case %d "%s" failed: Fewer grapheme clusters returned (%d) than expected (%d)`,
				testNum,
				testCase.original,
				index,
				len(testCase.expected))
		}
	}
}

// Test the Str() function.
func TestStr(t *testing.T) {
	gr := NewGraphemes("möp")
	gr.Next()
	gr.Next()
	gr.Next()
	if str := gr.Str(); str != "p" {
		t.Errorf(`Expected "p", got "%s"`, str)
	}
}

// Test the Bytes() function.
func TestBytes(t *testing.T) {
	gr := NewGraphemes("A👩‍❤️‍💋‍👩B")
	gr.Next()
	gr.Next()
	gr.Next()
	b := gr.Bytes()
	if len(b) != 1 {
		t.Fatalf(`Expected len("B") == 1, got %d`, len(b))
	}
	if b[0] != 'B' {
		t.Errorf(`Expected "B", got "%s"`, string(b[0]))
	}
}

// Test the Positions() function.
func TestPositions(t *testing.T) {
	gr := NewGraphemes("A👩‍❤️‍💋‍👩B")
	gr.Next()
	gr.Next()
	from, to := gr.Positions()
	if from != 1 || to != 28 {
		t.Errorf(`Expected from=%d to=%d, got from=%d to=%d`, 1, 28, from, to)
	}
}

// Test the Reset() function.
func TestReset(t *testing.T) {
	gr := NewGraphemes("möp")
	gr.Next()
	gr.Next()
	gr.Next()
	gr.Reset()
	gr.Next()
	if str := gr.Str(); str != "m" {
		t.Errorf(`Expected "m", got "%s"`, str)
	}
}

// Test retrieving clusters before calling Next().
func TestEarly(t *testing.T) {
	gr := NewGraphemes("test")
	r := gr.Runes()
	if r != nil {
		t.Errorf(`Expected nil rune slice, got %x`, r)
	}
	str := gr.Str()
	if str != "" {
		t.Errorf(`Expected empty string, got "%s"`, str)
	}
	b := gr.Bytes()
	if b != nil {
		t.Errorf(`Expected byte rune slice, got %x`, b)
	}
	from, to := gr.Positions()
	if from != 0 || to != 0 {
		t.Errorf(`Expected from=%d to=%d, got from=%d to=%d`, 0, 0, from, to)
	}
}

// Test retrieving more clusters after retrieving the last cluster.
func TestLate(t *testing.T) {
	gr := NewGraphemes("x")
	gr.Next()
	gr.Next()
	r := gr.Runes()
	if r != nil {
		t.Errorf(`Expected nil rune slice, got %x`, r)
	}
	str := gr.Str()
	if str != "" {
		t.Errorf(`Expected empty string, got "%s"`, str)
	}
	b := gr.Bytes()
	if b != nil {
		t.Errorf(`Expected byte rune slice, got %x`, b)
	}
	from, to := gr.Positions()
	if from != 1 || to != 1 {
		t.Errorf(`Expected from=%d to=%d, got from=%d to=%d`, 1, 1, from, to)
	}
}

// Test the GraphemeClusterCount function.
func TestCount(t *testing.T) {
	if n := GraphemeClusterCount("🇩🇪🏳️‍🌈"); n != 2 {
		t.Errorf(`Expected 2 grapheme clusters, got %d`, n)
	}
}
