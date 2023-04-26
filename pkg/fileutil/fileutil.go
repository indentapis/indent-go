// Package fileutil provides utility functions for working with files.
package fileutil

import (
	"fmt"
	"io"
	"os"
)

// PrintPermissionsFix prints a message to the user on how to fix file permissions.
func PrintPermissionsFix(w io.Writer, path string) (err error) {
	if _, err = fmt.Fprintf(w, "Permissions currently prevent writing to %q.\n\n", path); err != nil {
		return
	} else if _, err = fmt.Fprintln(w, "To fix this, run the following commands:"); err != nil {
		return
	}
	return printPermissionsFixCommand(w, path)
}

// isDir checks if path is a directory.
func isDir(path string) error {
	if info, err := os.Stat(path); err != nil {
		return err
	} else if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", path)
	}
	return nil
}
