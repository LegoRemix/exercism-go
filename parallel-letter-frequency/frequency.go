// Package letter helps compute the frequency of letters in a given text
package letter

import "runtime"

var procs = runtime.GOMAXPROCS(0)

// FreqMap is a mapping from characters to counts
type FreqMap map[rune]int

// Frequency computes count of letters in each string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// worker computes the frequency for a chunk specified to it
func worker(s string, outChan chan<- FreqMap) {
	outChan <- Frequency(s)
}

// ConcurrentFrequency attempts to compute the frequency of letters using parallelism
func ConcurrentFrequency(texts []string) FreqMap {
	pieces := procs / len(texts)
	if pieces == 0 {
		pieces = 1
	}

	//the number of workers from which we have to get answers
	var workerCount int

	mapChan := make(chan FreqMap)

	//for each text
	for _, txt := range texts {
		chunkSize := len(txt) / pieces
		for pos := 0; pos < len(txt); pos += chunkSize {
			end := pos + chunkSize
			if end > len(txt) {
				end = len(txt)
			}

			workerCount++
			go worker(txt[pos:end], mapChan)
		}
	}

	result := make(map[rune]int)
	for i := 0; i < workerCount; i++ {
		partial := <-mapChan
		for k, v := range partial {
			result[k] += v
		}
	}

	return result

}
