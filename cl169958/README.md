# CL169958

https://go-review.googlesource.com/c/go/+/169958

```
cmd/compile: better recovery after := (rather than =) in declarations

Before this fix, a mistaken := in a (const/type/var) declaration
ended that declaration with an error from which the parser didn't
recover well. This low-cost change will provide a better error
message and lets the parser recover perfectly.
```

# Demo

[main.go](src/main.go)

```
package main

func main() {
    var _ int := 0
}
```

Run

```
$ ./run.sh
...
go version go1.12.1 linux/amd64
...
# command-line-arguments
/root/main.go:4:15: syntax error: unexpected := at end of statement

go version devel +4145c5da1f Wed Apr 3 23:01:13 2019 +0000 linux/amd64
# command-line-arguments
/root/main.go:4:15: syntax error: unexpected :=, expecting =
```

A better error message will be shown after this commit.
