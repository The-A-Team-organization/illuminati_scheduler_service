package mailer

import (
	"strings"
	"testing"
)

func TestGetRandomWord(t *testing.T) {
    word := getRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}