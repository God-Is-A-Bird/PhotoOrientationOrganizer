package app

import (
	"log"
	"os"
	"sync"
)

func DirWorker(workerID, spawnSubProcesses int, jobs <-chan string) {
	for j := range jobs {
		defer WG.Done()

		f, err := os.Open(j)
		if err != nil {
			log.Print(err)
			continue
		}

		// Get a list of all items located in this dir
		flist, err := f.Readdirnames(-1)
		f.Close()
		if err != nil {
			log.Print(err)
			continue
		}

		var iWG sync.WaitGroup
		imageJobs := make(chan string, len(flist))

		// PER EACH DIR WORKER, create img-workers
		for w := 1; w <= spawnSubProcesses; w++ {
			go ImageWorker(w, imageJobs, &iWG)
		}

		EnumerateImageFiles(j, flist, imageJobs, &iWG)

		iWG.Wait()
		ProgressBar.Increment()
	}
}
