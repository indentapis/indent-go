package fileutil

import (
	"fmt"
	"io"
	"os"
)

// IsWritable returns true if the directory is writable.
func IsWritable(path string) error {
	if err := isDir(path); err != nil {
		return err
	} else if info, err := os.Stat(path); err != nil {
		return err
	} else if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return fmt.Errorf("directory %q is not writable, permissions are %s", path, info.Mode().Perm())
	}
	return nil
}

func printPermissionsFixCommand(w io.Writer, path string) error {
	if _, err := fmt.Fprintf(w, "\tmkdir %q\n", path); err != nil {
		return err
	} else if _, err = fmt.Fprintf(w, "\ttakeown /R /F %q\n", path); err != nil {
		return err
	}
	return nil
}
