package loadtest

import (
	"testing"
)

func TestRunLoadTest(t *testing.T) {
	url := "http://example.com"
	totalRequests := 10
	concurrency := 2

	report := RunLoadTest(url, totalRequests, concurrency)

	if report.TotalRequests != totalRequests {
		t.Errorf("Expected %d total requests, got %d", totalRequests, report.TotalRequests)
	}
}
