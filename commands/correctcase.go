package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func correct(path string) (string, string, error) {
	dirname, fname := filepath.Split(filepath.Clean(path))
	fd, fdErr := os.Open(dirname)
	if fdErr != nil {
		return dirname, fname, fdErr
	}
	defer fd.Close()
	fi, fiErr := fd.Readdir(-1)
	if fiErr != nil {
		return dirname, fname, fiErr
	}
	for _, fi1 := range fi {
		if strings.EqualFold(fi1.Name(), fname) {
			return dirname, fi1.Name(), nil
		}
	}
	return dirname, fname, fmt.Errorf("%s: not found.", path)
}

// correct path's case.
func CorrectCase(path string) (string, error) {
	dirname, fname, err := correct(path)
	if err != nil {
		return path, err
	}
	if len(dirname) > 0 {
		// NOT root directory.
		dirname, err = CorrectCase(dirname)
	}
	return filepath.Join(dirname, fname), nil
}
