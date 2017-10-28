package utilities

import (
	"log"
	"os"
	"path/filepath"
)

// ValidDir will check if a directory is valid and return a bool.
func ValidDir(dir string) (valid bool) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("directory does not exist")
		return
	}

	new := filepath.Join(dir, "tempDir")

	if err := os.Mkdir(new, 0755); err != nil {
		log.Println("unable to create directory")
		return
	}

	if err := os.Remove(new); err != nil {
		log.Println("unable to remove directory")
		return
	}

	return true

}

// // CreateDir will check if base is writable and create dir if it is writable.
// func CreateDir(base string, dir string) (err error) {

// 	return nil
// }
