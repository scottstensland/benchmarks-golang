
package main

import (
	"net/url"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	// testUrl := "https://www.example.com/path/file.html?param1=value1&param2=123"
	testUrl := "https://stackoverflow.com/questions/43245310/can-docker-be-deployed-on-version-2-6-x-of-the-linux-kernel"

	b.ResetTimer()
	for n := 0; n < b.N * 10000; n++ {
		_, err := url.Parse(testUrl)
		if err != nil {
			panic(err)
		}
	}
}

