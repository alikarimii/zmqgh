package shared

import sync "sync"

func Newlock() Locke {
	return Locke{c: make(chan struct{})}
}

type Locke struct {
	c chan struct{}
}

func (b *Locke) Unlock() {
	close(b.c)
}
func (b *Locke) Lock(fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		<-b.c
		fn()
	}()
	goroutineRunning.Wait()
}
