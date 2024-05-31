package app

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/God-Is-A-Bird/go-PhotoOrientationOrganizer/internal/utils"
	"github.com/cheggaaa/pb/v3"
)

var WG sync.WaitGroup
var ProgressBar *pb.ProgressBar
var PWD string
var FailedFiles []string
var MinWidth int
var MinHeight int

func Run(directory string, dirWorkers, imageWorkers, minWidth, minHeight int) {

	// Create program dir in which symlinks will be placed
	utils.CreatePOODir(directory)

	// Initilize values
	PWD = directory
	MinWidth = minWidth
	MinHeight = minHeight

	// Create a slice which contains all directory paths.
	log.Print("Preparing work...")
	subDirs, err := EnumerateSubdirectories(directory)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create Job Queue
	jobs := make(chan string, len(subDirs))

	// Progress!
	ProgressBar = pb.StartNew(len(subDirs))

	// Create directory workers
	log.Print("Spawning workers...")
	for w := 1; w <= dirWorkers; w++ {
		go DirWorker(w, imageWorkers, jobs)
	}

	log.Print("Staring work! Watch it happen below :)")
	log.Print("Note! Depending on your drive's random read performance, this may take a while!")
	log.Print("Note! This program ignores files without any extension to avoid wasting IOPS.")
	// Add each subdir to the work queue
	for _, j := range subDirs {
		WG.Add(1)
		jobs <- j
	}
	close(jobs)

	// Wait for directory workers to finish job queue
	WG.Wait()

	ProgressBar.Finish()
	log.Print("Done creating symlinks! Find them at: ", filepath.Join(PWD, ".poo"))

	// Prompt user to delete corrupt files
	utils.DeleteCorruptFiles(FailedFiles)

	// Exit Program
	log.Print("Exiting.\n")
	os.Exit(0)
}
