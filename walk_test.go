package godirwalkdgC

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestWalk(t *testing.T) {
	// Test the Walk function by walking the directory hierarchy rooted at the
	// current working directory.
	var paths []string
	err := Walk("..", &Options{
		Callback: func(osPathname string, de *Dirent) error {
			paths = append(paths, osPathname)
			return nil
		},
	})
	if err != nil {
		t.Fatalf("Walk() failed: %s", err)
	}

	// Sort the paths slice so that we can compare it to the output of the
	// filepath.Walk function.
	sort.Strings(paths)

	// Now walk the same directory hierarchy using the filepath.Walk function.
	var paths2 []string
	err = filepath.Walk("..", func(osPathname string, info os.FileInfo, err error) error {
		paths2 = append(paths2, osPathname)
		return nil
	})
	if err != nil {
		t.Fatalf("filepath.Walk() failed: %s", err)
	}

	// Sort the paths2 slice so that we can compare it to the output of the
	// godirwalk.Walk function.
	sort.Strings(paths2)

	// Compare the two slices.
	if len(paths) != len(paths2) {
		t.Fatalf("godirwalk.Walk() returned %d paths, filepath.Walk() returned %d paths", len(paths), len(paths2))
	}
	for i := 0; i < len(paths); i++ {
		if paths[i] != paths2[i] {
			t.Fatalf("godirwalk.Walk() returned %s, filepath.Walk() returned %s", paths[i], paths2[i])
		}
	}
}

