package utils

import (
	"github.com/khalt00/noitu/internal/dict"
	"testing"
)

func TestCombineString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"No words", []string{}, ""},
		{"One word", []string{"hello"}, "hello"},
		{"Multiple words", []string{"hello", "world"}, "hello world"},
		{"Empty strings", []string{"", "world"}, " world"},
		{"Spaces only", []string{"  ", "world"}, "   world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineString(tt.input...)
			if result != tt.expected {
				t.Errorf("CombineString(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetFirstConnectWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single word", "ăn", "ăn"},
		{"Multiple words", "ăn bám", "bám"},
		{"Multiple words with extra spaces", "   ăn    bám    ", "bám"},
		{"Empty string", "", ""},
		{"Single space", " ", ""},
		{"Multiple spaces", "   ", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFirstConnectWord(tt.input)
			if result != tt.expected {
				t.Errorf("GetConnectWord(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
func TestGetSecondConnectWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single word", "à", "à"},
		{"Multiple words", "à ơi", "à"},
		{"Multiple words with extra spaces", "   à   ơi   ", "à"},
		{"Empty string", "", ""},
		{"Single space", " ", ""},
		{"Multiple spaces", "   ", ""},
		{"Trailing spaces", "   à             ơi   ", "à"},
		{"No words", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetSecondConnectWord(tt.input)
			if result != tt.expected {
				t.Errorf("GetSecondConnectWord(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCompareCorrectConnectWord(t *testing.T) {
	dict.InitDict()
	tests := []struct {
		name           string
		first          string
		second         string
		expectedResult bool
	}{
		{"Both words valid and match", "ăn", "ăn", true},
		{"Second word invalid", "ăn", "uống", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareCorrectConnectWord(tt.first, tt.second)
			if result != tt.expectedResult {
				t.Errorf("CompareCorrectConnectWord(%q, %q) = %v; expected %v", tt.first, tt.second, result, tt.expectedResult)
			}
		})
	}
}
