package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := bufio.NewWriter(os.Stdout)
	// 標準出力をラップするbufio.Writerを作成
	for i := 0; i < 100; i++ {
		// bufio.Writeに対して書き込みを行う
		fmt.Fprintln(b, strings.Repeat("x", 100))
	}
	b.Flush()
}

// strace -e trace=write -c ./example-bufio > /dev/null
// % time     seconds  usecs/call     calls    errors syscall
// ------ ----------- ----------- --------- --------- ----------------
// 100.00    0.000213          71         3           write
// ------ ----------- ----------- --------- --------- ----------------
// 100.00    0.000213                     3           total
