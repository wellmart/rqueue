//
//  Ring Queue
//
//  Copyright (c) 2020 Wellington Marthas
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in
//  all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//  THE SOFTWARE.
//

package rqueue

const minQueueLength = 32

// Queue represents a queue data structure.
type Queue struct {
	items []interface{}

	head   int
	tail   int
	length int
}

// New returns a new Queue.
func New() *Queue {
	return &Queue{
		items: make([]interface{}, minQueueLength),
	}
}

// Capacity returns the capacity of queue.
func (q *Queue) Capacity() int {
	return len(q.items)
}

// Length returns number of elements in queue.
func (q *Queue) Length() int {
	return q.length
}

// Enqueue puts an item to the tail of the queue.
func (q *Queue) Enqueue(item interface{}) {
	if q.length == len(q.items) {
		q.resize()
	}

	q.items[q.tail] = item

	q.tail = (q.tail + 1) & (len(q.items) - 1)
	q.length++
}

// TryDequeue try removes the item at the head of the queue and returns it.
func (q *Queue) TryDequeue() (item interface{}, ok bool) {
	if q.length <= 0 {
		ok = false
		return
	}

	item = q.items[q.head]
	ok = true

	q.items[q.head] = nil

	q.head = (q.head + 1) & (len(q.items) - 1)
	q.length--

	if len(q.items) > minQueueLength && (q.length<<2) == len(q.items) {
		q.resize()
	}

	return
}

func (q *Queue) resize() {
	newItems := make([]interface{}, q.length<<1)

	if q.tail > q.head {
		copy(newItems, q.items[q.head:q.tail])
	} else {
		i := copy(newItems, q.items[q.head:])
		copy(newItems[i:], q.items[:q.tail])
	}

	q.items = newItems

	q.head = 0
	q.tail = q.length
}
