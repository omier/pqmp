package pqmp

import (
	"reflect"
	"testing"
)

func Test_priorityQueue_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name      string
		pq        *priorityQueue
		args      args
		want      *priorityQueue
		wantPanic bool
	}{
		{
			name: "Push to Empty Queue",
			pq:   &priorityQueue{},
			args: args{
				x: &Item{
					value:    "A",
					priority: 0,
				},
			},
			want: &priorityQueue{&Item{value: "A", priority: 0, index: 0}},
		},
		{
			name: "Push To non-empty Queue",
			pq:   &priorityQueue{&Item{value: "A", priority: 0, index: 0}},
			args: args{
				x: &Item{
					value:    "B",
					priority: 1,
				},
			},
			want: &priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
		},
		{
			name: "Push non *Item Value To Queue",
			pq:   &priorityQueue{&Item{value: "A", priority: 1, index: 0}},
			args: args{
				x: nil,
			},
			want:      nil,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("priorityQueue.Swap(%v) recover = %v, wantPanic = %v", tt.args.x, r, tt.wantPanic)
				}
			}()
			tt.pq.Push(tt.args.x)
			for i := 0; i < tt.pq.Len(); i++ {
				if (*tt.pq)[i].value != (*tt.want)[i].value ||
					(*tt.pq)[i].priority != (*tt.want)[i].priority ||
					(*tt.pq)[i].index != (*tt.want)[i].index {
					t.Errorf("priorityQueue.Push(%v) = %v, want %v", tt.args.x, *tt.pq, *tt.want)
				}
			}
		})
	}
}

func Test_priorityQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		pq   *priorityQueue
		want interface{}
	}{
		{
			name: "Pop From Empty Queue",
			pq:   &priorityQueue{},
			want: nil,
		},
		{
			name: "Pop From Queue With Single Item",
			pq:   &priorityQueue{&Item{value: "A", priority: 0, index: 0}},
			want: &Item{value: "A", priority: 0, index: -1},
		},
		{
			name: "Pop From Queue With Multiple Items",
			pq:   &priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			want: &Item{value: "B", priority: 1, index: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("priorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name      string
		pq        priorityQueue
		args      args
		want      priorityQueue
		wantPanic bool
	}{
		{
			name: "Swap First Element With 2nd Element in Queue",
			pq:   priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			args: args{i: 0, j: 1},
			want: priorityQueue{&Item{value: "B", priority: 1, index: 0}, &Item{value: "A", priority: 0, index: 1}},
		},
		{
			name: "Swap First Element With Itself",
			pq:   priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			args: args{i: 0, j: 0},
			want: priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
		},
		{
			name:      "Swap in Queue With Single Item and Index Out of Bounds",
			pq:        priorityQueue{&Item{value: "A", priority: 0, index: 0}},
			args:      args{i: 0, j: 1},
			want:      nil,
			wantPanic: true,
		},
		{
			name:      "Swap in Empty Queue",
			pq:        priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			args:      args{i: 0, j: 1},
			want:      nil,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("priorityQueue.Swap(%v,%v) recover = %v, wantPanic = %v", tt.args.i, tt.args.j, r, tt.wantPanic)
				}
			}()
			tt.pq.Swap(tt.args.i, tt.args.j)
			if tt.pq[tt.args.i].value != tt.want[tt.args.i].value ||
				tt.pq[tt.args.i].priority != tt.want[tt.args.i].priority ||
				tt.pq[tt.args.i].index != tt.want[tt.args.i].index ||
				tt.pq[tt.args.j].value != tt.want[tt.args.j].value ||
				tt.pq[tt.args.j].priority != tt.want[tt.args.j].priority ||
				tt.pq[tt.args.j].index != tt.want[tt.args.j].index {
				t.Errorf("priorityQueue.Swap(%v,%v) = %v, want %v", tt.args.i, tt.args.j, tt.pq, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name      string
		pq        priorityQueue
		args      args
		want      bool
		wantPanic bool
	}{
		{
			name: "First Element is Less Than 2nd Element in Queue",
			pq:   priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			args: args{i: 0, j: 1},
			want: true,
		},
		{
			name: "First Element isn't Less Than Itself",
			pq:   priorityQueue{&Item{value: "A", priority: 0, index: 0}, &Item{value: "B", priority: 1, index: 1}},
			args: args{i: 0, j: 0},
			want: false,
		},
		{
			name:      "Less in Empty Queue",
			pq:        priorityQueue{},
			args:      args{i: 0, j: 1},
			want:      false,
			wantPanic: true,
		},
		{
			name:      "Less in Queue With Single Item",
			pq:        priorityQueue{&Item{value: "A", priority: 0, index: 0}},
			args:      args{i: 0, j: 1},
			want:      false,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("priorityQueue.Less(%v,%v) recover = %v, wantPanic = %v", tt.args.i, tt.args.j, r, tt.wantPanic)
				}
			}()
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("priorityQueue.Less(%v,%v) = %v, want %v", tt.args.i, tt.args.j, got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_update(t *testing.T) {
	A := &Item{value: "A", priority: 0, index: 0}
	B := &Item{value: "B", priority: 1, index: 1}
	type args struct {
		item     *Item
		value    interface{}
		priority int
	}
	tests := []struct {
		name string
		pq   *priorityQueue
		args args
		want *Item
	}{
		{
			name: "update",
			pq:   &priorityQueue{A, B},
			args: args{
				item:     A,
				value:    "C",
				priority: 2,
			},
			want: &Item{value: "C", priority: 2, index: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.update(tt.args.item, tt.args.value, tt.args.priority)
			for _, item := range *tt.pq {
				if item == tt.args.item &&
					(item.value != tt.want.value ||
						item.priority != tt.want.priority ||
						item.index != tt.want.index) {
					t.Errorf("priorityQueue.update(%v) = %v, want %v", tt.args.item, item, tt.want)
				}
			}
		})
	}
}
