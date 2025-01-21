package main_test

import (
	"testing"
)

func TestSystemCanReportErrors(t *testing.T) {
	t.Error("Error logs the message and marks the test as failed, yet continues to execute")
	t.Fatal("Fatal logs the message and marks the test as failed, and it immediately aborts the test")
}
