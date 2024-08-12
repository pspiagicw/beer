package programs

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPrograms(t *testing.T) {
	files, err := filepath.Glob("*.b")

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			cmd := exec.Command("../osy", "run", file)

			_, err := cmd.Output()
			if err != nil {
				t.Errorf("Error for %s: %v", file, err)
			}
		})
	}

}
