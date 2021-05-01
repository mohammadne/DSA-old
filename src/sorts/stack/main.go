package main

import "fmt"

type Stack struct {
	top *StackItem
}

func (stack *Stack) pop() *StackItem {
	top := stack.top

	if top != nil {
		stack.top = top.bottom
	} else {
		stack.top = nil
	}

	return top
}

func (stack *Stack) push(item *StackItem) {
	stack.top = item
}

type StackItem struct {
	bottom *StackItem
	data   int64
}

func main() {
	i1 := StackItem{
		data:   3,
		bottom: nil,
	}

	i2 := StackItem{
		data:   2,
		bottom: &i1,
	}

	i3 := StackItem{
		data:   7,
		bottom: &i2,
	}

	mainStack := Stack{
		top: &i3,
	}

	helperStack := Stack{}
	sort(&mainStack, &helperStack)

	fmt.Println(mainStack)
	fmt.Println(helperStack)
}

func sort(stack *Stack, helperStack *Stack) {
	popedItemFromStack := stack.pop()
	popedItemFromHelperStack := helperStack.pop()

	if popedItemFromStack == nil || popedItemFromHelperStack == nil {
		return
	} else {
		if popedItemFromHelperStack.data <= popedItemFromStack.data {
			helperStack.push(popedItemFromHelperStack)
			helperStack.push(popedItemFromStack)
			sort(stack, helperStack)
		} else {
			stack.push(popedItemFromHelperStack)
			stack.push(popedItemFromStack)
			sort(stack, helperStack)
		}
	}

}
