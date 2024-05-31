package app

import (
	"path/filepath"
	"sync"

	"github.com/God-Is-A-Bird/go-PhotoOrientationOrganizer/internal/utils"
)

func EnumerateImageFiles(dirpath string, paths []string, jobs chan<- string, WG *sync.WaitGroup) {

	for _, file := range paths {
		isDir := utils.IsDirectory(filepath.Join(dirpath, file))
		if isDir {
			continue
		}

		validExt := utils.IsValidImageExtension(file)
		if !validExt {
			continue
		}

		WG.Add(1)
		jobs <- filepath.Join(dirpath, file)
	}
	close(jobs)

}
