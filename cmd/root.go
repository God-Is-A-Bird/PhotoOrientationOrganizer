package cmd

import (
	"flag"

	"github.com/God-Is-A-Bird/go-PhotoOrientationOrganizer/internal/app"
	"github.com/God-Is-A-Bird/go-PhotoOrientationOrganizer/internal/utils"
)

func Execute() {
	// Define flags
	dirPtr := flag.String("dir", utils.GetPWD(), "Directory you would like the program to search for images in and create .poo dir with symlinks.")
	dirWorkersPtr := flag.Int("dir-workers", 2, "How many directories would you like the program to process at once? Default: 2")
	imageWorkersPtr := flag.Int("img-workers", 50, "Per each directory worker, how many images should the program process at once?. Default: 50")
	minWidthPtr := flag.Int("min-width", 0, "Images with less than this will not be symlinked")
	minHeightPtr := flag.Int("min-height", 0, "Images with less than this will not be symlinked")

	// Parse flags
	flag.Parse()

	// Verify flag values are valid
	utils.VerifyWD(*dirPtr)
	utils.VerifyDirWorkers(*dirWorkersPtr)
	utils.VerifyImageWorkers(*imageWorkersPtr)

	/*
		GO spawns new threads for GO routines if:
			"A Go program creates a new thread only when a goroutine is ready to run but all the existing threads are blocked
			in system calls, cgo calls, or are locked to other goroutines due to use of runtime.LockOSThread.""
		By default, the MaxThreads allowed is 10,000. Since this program is IO Bound, it is almost gaurnteed that attempting
		to spawn more than 10,000 workers in total will result in a crash. Thus, dir-workers * img-workers should remain below 10,000.
	*/
	utils.VerifyMaxWorkers(*dirWorkersPtr, *imageWorkersPtr)

	// Run Program
	app.Run(*dirPtr, *dirWorkersPtr, *imageWorkersPtr, *minWidthPtr, *minHeightPtr)
}
