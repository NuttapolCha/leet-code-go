package main

func main() {

}

type Collection []int

func (c Collection) Len() int           { return len(c) }
func (c Collection) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Collection) Less(i, j int) bool { return c[i] > c[j] }

func (c Collection) Remove(e int) {
	return
}

func lastStoneWeight(stones []int) int {
	return 0
}
