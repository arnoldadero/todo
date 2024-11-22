package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// stripAnsi removes ANSI color codes from a string
func stripAnsi(str string) string {
	// This pattern matches all ANSI escape sequences
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(str, "")
}

// captureOutput captures stdout during a test
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		if err != nil && err != io.ErrClosedPipe {
			panic(err)
		}
		outC <- stripAnsi(buf.String())
		close(outC)
	}()

	f()
	w.Close()
	os.Stdout = old

	return <-outC
}

func TestMainIntegration(t *testing.T) {
	// Get absolute path for the test binary
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	binPath := filepath.Join(pwd, "todo-test")

	// Build the test binary
	cmd := exec.Command("go", "build", "-o", binPath, "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build todo binary: %v\nBuild output: %s", err, output)
	}
	defer os.Remove(binPath)

	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "todo-test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set up the todo file path
	todoFile := filepath.Join(tmpDir, "todo.json")

	tests := []struct {
		name    string
		args    []string
		wantOut string
	}{
		{
			name:    "No command",
			args:    []string{},
			wantOut: "Usage:",
		},
		{
			name:    "Add todo",
			args:    []string{"-add", "Test todo item"},
			wantOut: "Added todo: Test todo item",
		},
		{
			name:    "List todos - one item",
			args:    []string{"-list"},
			wantOut: "Test todo item",
		},
		{
			name:    "Complete todo",
			args:    []string{"-complete", "1"},
			wantOut: "Completed todo #1",
		},
		{
			name:    "List todos - completed item",
			args:    []string{"-list"},
			wantOut: "Test todo item",
		},
		{
			name:    "Add another todo",
			args:    []string{"-add", "Another test item"},
			wantOut: "Added todo: Another test item",
		},
		{
			name:    "Delete todo",
			args:    []string{"-del", "1"},
			wantOut: "Deleted todo #1",
		},
		{
			name:    "List todos - after deletion",
			args:    []string{"-list"},
			wantOut: "Another test item",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binPath, tt.args...)
			cmd.Env = append(os.Environ(), "TODO_FILE="+todoFile)
			output, err := cmd.CombinedOutput()
			outputStr := stripAnsi(string(output))

			if err != nil && !strings.Contains(outputStr, tt.wantOut) {
				t.Errorf("Command failed: %v\nOutput: %s", err, output)
			}

			if !strings.Contains(outputStr, tt.wantOut) {
				t.Errorf("Expected output containing %q, got %q", tt.wantOut, outputStr)
			}
		})
	}
}
