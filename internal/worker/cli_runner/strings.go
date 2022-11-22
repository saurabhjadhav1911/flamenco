package cli_runner

import (
	"unicode/utf8"
)

// splitOnCharacterBoundary splits `b` such that `valid` + `tail` = `b` and
// `valid` is valid UTF-8.
func splitOnCharacterBoundary(b []byte) (valid []byte, tail []byte) {
	totalLength := len(b)
	tailBytes := 0
	for {
		valid = b[:totalLength-tailBytes]
		r, size := utf8.DecodeLastRune(valid)
		switch {
		case r == utf8.RuneError && size == 0:
			// valid is empty, which means 'b' consists of only non-UTF8 bytes.
			return valid, b
		case r == utf8.RuneError && size == 1:
			// The last bytes do not form a valid rune. See what happens if we move
			// one byte from `valid` to `tail`.
			tailBytes++
			continue
		case r == utf8.RuneError:
			// This shouldn't happen, RuneError should only be returned with size 0 or 1.
			panic(size)
		default:
			return valid, b[totalLength-tailBytes:]
		}
	}
}
