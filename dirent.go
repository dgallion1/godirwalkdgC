package godirwalkdgC

import (
	"os"
)

// Dirent stores the name and file system mode type of discovered file system
// entries.
type Dirent struct {
	name     string      // base name of the file system entry.
	path     string      // path name of the file system entry.
	modeType os.FileMode // modeType is the type of file system entry.
}
