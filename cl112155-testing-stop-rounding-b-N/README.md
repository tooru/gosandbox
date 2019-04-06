# CL112155

https://go-review.googlesource.com/c/go/+/112155

```
testing: stop rounding b.N

The original goal of rounding to readable b.N
was to make it easier to eyeball times.
However, proper analysis requires tooling
(such as benchstat) anyway.

Instead, take b.N as it comes.
This will reduce the impact of external noise
such as GC on benchmarks.

This requires reworking our iteration estimates.
We used to calculate the estimated ns/op
and then divide our target ns by that estimate.
However, this order of operations was destructive
when the ns/op was very small; rounding could
hide almost an order of magnitude of variation.
Instead, multiply first, then divide.
Also, make n an int64 to avoid overflow.

Prior to this change, we attempted to cap b.N at 1e9.
Due to rounding up, it was possible to get b.N as high as 2e9.
This change consistently enforces the 1e9 cap.

This change also reduces the wall time required to run benchmarks.

Here's the impact of this change on the wall time to run
all benchmarks once with benchtime=1s on some std packages:

name           old time/op       new time/op       delta
bytes                 306s ± 1%         238s ± 1%  -22.24%  (p=0.000 n=10+10)
encoding/json         112s ± 8%          99s ± 7%  -11.64%  (p=0.000 n=10+10)
net/http             54.7s ± 7%        44.9s ± 4%  -17.94%  (p=0.000 n=10+9)
runtime               957s ± 1%         714s ± 0%  -25.38%  (p=0.000 n=10+9)
strings               262s ± 1%         201s ± 1%  -23.27%  (p=0.000 n=10+10)
[Geo mean]            216s              172s       -20.23%
```

# Demo

Run

```
$ ./run.sh
...
go version go1.12.1 linux/amd64
...
goos: linux
goarch: amd64
BenchmarkFoo-2   	50000000	        34.3 ns/op
PASS
ok  	command-line-arguments	1.774s

go version devel +03a79e94ac Wed Mar 20 21:19:16 2019 +0000 linux/amd64
goos: linux
goarch: amd64
BenchmarkFoo-2   	31800799	        33.0 ns/op
PASS
ok  	command-line-arguments	1.135s
```
