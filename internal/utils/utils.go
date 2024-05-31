package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func GetPWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func VerifyWD(wd string) {
	_, err := os.Stat(wd)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func VerifyDirWorkers(workers int) {
	if workers < 1 {
		log.Fatal("Can't have less than one directory worker buddy!")
	}

	return
}

func VerifyImageWorkers(workers int) {
	if workers < 1 {
		log.Fatal("Can't have less than one image worker buddy!")
	}

	return
}

func VerifyMaxWorkers(dirWorkers, imgWorkers int) {
	if (dirWorkers * imgWorkers) > 10_000 {
		log.Fatal(`
/*
	GO spawns new threads for GO routines if:
		"A Go program creates a new thread only when a goroutine is ready to run but all the existing threads are blocked
		in system calls, cgo calls, or are locked to other goroutines due to use of runtime.LockOSThread.""
	By default, the MaxThreads allowed is 10,000. Since this program is IO Bound, it is almost gaurnteed that attempting
	to spawn more than 10,000 workers in total will result in a crash. Thus, dir-workers * img-workers should remain below 10,000.
*/`)
	}
}

func CreatePOODir(path string) {
	os.Mkdir(filepath.Join(path, ".poo"), 0700)
	os.Mkdir(filepath.Join(path, ".poo", "Portrait"), 0700)
	os.Mkdir(filepath.Join(path, ".poo", "Landscape"), 0700)
}

func IsDirEmpty(dir string) (bool, error) {

	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in one file
	_, err = f.Readdir(1)

	// if the file is EOF, the dir is empty
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Print(err)
		return false
	}

	return fileInfo.IsDir()
}

// Checks to see if image is JPG, PNG, GIF. Returns false if no extension is found.
func IsValidImageExtension(path string) bool {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	extension := strings.ToLower(filepath.Ext(path))
	if slices.Contains(allowedExtensions, extension) {
		return true
	}

	return false
}

// Ptompt user if they would like to delete files which are likely corrupt
func DeleteCorruptFiles(paths []string) {
	if len(paths) == 0 {
		return
	}

	for _, path := range paths {
		fmt.Println("Error decoding file, likely corrupt: ", path)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Should we help you by deleting these likely corrupt files? (y/n)")
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if fmt.Sprintf("%c", char) == "Y" || fmt.Sprintf("%c", char) == "y" {
			fmt.Println("Deleting...")
			for _, path := range paths {
				err := os.Remove(path)
				if err != nil {
					log.Println(err)
				}
			}
			return
		} else if fmt.Sprintf("%c", char) == "N" || fmt.Sprintf("%c", char) == "n" {
			return
		}

	}

}
