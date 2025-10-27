package utils

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func SerializePasswordHash(password string) ([]byte, error){
	payload := map[string]string{
		"password": string(password),
	}
	data, err := json.Marshal(payload)
	return data, err
}