package app

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ImageWorker(workerID int, jobs <-chan string, iWG *sync.WaitGroup) {
	for j := range jobs {
		defer iWG.Done()

		// Check if file has already been symlinked
		_, err := os.Stat(filepath.Join(PWD, ".poo", "Portrait", strings.TrimPrefix(j, PWD)))
		if err == nil {
			continue
		}
		_, err = os.Stat(filepath.Join(PWD, ".poo", "Landscape", strings.TrimPrefix(j, PWD)))
		if err == nil {
			continue
		}

		f, err := os.Open(j)
		defer f.Close()
		if err != nil {
			log.Print(err)
			return
		}

		// Read the header of the image to get the resolution
		cfg, _, err := image.DecodeConfig(f)
		if err != nil {
			FailedFiles = append(FailedFiles, j)
			continue
		}

		// If files with no extensions are read in, go ahead and add the appropriate ext to the file.
		//if filepath.Ext(j) == "" {
		//	err := os.Rename(j, j+"."+imgext)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}

		// Skip files that don't meet resolution requirements
		if cfg.Width < MinWidth || cfg.Height < MinHeight {
			continue
		}

		if cfg.Height/cfg.Width >= 1 {
			err := os.Symlink(j, filepath.Join(PWD, ".poo", "Portrait", strings.TrimPrefix(j, PWD)))
			if err != nil {
				log.Fatal(err)
			}
		} else {

			err := os.Symlink(j, filepath.Join(PWD, ".poo", "Landscape", strings.TrimPrefix(j, PWD)))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
