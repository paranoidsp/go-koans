package go_koans


func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; ; /* infinite loop */ i++ {
		// your code goes here
        if (i >= 100) {
            break
        }

        // fmt.Println("Checking prime: ", i)
        if (isPrimeNumber(i)) {
            // fmt.Println("Found prime: ", i)
            channel <- i;
        }
		assert(i < 100) // i is afraid of heights

	}
}

func aboutConcurrency() {
	ch := make(chan int)

	// your code goes here
    go findPrimeNumbers(ch)

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
