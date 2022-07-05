// Implement a first in first out (FIFO) queue using only two stacks.
// The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

package datastructures

type MyQueue struct {
	popStack  stackInt
	pushStack stackInt
}

func Constructor() MyQueue {
	return MyQueue{
		popStack:  make(stackInt, 0),
		pushStack: make(stackInt, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack.Push(x)
}

func (this *MyQueue) movePushStack() {
	for this.pushStack.Len() > 0 {
		this.popStack.Push(this.pushStack.Peek())
		this.pushStack.Pop()
	}
}

func (this *MyQueue) Pop() int {
	if this.popStack.Len() == 0 {
		this.movePushStack()
	}
	dead := this.popStack.Peek()
	this.popStack.Pop()
	return dead
}

func (this *MyQueue) Peek() int {
	if this.popStack.Len() == 0 {
		this.movePushStack()
	}
	return this.popStack.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.popStack.Len() == 0 && this.pushStack.Len() == 0
}

type stackInt []int

func (s *stackInt) Push(val int) {
	*s = append(*s, val)
}

func (s *stackInt) Pop() {
	if s.Len() > 0 {
		*s = (*s)[:s.Len()-1]
	}
}

func (s *stackInt) Peek() int {
	if s.Len() == 0 {
		return 0
	}
	return (*s)[s.Len()-1]
}

func (s *stackInt) Len() int {
	return len(*s)
}
