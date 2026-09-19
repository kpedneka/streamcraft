package telemetry

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewLogger_JSONOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "test-service")

	logger.Info("test message", "key1", "value1", "key2", 42)

	var record map[string]any
	if err := json.NewDecoder(&buf).Decode(&record); err != nil {
		t.Fatalf("output is not valid JSON: %v — output: %q", err, buf.String())
	}

	if record["service"] != "test-service" {
		t.Errorf("service field: got %q, want %q", record["service"], "test-service")
	}

	if record["key1"] != "value1" {
		t.Errorf("key1 field: got %q, want %q", record["key1"], "value1")
	}

	if record["msg"] != "test message" {
		t.Errorf("msg field: got %q, want %q", record["msg"], "test message")
	}

}
