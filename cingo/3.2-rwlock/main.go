package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) { // ❶
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) // ❷
		}
	}
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}
	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()
	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m))
	}
}

// Readers  RWMutext      Mutex
// 1        77.927µs      40.027µs
// 2        19.859µs      8.69µs
// 4        15.443µs      37.677µs
// 8        45.243µs      53.353µs
// 16       22.826µs      58.196µs
// 32       93.805µs      87.503µs
// 64       210.266µs     66.012µs
// 128      145.29µs      225.822µs <-
// 256      469.535µs     96.393µs
// 512      162.74µs      188.363µs
// 1024     354.267µs     381.72µs
// 2048     700.164µs     770.813µs
// 4096     1.382192ms    1.51859ms
// 8192     3.123486ms    3.435976ms
// 16384    6.67749ms     6.930911ms
// 32768    12.84362ms    11.930864ms <-
// 65536    23.837728ms   33.30645ms
// 131072   53.776079ms   50.342382ms <-
// 262144   251.275632ms  145.293453ms
// 524288   223.945188ms  422.895467ms
