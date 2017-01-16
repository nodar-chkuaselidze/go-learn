package main

import (
	"testing"
)

const parallelism = 8

func BenchmarkCmapLockSet(b *testing.B) {
	cmapl := newCMapLockInt()

	b.SetParallelism(parallelism)

	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapl.set(12, 1523)
		}
	})
}

func BenchmarkCmapLockGet(b *testing.B) {
	cmapl := newCMapLockInt()
	cmapl.set(123, 14111)

	b.SetParallelism(parallelism)
	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapl.get(123)
		}
	})
}

func BenchmarkCmapLockDel(b *testing.B) {
	cmapl := newCMapLockInt()

	for i := 0; i < b.N; i++ {
		cmapl.set(i, i)
	}

	b.SetParallelism(parallelism)
	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapl.del(1)
		}
	})
}

func BenchmarkCmapChanSet(b *testing.B) {
	cmapc := newCapChanInt()

	b.SetParallelism(parallelism)

	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapc.set(12, 1523)
		}
	})

}

func BenchmarkCmapChanGet(b *testing.B) {
	cmapc := newCapChanInt()
	cmapc.set(123, 14111)

	b.SetParallelism(parallelism)
	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapc.get(123)
		}
	})
}

func BenchmarkCmapChanDel(b *testing.B) {
	cmapc := newCapChanInt()
	cmapc.set(123, 1)

	b.SetParallelism(parallelism)
	b.RunParallel(func(t *testing.PB) {
		for t.Next() {
			cmapc.del(123)
		}
	})
}
