package main

func printIntArray(a []int) {
	for _, v := range a {
		println(v)
	}
}

func main() {
	a := []int{5, 2, 6, 3, 1, 4}
	b := []int{9, 1, 3, 5, 6, 2}
	printIntArray(a)
	printIntArray(b)
}
