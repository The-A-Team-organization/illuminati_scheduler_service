package controllers

import (
	"illuminati/go/microservice/service"
	"illuminati/go/microservice/utils"
	"net/http"
	"sync"
)

type MutexManager struct{
	sync.Mutex
}

func (m MutexManager) Compromised(w http.ResponseWriter, r *http.Request) {
	
	m.Lock()
	defer m.Unlock()
	newWord := service.GetRandomWord()
	participants, _ := service.GetAppParticipants()
	service.SendWordEmail(newWord,participants)

	hashed, err := utils.HashPassword(newWord)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	data, err :=  utils.SerializePasswordHash(hashed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
  		return
	}

	w.Write(data)

}