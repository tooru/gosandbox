package main

import "sync/atomic"

func main() {
    var a int64 = 1

    atomic.StoreInt64(&a, 2)
}
