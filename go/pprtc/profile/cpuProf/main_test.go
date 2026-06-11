// main_test.go

package main

import (
	"testing"

	"os"

	"runtime/pprof"
)

func BenchmarkAdd(b *testing.B) {

	f, err := os.Create("cpu.pprof")

	if err != nil {

		b.Fatal(err)

	}

	defer f.Close()

	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {

		Add(3, 5)

	}

}
