// A Go string is a read-only slice of bytes. The language and the standard library
// treat string specially - as cointainers of text encoded in UTF-8.
// In Go, the concept of a character is called a rune - it's an integer that
// represents a Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// word "hello" in the Thai language.
	const s = "สวัสดี"

	// Since strings are equivalent to []byte, this will produce
	// the length of the raw bytes stored within
	fmt.Println("Len: ", len(s))

	// This loop generates teh hex values of all the bytes that constitute the code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	
	fmt.Println()

	// To count how many *runes* are in a string, there is a utf8 package.
	// Some Thai characters are represented by multiple UTF-8 code points
	// so the result of this count may be surprising
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	// It decodes each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInStrings")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, widdth := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = widdth

		examineRune(runeValue)
	}
}

	func examineRune(r rune) {

		// values enclosed in single quotes are rune literals.
		// We can compare a rune value to a rune literal directly
		if r == 't' {
			fmt.Println("found tee")
		} else if r == 'ส' {
			fmt.Println("found so sua")
		}
	}