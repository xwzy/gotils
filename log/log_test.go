package log

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestLogFunctions(t *testing.T) {
	// Temporarily replace the log output
	var buf bytes.Buffer
	originalOutput := log.Out
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(originalOutput)
	}()

	// Test Info
	Info("test info")
	assertLogOutput(t, &buf, "info", "test info")

	// Test Infof
	Infof("test %s", "infof")
	assertLogOutput(t, &buf, "info", "test infof")

	// Test Error
	Error("test error")
	assertLogOutput(t, &buf, "error", "test error")

	// Test Errorf
	Errorf("test %s", "errorf")
	assertLogOutput(t, &buf, "error", "test errorf")

	// Test Warn
	Warn("test warn")
	assertLogOutput(t, &buf, "warning", "test warn")

	// Test Warnf
	Warnf("test %s", "warnf")
	assertLogOutput(t, &buf, "warning", "test warnf")

	// Test Debug (this won't output anything as default level is Info)
	Debug("test debug")
	assert.Empty(t, buf.String())

	// Test Debugf (this won't output anything as default level is Info)
	Debugf("test %s", "debugf")
	assert.Empty(t, buf.String())
}

func assertLogOutput(t *testing.T, buf *bytes.Buffer, expectedLevel, expectedMessage string) {
	var logEntry map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &logEntry)
	assert.NoError(t, err)

	assert.Equal(t, expectedLevel, logEntry["level"])
	assert.Equal(t, expectedMessage, logEntry["msg"])

	buf.Reset()
}

func TestLogToFile(t *testing.T) {
	// Create a temporary directory for the log file
	tempDir, err := os.MkdirTemp("", "log_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Set up a temporary log file
	logFilePath := filepath.Join(tempDir, "test.log")

	// Configure lumberjack for the test
	testRotateFileHook := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   false,
	}

	// Set logrus to use the test file
	originalOutput := log.Out
	log.SetOutput(testRotateFileHook)
	defer func() {
		log.SetOutput(originalOutput)
	}()

	// Write some log entries
	Info("This is a test info log")
	Error("This is a test error log")
	Warn("This is a test warning log")

	// Verify that the file was created and contains content
	fileInfo, err := os.Stat(logFilePath)
	assert.NoError(t, err)
	assert.True(t, fileInfo.Size() > 0, "Log file should not be empty")

	// Read the file content (optional, for debugging)
	content, err := os.ReadFile(logFilePath)
	assert.NoError(t, err)
	t.Logf("Log file content: %s", string(content))
}
