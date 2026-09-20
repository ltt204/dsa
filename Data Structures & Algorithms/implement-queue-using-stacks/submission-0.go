// 232. Implement Queue using Stacks
// https://leetcode.com/problems/implement-queue-using-stacks
//
// Use 2 Stack. We it goes to pop, we will move all the element within the first stack to the second one. Why?
// Example:
// 		1stStack=[1,2,3,4]; 2ndStack= [] (empty at first)
//		Hit pop: move all element from the 1stStack to 2ndStack, we have: 2ndStack=[4,3,2,1]  and 1sStack=[]
//	Now when every we pop, we will get the top of second stack, this way we are not violating the restriction of the problem.
// the logic above apply also for `peek()` method
//
// The Pop and peek are still considered as O(1), but Amortized-O(1). You can see all element are only moved at most ONCE,
// leading the average time complexity coverage to O(1)
//
// Note: Here is the restriction:
// ```
// peek & pop --> cannot access from index 0
//     because we are using stack for the implementation

// push --> cannot append from the index 0
//      because we are implementing the queue
// ```

package implementqueueusingstacks

type MyQueue struct {
	Stack    []int
	StackTop []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) move() {
	for len(this.Stack) > 0 {
		n := len(this.Stack)
		val := this.Stack[n-1]
		this.Stack = this.Stack[:n-1]

		this.StackTop = append(this.StackTop, val)
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.StackTop) == 0 {
		this.move()
	}

	val := this.StackTop[len(this.StackTop)-1]
	this.StackTop = this.StackTop[:len(this.StackTop)-1]

	return val
}

func (this *MyQueue) Peek() int {
	if len(this.StackTop) == 0 {
		this.move()
	}

	return this.StackTop[len(this.StackTop)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.Stack) == 0 && len(this.StackTop) == 0 {
		return true
	}

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
