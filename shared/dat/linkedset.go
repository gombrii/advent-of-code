package dat

import (
	"fmt"
	"strings"
)

type Linked[T comparable] struct {
	data map[T]*T
	head T
	tail T
}

func NewLinked[T comparable]() Linked[T] {
	return Linked[T]{data: make(map[T]*T)}
}

func NewLinkedFrom[T comparable](slice []T) Linked[T] {
	linked := Linked[T]{data: make(map[T]*T)}
	if len(slice) == 0 {
		return linked
	}

	linked.head = slice[0]
	last := slice[0]
	for i := 1; i < len(slice); i++ {
		linked.data[last] = &slice[i]
		last = slice[i]
	}

	linked.data[last] = nil
	linked.tail = last
	return linked
}

func (l *Linked[T]) Head() T {
	return l.head
}

func (l *Linked[T]) Tail() T {
	return l.tail
}

func (l *Linked[T]) Len() int {
	return len(l.data)
}

func (l *Linked[T]) Append(item T) {
	//check if exists
	//Check so the last tail is not item
	if len(l.data) == 0 {
		l.data[item] = nil
		l.head = item
	} else {
		l.data[l.tail] = &item
		l.data[item] = nil
	}
	l.tail = item
}

func (l *Linked[T]) Prepend(item T) {
	//check if exists
	//Check so the last head is not item
	if len(l.data) == 0 {
		l.data[item] = nil
		l.tail = item
	} else {
		head := l.head
		l.data[item] = &head
	}
	l.head = item
}

func (l *Linked[T]) Remove(item T) {
	next, exists := l.data[item]
	if !exists {
		return
	}
	var prev T
	for _, v := range l.Iter {
		if v == item {
			break
		}
		prev = v
	}

	delete(l.data, item)
	l.data[prev] = next
}

func (l *Linked[T]) Insert(item T, after T) error {
	//check so item and after are not the same (if needed)
	if _, exists := l.data[after]; !exists {
		// panic instead??
		return fmt.Errorf("after is no item in list [%v]", after)
	}
	_, exists := l.data[item]
	if exists {
		l.Remove(item)
	}
	prev := after
	next := l.data[prev]

	l.data[prev] = &item
	l.data[item] = next
	return nil
}

func (l Linked[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for _, v := range l.Iter {
		builder.WriteString(fmt.Sprintf("%v", v))
		if v != l.tail {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (l *Linked[T]) Iter(yield func(int, T) bool) {
	if len(l.data) == 0 {
		return
	}
	last := &l.head
	for i := 0; last != nil; i++ {
		if !yield(i, *last) {
			return
		}
		last = l.data[*last]
	}
}

func (l *Linked[T]) IndexOf(item T) int {
	//check if exists
	if l.head == item {
		return 0
	}
	if l.tail == item {
		return len(l.data) - 1
	}

	for i, v := range l.Iter {
		if v == item {
			return i
		}
	}
	return 0 // should return error or panic, probably panic
}

func (l *Linked[T]) Contains(item T) bool {
	_, exits := l.data[item]
	return exits
}
