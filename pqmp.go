package pqmp

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A priorityQueue implements heap.Interface and holds Items.
type priorityQueue []*Item

// Len is the number of elements in the collection.
func (pq priorityQueue) Len() int { return len(pq) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

// Swap swaps the elements with indexes i and j.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds x as element with index Len()
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item, ok := x.(*Item)
	if !ok {
		panic("x.(*Item) == nil")
	}

	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns element Len() - 1.
func (pq *priorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}

	item := (*pq)[pq.Len()-1]
	(*pq)[pq.Len()-1] = nil // avoid memory leak
	item.index = -1         // for safety
	*pq = (*pq)[0 : pq.Len()-1]

	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *priorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
