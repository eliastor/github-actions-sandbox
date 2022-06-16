package unit1

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise4_Syntax(t *testing.T) {
}

func TestExercise4_Output(t *testing.T) {
	codepath := "../../unit1/exercises/e4/main.go"

	cmd := exec.Command("timeout", "10", "go", "run", codepath)
	tmpDir, err := os.MkdirTemp("", "gotest-*")
	if !assert.NoError(t, err) {
		assert.FailNow(t, "test error: cannot create temporary dir")
	}
	defer os.RemoveAll(tmpDir)

	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),
		"CGO_ENABLED=0",
		"GOCACHE=" + tmpDir,
	}

	outBuf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	err = cmd.Run()
	if !assert.NoError(t, err, errBuf.String()) {
		assert.FailNow(t, "test error: \"go run\" failed", err)
		return
	}

	assert.NoError(t, err, "Can't compile and run code")
	assert.Empty(t, errBuf.String(), "stderr of your code must be empty", errBuf.String())
	assert.NotEmpty(t, outBuf.String(), "stdout of your code must not be empty")

	assert.Equal(t, "You cracked the code!", outBuf.String())
}
