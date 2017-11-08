package main

import (
	"bufio"
	"io"
)

// START OMIT

// Read introduces each line from io.Reader in a channel
func Read(in io.Reader) <-chan string {
	// Create the channel we will write into
	out := make(chan string)

	go func() { // Do this in the background (concurrency!)
		scanner := bufio.NewScanner(in)
		// Read input line by line
		for scanner.Scan() {
			// Write line as text into channel
			out <- scanner.Text()
		}
		close(out) // Finished reading input, close
		// channel. No one else will be able to
		// write into it.
	}()
	return out // Return the channel we are going to fill
}

// END OMIT

// QSTART OMIT

// Process reads elements from channel and processes them
func Process(list <-chan string) { // <-chan means read only
	for l := range list { // Extract element from channel
		// Do stuff with element
	}
	// Channel is now empty (and closed), we are done
}

// QEND OMIT
