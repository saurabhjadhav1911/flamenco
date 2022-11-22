package cli_runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitOnCharacterBoundary(t *testing.T) {
	// Test with strings, as those are easier to type.
	tests := []struct {
		name      string
		input     string
		wantValid string
		wantTail  string
	}{
		{"empty", "", "", ""},
		{"trivial", "abc", "abc", ""},
		{"valid", "Stüvel", "Stüvel", ""},
		{"cats", "🐈🐈\xf0\x9f\x90\x88", "🐈🐈🐈", ""},
		{"truncated-cats-1", "🐈🐈\xf0\x9f\x90", "🐈🐈", "\xf0\x9f\x90"},
		{"truncated-cats-2", "🐈🐈\xf0\x9f", "🐈🐈", "\xf0\x9f"},
		{"truncated-cats-3", "🐈🐈\xf0", "🐈🐈", "\xf0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotTail := splitOnCharacterBoundary([]byte(tt.input))
			assert.Equal(t, tt.input, string(gotValid)+string(gotTail))
			assert.Equal(t, tt.wantValid, string(gotValid))
			assert.Equal(t, tt.wantTail, string(gotTail))
		})
	}
}
