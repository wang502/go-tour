package main

import (
	"errors"
	"fmt"
	"sync"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type walkFunc func(val int) error

func walk(root *treeNode, walkFn walkFunc) error {
	if root == nil {
		return nil
	}
	err := walkFn(root.val)
	if err != nil {
		return err
	}
	err = walk(root.left, walkFn)
	if err != nil {
		return err
	}
	return walk(root.right, walkFn)
}

func walkTree(root *treeNode, done <-chan struct{}) (<-chan int, <-chan error) {
	vals := make(chan int)
	errc := make(chan error, 1)
	go func() {
		defer close(vals)
		errc <- walk(root, func(val int) error {
			select {
			case vals <- val:
			case <-done:
				return errors.New("tree walk canceled")
			}
			return nil
		})
	}()

	return vals, errc
}

func valDigester(vals <-chan int, result chan<- int, done <-chan struct{}) {
	for val := range vals {
		square := val * val
		select {
		case result <- square:
		case <-done:
			return
		}
	}
}

func squreAllInt(root *treeNode) ([]int, error) {
	done := make(chan struct{})
	defer close(done)

	result := make(chan int)
	vals, errc := walkTree(root, done)
	const numDigesters = 10
	var wg sync.WaitGroup
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			valDigester(vals, result, done)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var res []int
	for r := range result {
		res = append(res, r)
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return res, nil
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		c <- 10
	}()
	return c
}

func main() {

	root := &treeNode{
		10,
		&treeNode{
			5,
			&treeNode{1, nil, nil},
			&treeNode{7, nil, nil},
		},
		&treeNode{20, nil, nil},
	}

	res, err := squreAllInt(root)
	if err != nil {
		fmt.Println(err)
	}
	for _, num := range res {
		fmt.Println(num)
	}

	/*
		c := gen()
		fmt.Println(<-c)

		c2 := make(chan int)
		go func() {
			fmt.Println(<-c2)
		}()
		c2 <- 10

		c3 := make(chan int, 1)
		c3 <- 10
		fmt.Println(<-c3)

		c4 := make(chan int)
		c4 <- 10
		fmt.Println(<-c4)
	*/
}
