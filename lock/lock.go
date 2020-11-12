package lock

import (
	"sync"
)

type mutex struct {
	c chan struct{}
}

// TryLocker represents an object that can be locked, unlocked and trylocked.
type TryLocker interface {
	sync.Locker
	TryLock() bool
}

// NewTryLocker returns an initialized object that implements TryLocker.
func NewTryLocker() TryLocker {
	return &mutex{c: make(chan struct{}, 1)}
}

// Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available.
func (m *mutex) Lock() {
	m.c <- struct{}{}
}

// Unlock unlocks m. It is a run-time error if m is not locked on entry to Unlock.
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.
func (m *mutex) Unlock() {
	select {
	case <-m.c:
		return
	default:
		panic("sync: unlock of unlocked mutex")
	}
}

// TryLock acquires the lock only if it is available and returns immediately with the value true,
// if the lock is not available then this method will return immediately with the value false.
func (m *mutex) TryLock() bool {
	select {
	case m.c <- struct{}{}:
		// lock acquired
		return true
	default:
		// lock not acquired
		return false
	}
}
