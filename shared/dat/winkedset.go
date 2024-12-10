package dat

import (
	"fmt"
	"strings"
)

const falloffCount = 500

type Winked[T comparable] struct {
	data map[T]*T
	wgts map[T]float64
	head T
	tail T
	tick int
}

func NewWinked[T comparable]() Winked[T] {
	return Winked[T]{data: make(map[T]*T), wgts: make(map[T]float64)}
}

func NewWinkedFrom[T comparable](slice []T) Winked[T] {
	new := Winked[T]{data: make(map[T]*T), wgts: make(map[T]float64)}
	if len(slice) == 0 {
		return new
	}

	wIncr := 1 / float64(len(slice))

	new.head = slice[0]
	last := slice[0]
	for i := 1; i < len(slice); i++ {
		new.data[last] = &slice[i]
		new.wgts[last] = float64(i) * wIncr
		last = slice[i]
	}

	new.data[last] = nil
	new.wgts[last] = 1
	new.tail = last
	return new
}

func (l *Winked[T]) Head() T {
	return l.head
}

func (l *Winked[T]) Tail() T {
	return l.tail
}

func (l *Winked[T]) Len() int {
	return len(l.data)
}

func (l *Winked[T]) Append(item T) {
	l.acc()
	//check if exists
	//Check so the last tail is not item
	if len(l.data) == 0 {
		l.data[item] = nil
		l.wgts[item] = 0
		l.head = item
	} else {
		l.data[l.tail] = &item
		l.data[item] = nil
		l.wgts[item] = (l.wgts[l.tail] + 2) / 2
	}
	l.tail = item
}

func (l *Winked[T]) Prepend(item T) {
	l.acc()
	//check if exists
	//Check so the last head is not item
	if len(l.data) == 0 {
		l.data[item] = nil
		l.wgts[item] = 0
		l.tail = item
	} else {
		head := l.head
		l.wgts[head] = l.wgts[*l.data[head]] / 2
		l.data[item] = &head
		l.wgts[item] = 0
	}
	l.head = item
}

func (l *Winked[T]) Remove(item T) {
	next, exists := l.data[item]
	if !exists {
		return
	}
	if item == l.head {
		l.head = *l.data[item]
		delete(l.data, item)
		delete(l.wgts, item)
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

func (l *Winked[T]) Insert(item T, after T) {
	l.acc()
	//check so item and after are not the same (if needed)
	if _, exists := l.data[after]; !exists {
		// panic instead??
		panic(fmt.Errorf("after is no item in list [%v]", after))
	}
	_, exists := l.data[item]
	if exists {
		l.Remove(item)
	}
	if after == l.tail {
		l.Append(item)
		return
	}

	prevI := after
	prevW := l.wgts[after]
	nextI := l.data[prevI]
	nextW := l.wgts[*nextI]

	l.data[prevI] = &item
	l.data[item] = nextI
	l.wgts[item] = (prevW + nextW) / 2
}

func (l Winked[T]) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for _, v := range l.Iter {
		builder.WriteString(fmt.Sprintf("{%v:%f}", v, l.wgts[v]))
		if v != l.tail {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (l *Winked[T]) Iter(yield func(int, T) bool) {
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

func (l *Winked[T]) Contains(item T) bool {
	_, exits := l.data[item]
	return exits
}

func (l *Winked[T]) After(item T, other T) bool {
	otherW, exists := l.wgts[other]
	if !exists {
		panic(fmt.Sprintf("other doesn't exist in set [%v]", other))
	}
	currW := l.wgts[item]
	return currW > otherW
}

func (l *Winked[T]) acc() {
	l.tick++
	if l.tick > falloffCount {
		fmt.Println("---BEFORE REWEIGHING---")
		for _, v := range l.Iter {
			fmt.Println(v, l.wgts[v])
		}
		wIncr := 1 / float64(len(l.data))
		for i, v := range l.Iter {
			l.wgts[v] = float64(i) * wIncr
		}
		fmt.Println("---AFTER REWEIGHING---")
		for _, v := range l.Iter {
			fmt.Println(v, l.wgts[v])
		}
	}
}
