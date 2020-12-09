package stack_slice_implemention

type Stack struct {
	top int
	stack []int
}

func (stack *Stack) Push(value int) {
	stack.stack = append(stack.stack, value)
	stack.top++
}

func (stack *Stack) Pop() int {
	res := stack.stack[stack.top - 1]
	stack.top -= 1
	stack.stack	 = stack.stack[0:stack.top]
	return res
}