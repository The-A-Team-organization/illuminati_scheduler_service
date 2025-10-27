package scheduler

import (
	"bytes"
	"fmt"
	"illuminati/go/microservice/service"
	"illuminati/go/microservice/utils"
	"net/http"
)
func newPassword(){

	newWord := service.GetRandomWord()
	participants, _ := service.GetAppParticipants()
	service.SendWordEmail(newWord,participants)
	hashed, err := utils.HashPassword(newWord)
	if err != nil {
		fmt.Println(err)
	}
	data, err :=  utils.SerializePasswordHash(hashed)
	if err != nil {
		
	}
	http.Post("", "application/json", bytes.NewBuffer(data))

}