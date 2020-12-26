package stackslice

// Stack 用slice实现stack
type Stack struct {
	top   int
	stack []int
}

// Push 入栈
func (stack *Stack) Push(value int) {
	stack.stack = append(stack.stack, value)
	stack.top++
}

// Pop 出栈
func (stack *Stack) Pop() int {
	res := stack.stack[stack.top-1]
	stack.top--
	stack.stack = stack.stack[0:stack.top]
	return res
}
