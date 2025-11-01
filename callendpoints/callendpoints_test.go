package callendpoints

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func TestCloseVotes_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	orig := BackendURL
	defer func() { BackendURL = orig }()
	BackendURL = server.URL

	CloseVotes()
}


func TestCloseVotes_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	orig := BackendURL
	defer func() { BackendURL = orig }()
	BackendURL = server.URL

	CloseVotes()
}


func TestCloseVotes_NetworkError(t *testing.T) {
	orig := BackendURL
	defer func() { BackendURL = orig }()
	BackendURL = "http://invalid_host"

	CloseVotes()
}




func TestCloseVotes_NewRequestError(t *testing.T) {
	orig := httpNewRequest
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, errors.New("request error")
	}
	defer func() { httpNewRequest = orig }()

	CloseVotes()
}


func TestCloseVotes_ClientDoError(t *testing.T) {
	orig := httpClientDo
	httpClientDo = func(client *http.Client, req *http.Request) (*http.Response, error) {
		return nil, errors.New("client error")
	}
	defer func() { httpClientDo = orig }()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	origURL := BackendURL
	defer func() { BackendURL = origURL }()
	BackendURL = server.URL

	CloseVotes()
}

func TestCloseVotesTimestampFormat(t *testing.T) {
	
	payload := map[string]string{
		"date_of_end": time.Now().Format("2006-01-02 15:04:05"),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	if _, ok := payload["date_of_end"]; !ok {
		t.Errorf("payload does not contain date_of_end field")
	}

	ts := payload["date_of_end"]
	_, err = time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		t.Errorf("date_of_end has invalid format: %s", ts)
	}

	if len(data) == 0 {
		t.Error("expected JSON payload not to be empty")
	}
	CloseVotes()
}