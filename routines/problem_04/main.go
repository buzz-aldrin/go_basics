package main

import (
	"fmt"

	"sync"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	elem1 := make([]int, 0)
	elem2 := make([]int, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := range ch1 {
			elem1 = append(elem1, i)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := range ch2 {
			elem2 = append(elem2, i)
		}
		wg.Done()
	}()
	wg.Wait()

	if len(elem1) != len(elem2) {
		return false
	}

	for i := 0; i < len(elem1); i++ {
		if elem1[i] != elem2[i] {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)

	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t2, t3))
}
