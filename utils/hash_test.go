package utils

import "testing"


func Test_HashPassword(t *testing.T){
	hashed, _ := HashPassword("newWord")
	
	if hashed == "" {
		t.Error("bcrypt doesnt work!")
	}
}

func Test_SerializePasswordHash(t *testing.T){
	hashed, _ := SerializePasswordHash("newWord")
	
	if hashed == nil {
		t.Error("bcrypt doesnt work!")
	}
}