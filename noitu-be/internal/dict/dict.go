package dict

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

var Dict map[string]bool

func InitDict() {
	slog.Info("Starting to init dict")
	Dict = GetDict()
	slog.Info("Init successfully")
}

func GetDictBuffer() map[string]bool {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	words := make(map[string]bool, 1024)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var jsonObject map[string]string
		err := json.Unmarshal([]byte(scanner.Text()), &jsonObject)
		if err != nil {
			slog.Info("err parsing", slog.Any("err", err))
			continue
		}

		if text, exists := jsonObject["text"]; exists {
			words[text] = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return words
}

func GetDict() map[string]bool {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	words := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var jsonObject map[string]string
		err := json.Unmarshal([]byte(scanner.Text()), &jsonObject)
		if err != nil {
			slog.Info("err parsing", slog.Any("err", err))
			continue
		}

		if text, exists := jsonObject["text"]; exists {
			words[text] = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return words
}

func GetRandomWord() string {
	rand.Seed(uint64(time.Now().UnixNano()))

	keys := make([]string, 0, len(Dict))
	for key := range Dict {
		keys = append(keys, key)
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	if len(keys) > 0 {
		return keys[0]
	}
	return ""
}

func IsValidWord(word string) bool{
	return Dict[word]
}
