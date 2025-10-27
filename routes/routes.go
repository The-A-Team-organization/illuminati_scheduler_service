package routes

import (
	"illuminati/go/microservice/controllers"
	"net/http"
)

func SetupAPI(m controllers.MutexManager){
	http.HandleFunc("/compromised", m.Compromised)
}

