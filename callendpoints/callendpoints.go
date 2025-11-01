package callendpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var (
    BackendURL         = "http://backend:8000"
    EndpointVotesClose = "/api/votes/vote_close/"
    jsonMarshal    = json.Marshal
	httpNewRequest = http.NewRequest
)

func CloseVotes() {
    url := BackendURL + EndpointVotesClose
    
    payload := map[string]string{
        "date_of_end": time.Now().Format("2006-01-02 15:04:05"),
    }

    body, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling payload: %v", err)
        return
    }

   req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")


	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

    if err != nil {
        log.Printf("Error calling %s: %v", url, err)
        return
    }
    defer resp.Body.Close()

    log.Printf("Called %s - Status: %s, Timestamp: %s", url, resp.Status, payload["date_of_end"])


}
