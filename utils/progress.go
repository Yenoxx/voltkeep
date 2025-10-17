package utils

import (
	"sync"
)

// Struct used to track progress of operations
type Progress struct {
	count uint64
	total uint64
	ratio float32

	mu sync.Mutex
}

func CreateProgress(total any) *Progress {
	ntotal, ok := total.(uint64)
	if !ok {
		panic("uint expected")
	}

	p := &Progress{}
	p.total = ntotal
	return p
}

func (p *Progress) Increment() {
	p.mu.Lock()
	p.count += 1
	p.ratio = float32(p.count) / float32(p.total)
	p.mu.Unlock()
}

func (p *Progress) Ratio() float32 {
	return p.ratio
}
