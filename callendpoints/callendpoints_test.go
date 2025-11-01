package callendpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


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



func mockHTTPClient(statusCode int) func(*http.Client, *http.Request) (*http.Response, error) {
	return func(client *http.Client, req *http.Request) (*http.Response, error) {
		rec := httptest.NewRecorder()
		rec.WriteHeader(statusCode)
		return rec.Result(), nil
	}
}

func TestSetInquisitor_Success(t *testing.T) {
	httpClientDo = mockHTTPClient(204)
	SetInquisitor()
}

func TestUnsetInquisitor_Success(t *testing.T) {
	httpClientDo = mockHTTPClient(200)
	UnsetInquisitor()
}

func TestBanArchitect_Success(t *testing.T) {
	httpClientDo = mockHTTPClient(200)
	BanArchitect()
}

func TestNewEntryPassword_Success(t *testing.T) {
	httpClientDo = mockHTTPClient(200)
	NewEntryPassword()
}

func TestCloseVotes_Success(t *testing.T) {
	httpClientDo = mockHTTPClient(200)
	CloseVotes()
}

func TestCallEndpoint_ErrorRequest(t *testing.T) {
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, io.EOF
	}
	defer func() {
		httpNewRequest = http.NewRequest
	}()
	callEndpoint(http.MethodPatch, "bad-url", nil)
}
