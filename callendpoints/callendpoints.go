package callendpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var (
	BackendURL         = "http://host.docker.internal:8000"
	EndpointVotesClose = "/api/votes/vote_close/"
    EndpointManageInq  = "/api/votes/manage_inquisitor/"
    EndpointBanArchitect = "/api/votes/ban_architect/"


	
	httpNewRequest = http.NewRequest
	httpClientDo   = func(client *http.Client, req *http.Request) (*http.Response, error) {
		return client.Do(req)
	}
)


func CloseVotes() {
	payload := map[string]string{
		"date_of_end": time.Now().Format("2006-01-02 15:04:05"),
	}
	callEndpoint(http.MethodPatch, BackendURL+EndpointVotesClose, payload)
}


func SetInquisitor() {
	callEndpoint(http.MethodPatch, BackendURL+EndpointManageInq, nil)
}

func UnsetInquisitor() {
	callEndpoint(http.MethodDelete, BackendURL+EndpointManageInq, nil)
}

func BanArchitect() {
	callEndpoint(http.MethodDelete, BackendURL+EndpointBanArchitect, nil)
}

func callEndpoint(method, url string, bodyData interface{}) {
	var body *bytes.Buffer
		jsonBody, err := json.Marshal(bodyData)
		if err != nil {
			log.Printf("Error marshalling body: %v", err)
			return
		}
		body = bytes.NewBuffer(jsonBody)

	req, err := httpNewRequest(method, url, body)
	if err != nil {
		log.Printf("Error creating %s request: %v", method, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClientDo(client, req)
	if err != nil {
		log.Printf("Error calling %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Called %s %s - Status: %s", method, url, resp.Status)
}