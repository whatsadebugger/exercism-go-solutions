package letter

import "sync"

var mapMutex = &sync.Mutex{}

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	wg := sync.WaitGroup{}

	wg.Add(len(strings))

	for _, v := range strings {
		go func(text string) {
			defer wg.Done()
			fm := Frequency(text)
			mapMutex.Lock()
			for k, value := range fm {
				m[k] += value
			}
			mapMutex.Unlock()
		}(v)
	}

	wg.Wait()

	return m
}
