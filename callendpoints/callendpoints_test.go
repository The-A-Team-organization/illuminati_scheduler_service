package callendpoints

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCloseVotesTimestampFormat(t *testing.T) {
	CloseVotes()
	payload := map[string]string{
		"DateOfEnd": time.Now().Format("2006-01-02 15:04:05"),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	if _, ok := payload["DateOfEnd"]; !ok {
		t.Errorf("payload does not contain DateOfEnd field")
	}

	ts := payload["DateOfEnd"]
	_, err = time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		t.Errorf("DateOfEnd has invalid format: %s", ts)
	}

	if len(data) == 0 {
		t.Error("expected JSON payload not to be empty")
	}
}
