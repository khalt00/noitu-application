package utils

import (
	"fmt"
	"strings"

	"github.com/khalt00/noitu/internal/dict"
)

func CombineString(rest ...string) string {
	return strings.Join(rest, " ")
}

// Get connect word, if it only have 1 word
// Ex: à => return à
// if it have more than 1 word get last word
// Ex: ăn bám => return bám
func GetFirstConnectWord(word string) string {
	temp := strings.Fields(word)
	fmt.Println(temp)

	if len(temp) == 0 {
		return ""
	}
	if len(temp) == 1 {
		return temp[0]
	}
	return temp[len(temp)-1]
}

// Get connect word
// Ex: à => à
// if it have more than 1 word, get first word
func GetSecondConnectWord(word string) string {
	temp := strings.Fields(word)

	if len(temp) <= 0 {
		return ""
	}
	return temp[0]
}

// First word is the current word
// Second word is the incoming word
// Check if second word exists in dictionary
// Get Connect word
func CompareCorrectConnectWord(first, second string) bool {
	existInDict := dict.IsValidWord(second)
	if !existInDict {
		return false
	}
	firstConnectWord := GetFirstConnectWord(first)
	secondConnectWord := GetSecondConnectWord(second)

	return firstConnectWord == secondConnectWord
}
