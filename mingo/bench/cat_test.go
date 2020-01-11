// cat_test.go
package cat

import (
	"testing"
)

// seedはベンチマーク用のトークンをつくる
// 長さを受け取り、指定された長さの文字列のスライスを生成する // 今回は、単純に "a" を n 個ならべたスライスを生成する
func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

// benchはベンチマーク用のヘルパ
// テストしたい文字列の組み合わせ長と、文字列結合のための
// 手続きを渡す。それについてベンチマークを実行させる
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}
func BenchmarkCat3(b *testing.B)     { bench(b, 3, cat) }
func BenchmarkBuf3(b *testing.B)     { bench(b, 3, buf) }
func BenchmarkCat100(b *testing.B)   { bench(b, 10, cat) }
func BenchmarkBuf100(b *testing.B)   { bench(b, 10, buf) }
func BenchmarkCat10000(b *testing.B) { bench(b, 10000, cat) }
func BenchmarkBuf10000(b *testing.B) { bench(b, 10000, buf) }
