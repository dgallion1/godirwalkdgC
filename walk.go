package godirwalk

import (
	"errors"
	"io/fs"
	"path/filepath"
)

// Options provide parameters for how the Walk function operates.
type Options struct {
	Unsorted bool
	Callback WalkFunc
}

// ErrorAction defines a set of actions the Walk function could take based on
// the occurrence of an error while walking the file system. See the
// documentation for the ErrorCallback field of the Options structure for more
// information.
type ErrorAction int

const (
	// Halt is the ErrorAction return value when the upstream code wants to halt
	// the walk process when a runtime error takes place. It matches the default
	// action the Walk function would take were no ErrorCallback provided.
	Halt ErrorAction = iota

	// SkipNode is the ErrorAction return value when the upstream code wants to
	// ignore the runtime error for the current file system node, skip
	// processing of the node that caused the error, and continue walking the
	// file system hierarchy with the remaining nodes.
	SkipNode
)

// SkipThis is used as a return value from WalkFuncs to indicate that the file
// system entry named in the call is to be skipped. It is not returned as an
// error by any function.
var SkipThis = errors.New("skip this directory entry")

type WalkFunc func(osPathname string, directoryEntry *Dirent) error
type WalkFunc2 func(path string, info fs.FileInfo, err error) error

func Walk(pathname string, options *Options) error {

	visit := func(path string, di fs.DirEntry, err error) error {
		dirent := Dirent{di.Name(), path, di.Type()}
		return options.Callback(path, &dirent)
	}

	return filepath.WalkDir(pathname, visit)
}
