package go_koans


func aboutChannels() {
	ch := make(chan string, 2)

	assert(len(ch) == 0) // channels are like buffers

	ch <- "foo" // i mean, "metaphors are like similes"

	assert(len(ch) == 1) // they can be queried for queued items

	assert(<-ch == "foo") // items can be popped out of them

	assert(len(ch) == 0) // and len() always reflects the "current" queue status

	// the 'go' keyword runs a function-call in a new "goroutine"
	// which executes "concurrently" with the calling "goroutine"
	go func() {
		// your code goes here
        // fmt.Println("\n Initial length: ", len(ch))
        for {
            if len(ch) == 2 {
                _ = <-ch
            }
        }
	}()

	assert(len(ch) == 0) // we'll need to make room for the queue, or suffer deadlocks

    // fmt.Println("\n Initial length outside: ", len(ch))
	ch <- "bar"   // this send will succeed
	ch <- "quux"  // there's enough room for this send too
	ch <- "extra" // but the buffer only has two slots
}
