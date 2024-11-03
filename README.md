# go-ds
Efficient implementations of data structures such as list, stack and queue. Ready to GO !!!


## Examples

### Lists
 
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

	fmt.Println(s.Top()) // 2 true

	ok := s.Pop()
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

	fmt.Println(q.Top()) // 1
	
	ok := q.Pop() // true
	if ok {
		fmt.Println("pop successful")
	}

	q.Clear()
}
```

