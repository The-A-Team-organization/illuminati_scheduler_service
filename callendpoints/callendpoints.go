package callendpoints

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
    BackendURL         = "http://backend:8000"
    EndpointVotesClose = "/api/votes/vote_close"
)

func CloseVotes() {
    url := BackendURL + EndpointVotesClose
    
    payload := map[string]string{
        "DateOfEnd": time.Now().Format("2006-01-02 15:04:05"),
    }

    body, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling payload: %v", err)
        return
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
    if err != nil {
        log.Printf("Error calling %s: %v", url, err)
        return
    }
    defer resp.Body.Close()

    log.Printf("Called %s - Status: %s, Timestamp: %s", url, resp.Status, payload["DateOfEnd"])


}
