package proxy

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestExtract(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("can't create temp directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	file, err := os.Open("./test/sample.zip")
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	client := &GithubClientImpl{}
	err = client.extract(file.Name(), tempDir)
	if err != nil {
		t.Fatalf("failed extract: %s", err)
	}

	_, err = os.Stat(filepath.Join(tempDir, "test"))
	if err != nil {
		t.Fatalf("file not found: %s", err)
	}
}
