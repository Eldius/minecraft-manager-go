package versions

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestParseVersionsFile(t *testing.T) {
	path, err := filepath.Abs(".")
	if err != nil {
		t.Errorf("Failed to parse current location")
	}
	t.Logf("path: %s", path)
	f, err := os.Open(filepath.Join(path, "samples", "version_manifest.json"))
	if err != nil {
		t.Errorf("Failed to open sample file\n%s", err.Error())
	}
	parsed, err := parseVersionsFile(f)
	if err != nil {
		t.Errorf("Failed to parse sample file\n%s", err.Error())
	}
	fmt.Println("versions: ", len(parsed.Versions))
}