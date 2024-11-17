# go-ds
Efficient implementations of data structures such as list, stack and queue. Ready to GO !!!

[![Go Report Card](https://goreportcard.com/badge/github.com/yashwanth-reddy909/go-ds)](https://goreportcard.com/report/github.com/yashwanth-reddy909/go-ds) 
[![Github Workflow](https://github.com/yashwanth-reddy909/go-ds/actions/workflows/main.yaml/badge.svg)](https://github.com/yashwanth-reddy909/go-ds/actions/workflows/main.yaml/badge.svg)

## Installation

```bash
go get github.com/yashwanth-reddy909/go-ds
```

## Examples

### list
 
```go
package main

import (
	"fmt"
	"github.com/yashwanth-reddy909/go-ds/list"
)

func main() {
	// Create a new list
	l := list.New[int]()

	// Append elements to the list
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Get the element at index 1
	x, ok := l.Get(1)
	fmt.Println("Element at index 1:", x, ok)

	// Remove the first element
	l.Remove(0) // [2, 3]

	l.Clear()
}
```

### stack
```go
package main

import (
	"fmt"
	"github.com/yashwanth-reddy909/go-ds/stack"
)

func main() {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)

	top, ok := s.Peek()
	if ok {
		fmt.Println("top element:", top) // 2
	}

	ok = s.Pop()
	if ok {
		fmt.Println("pop successful")
	}

	s.Clear()
}
```

### queue
```go
package main

import (
	"fmt"
	"github.com/yashwanth-reddy909/go-ds/queue"
)

func main() {
	q := queue.New[int]()

	q.Push(1)
	q.Push(2)

	front, ok := q.Peek()
	if ok {
		fmt.Println("front element:", front) // 1
	}

	ok = q.Pop() // true
	if ok {
		fmt.Println("pop successful")
	}

	q.Clear()
}
```

### heap
```go
package main

import (
	"fmt"

	"github.com/yashwanth-reddy909/go-ds/heap"
)

func main() {
	heap := heap.New[int]()
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	k, isPresent := heap.Peek()
	if isPresent {
		fmt.Println("top element is", k) // 1
	}

	heap.Clear()
}
```