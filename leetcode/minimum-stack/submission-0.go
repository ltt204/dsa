type MinStack struct {
	items []int
}

func Constructor() MinStack {
	return MinStack {
		items: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	this.items = this.items[:(len(this.items) - 1)]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	minVal := math.MaxInt

	for _, val := range this.items {
		if minVal > val {
			minVal = val
		}
	}

	return minVal
}