type element struct {
	val int
	min int

}

type MinStack struct {
	elements []element
}

func Constructor() MinStack {
	return MinStack {
		elements: []element{},
	}
}

func (this *MinStack) Push(val int) {
	min := val

	if len(this.elements) == 0 {
		this.elements = append(this.elements, element{
			val: val,
			min: min,
		})
		return
	}

	currMin := this.elements[len(this.elements) - 1].min
	
	if currMin < val {
		min = currMin
	}

	this.elements = append(this.elements, element {
		val: val,
		min: min,
	})
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:(len(this.elements) - 1)]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1].val
}

func (this *MinStack) GetMin() int {
	return this.elements[len(this.elements)-1].min
}