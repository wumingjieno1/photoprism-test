package fs

import (
	"path/filepath"
	"strconv"
	"strings"
)

// StripSequence removes common sequence patterns at the end of file names.
func StripSequence(name string) string {
	// Strip numeric extensions like .00000, .00001, .4542353245,.... (at least 5 digits).
	if dot := strings.LastIndex(name, "."); dot != -1 && len(name[dot+1:]) >= 5 {
		if i, err := strconv.Atoi(name[dot+1:]); err == nil && i >= 0 {
			name = name[:dot]
		}
	}

	// Other common sequential naming schemes.
	if end := strings.Index(name, "("); end != -1 {
		// Copies created by Chrome & Windows, example: IMG_1234 (2).
		name = name[:end]
	} else if end := strings.Index(name, " copy"); end != -1 {
		// Copies created by OS X, example: IMG_1234 copy 2.
		name = name[:end]
	}

	name = strings.TrimSpace(name)

	return name
}

// BasePrefix returns the filename base without any extensions and path.
func BasePrefix(fileName string, stripSequence bool) string {
	name := StripKnownExt(StripExt(filepath.Base(fileName)))

	if !stripSequence {
		return name
	}

	return StripSequence(name)
}

// RelPrefix returns the relative filename.
func RelPrefix(fileName, dir string, stripSequence bool) string {
	if name := RelName(fileName, dir); name != "" {
		return AbsPrefix(name, stripSequence)
	}

	return BasePrefix(fileName, stripSequence)
}

// AbsPrefix returns the directory and base filename without any extensions.
func AbsPrefix(fileName string, stripSequence bool) string {
	return filepath.Join(filepath.Dir(fileName), BasePrefix(fileName, stripSequence))
}
