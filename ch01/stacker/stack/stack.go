package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Push(elem interface{}) {
	*stack = append(*stack, elem)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[stack.lastElemIdx()], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("can't Pop() an empty stack")
	}

	elem := (*stack)[stack.lastElemIdx()]
	*stack = (*stack)[:stack.lastElemIdx()]

	return elem, nil
}

func (stack Stack) lastElemIdx() int {
	return stack.Len() - 1
}