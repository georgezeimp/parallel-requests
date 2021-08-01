package concurrency

import "sync"

type ConcurrentSlice struct {
	sync.RWMutex
	items []string
}

func NewConcurrentSlice(items []string) *ConcurrentSlice {
	return &ConcurrentSlice{
		items: items,
	}
}

func (slice *ConcurrentSlice) Append(item string) {
	slice.Lock()
	defer slice.Unlock()

	slice.items = append(slice.items, item)
}

func (slice *ConcurrentSlice) GetItems() []string {
	return slice.items
}
