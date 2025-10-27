package service

import (
	"strings"
	"testing"
)

func TestGetRandomWord(t *testing.T) {
    word := GetRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}