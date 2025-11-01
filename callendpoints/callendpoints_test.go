package callendpoints

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCloseVotes_SendsPatchRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH method, got %s", r.Method)
		}

		if !strings.HasSuffix(r.URL.Path, EndpointVotesClose) {
			t.Errorf("unexpected endpoint: %s", r.URL.Path)
		}

		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("expected Content-Type: application/json, got %s", ct)
		}

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var payload map[string]string
		err := json.Unmarshal(body, &payload)
		if err != nil {
			t.Errorf("invalid JSON payload: %v", err)
		}

		ts, ok := payload["date_of_end"]
		if !ok {
			t.Errorf("missing date_of_end field in payload")
		} else if _, err := time.Parse("2006-01-02 15:04:05", ts); err != nil {
			t.Errorf("invalid timestamp format: %s", ts)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	origBackendURL := BackendURL
	defer func() { BackendURL = origBackendURL }()
	BackendURL = server.URL

	CloseVotes()
}

func TestCloseVotes_JSONMarshalError(t *testing.T) {
	origMarshal := jsonMarshal
	jsonMarshal = func(v interface{}) ([]byte, error) {
		return nil, errors.New("marshal error")
	}
	defer func() { jsonMarshal = origMarshal }()

	CloseVotes()
}

func TestCloseVotes_NewRequestError(t *testing.T) {
	origRequest := httpNewRequest
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, errors.New("request error")
	}
	defer func() { httpNewRequest = origRequest }()

	CloseVotes()
}
func TestCloseVotesTimestampFormat(t *testing.T) {
	CloseVotes()
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
}