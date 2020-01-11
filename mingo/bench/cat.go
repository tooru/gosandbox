// cat.go
package cat

import (
	"bytes"
	"strings"
)

// catは += 演算子を使って文字列を結合する
func cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// bufは bytes.Buffer を使って文字列を結合する
func buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		// NOTICE: エラーは無視している
		b.WriteString(s)
	}
	return b.String()
}

// sbは strings.Builder を使って文字列を結合する
func sb(ss ...string) string {
	var b strings.Builder
	for _, s := range ss {
		// NOTICE: エラーは無視している
		b.WriteString(s)
	}
	return b.String()
}

// BenchmarkCat3-4       	 5054496	       228 ns/op	      54 B/op	       3 allocs/op
// BenchmarkBuf3-4       	 5868398	       207 ns/op	     115 B/op	       3 allocs/op
// BenchmarkSb3-4        	 7277700	       155 ns/op	      56 B/op	       2 allocs/op
// BenchmarkCat100-4     	 1985695	       607 ns/op	     224 B/op	      10 allocs/op
// BenchmarkBuf100-4     	 3632821	       333 ns/op	     240 B/op	       3 allocs/op
// BenchmarkSb100-4      	 3736467	       304 ns/op	     184 B/op	       3 allocs/op
// BenchmarkCat10000-4   	      82	  14303326 ns/op	53327813 B/op	   10000 allocs/op
// BenchmarkBuf10000-4   	    6373	    201854 ns/op	  211376 B/op	      11 allocs/op
// BenchmarkSb10000-4    	    7282	    148729 ns/op	  212344 B/op	      18 allocs/op
