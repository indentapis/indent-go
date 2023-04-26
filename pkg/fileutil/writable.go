//go:build !windows

package fileutil

import (
	"fmt"
	"io"

	"golang.org/x/sys/unix"
)

// IsWritable returns true if the directory is writable.
func IsWritable(path string) error {
	if err := isDir(path); err != nil {
		return err
	}
	return unix.Access(path, unix.W_OK)
}

func printPermissionsFixCommand(w io.Writer, path string) error {
	_, err := fmt.Fprintf(w, "sudo mkdir -p %s && sudo chown -R $(whoami) %s", path, path)
	return err
}
