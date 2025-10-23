package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
//
// We use the fork/join pattern.
// Here's one execution of `go test -bench .`
//
//	BenchmarkSequentialFrequency-10    	   16345	     72857 ns/op
//	BenchmarkConcurrentFrequency-10    	   20041	     61049 ns/op
//
// We can see that the concurrent version shows a lower number of
// nano seconds per operation (ns/op) than the sequential version.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	freqs := make(chan FreqMap)
	total := make(chan FreqMap)

	// Start joiner first
	go func() {
		fm := FreqMap{}
		for f := range freqs { // read until `freqs` closed
			for k, v := range f {
				fm[k] += v
			}
		}
		total <- fm
	}()

	// Fork workers
	for _, text := range texts {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			freqs <- Frequency(t)
		}(text)
	}

	// Wait for all workers, then close the channel
	wg.Wait()
	close(freqs)

	// Collect the total result
	return <-total
}
