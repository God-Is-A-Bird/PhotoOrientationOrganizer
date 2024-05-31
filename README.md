Are you a programmer of culture with one portrait monitor for Spotify/Discord/Docs and one landscape monitor for all your other needs? Do you have 100,000+ wallpapers that you wish you could have organized by orientation so that they fit your monitors appropriately? Use POO (Photo Orientation Organizer)!

This program creates a `.poo` directory in the PWD mirroring the PWD's directory structure and places symlinks of all images into `.poo/Portrait` or `.poo/Landscape`.


How To Install (requires you to already have Go installed on your system):

`go install github.com/God-Is-A-Bird/PhotoOrientationOrganizer@latest`


Usage of `PhotoOrientationOrganizer`:
```
  -dir string
    	Directory you would like the program to search for images in and create .poo dir with symlinks.
  -dir-workers int
    	How many directories would you like the program to process at once? Default: 2 (default 2)
  -img-workers int
    	Per each directory worker, how many images should the program process at once?. Default: 50 (default 50)
  -min-height int
    	Images with less than this will not be symlinked
  -min-width int
    	Images with less than this will not be symlinked
```

Depending on your file structure, you may find it beneficial to play around with `-dir-workers` and `-img-workers`. These control the number of GO Routines spawned. Remember, you're likely limited by your drive's read speeds, not CPU, so choose values that make sense. On Linux, you can use `iostat [drive] [timeduration]` (ex. `iostat sda 5` to print every five seconds) to monitor your drive's `%iowait` and `%idle`. If `%iowait` is nearly maxed out, spawning more workers won't help and may actually hurt performance. 

![image](https://github.com/God-Is-A-Bird/PhotoOrientationOrganizer/assets/27874321/8ea798a7-d1b1-4111-9fdc-0f99f2c1d77a)
