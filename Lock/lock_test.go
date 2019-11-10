package lock

import (
	"testing"
)

func Test_mutex_Lock(t *testing.T) {
	unlockedChan := make(chan struct{}, 1)

	type fields struct {
		c chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Lock a unlocked mutex",
			fields: fields{c: unlockedChan},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutex{
				c: tt.fields.c,
			}
			m.Lock()
		})
	}
}

func Test_mutex_Unlock(t *testing.T) {
	lockedChan := make(chan struct{}, 1)
	lockedChan <- struct{}{}
	unlockedChan := make(chan struct{}, 1)

	type fields struct {
		c chan struct{}
	}
	tests := []struct {
		name      string
		fields    fields
		wantPanic bool
	}{
		{
			name:   "Unlock a locked mutex",
			fields: fields{c: lockedChan},
		},
		{
			name:      "Unlock a unlocked mutex",
			fields:    fields{c: unlockedChan},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("mutex.Unlock() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			m := &mutex{
				c: tt.fields.c,
			}
			m.Unlock()
		})
	}
}

func Test_mutex_TryLock(t *testing.T) {
	lockedChan := make(chan struct{}, 1)
	lockedChan <- struct{}{}
	unlockedChan := make(chan struct{}, 1)

	type fields struct {
		c chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TryLock a unlocked mutex",
			fields: fields{c: unlockedChan},
			want:   true,
		},
		{
			name:   "TryLock a locked mutex",
			fields: fields{c: lockedChan},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mutex{
				c: tt.fields.c,
			}
			if got := m.TryLock(); got != tt.want {
				t.Errorf("mutex.TryLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTryLocker(t *testing.T) {
	tests := []struct {
		name string
		want func(t TryLocker) bool
	}{
		{
			name: "NewTryLocker returns unlocked mutex",
			want: func(t TryLocker) bool { return t != nil },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTryLocker(); !tt.want(got) {
				t.Errorf("NewTryLocker() = %v, want %v", got, tt.want(got))
			}
		})
	}
}
