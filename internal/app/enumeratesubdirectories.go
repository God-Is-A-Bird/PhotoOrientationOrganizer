package app

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/God-Is-A-Bird/go-PhotoOrientationOrganizer/internal/utils"
)

func EnumerateSubdirectories(dir string) ([]string, error) {
	var subDirs []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Print(err)
			return err
		}

		if d.IsDir() {
			// Check if inside .poo dir
			if filepath.Base(path) == ".poo" {
				return filepath.SkipDir
			}

			// Create symlink dir mirror
			err := os.MkdirAll(filepath.Join(PWD, ".poo", "Portrait", strings.TrimPrefix(path, PWD)), 0700)
			if err != nil {
				log.Fatal(err)
			}
			err = os.MkdirAll(filepath.Join(PWD, ".poo", "Landscape", strings.TrimPrefix(path, PWD)), 0700)
			if err != nil {
				log.Fatal(err)
			}

			// Check if directory is empty, is so, don't add it to the slice
			isEmpty, err := utils.IsDirEmpty(path)
			if err != nil {
				log.Print(err)
				return err
			}

			if !isEmpty {
				subDirs = append(subDirs, path)
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return subDirs, nil
}
