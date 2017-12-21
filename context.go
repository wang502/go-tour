package main

import "context"
import "fmt"
import "sync"

var (
	numDigesters = 1
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		defer close(dst)
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func sq(ctx context.Context, c <-chan int, res chan<- int) error {
	for n := range c {
		square := n * n
		select {
		case <-ctx.Done():
			return ctx.Err()
		case res <- square:
		}
	}
	return nil
}

type message struct {
	responseChan chan<- int
	parameter    string
	ctx          context.Context
}

func processMessages(work <-chan message) {
	for job := range work {
		select {
		case <-job.ctx.Done():
			continue
		default:
		}

		hardToCalculate := len(job.parameter)
		select {
		case <-job.ctx.Done():
		case job.responseChan <- hardToCalculate:
		}
	}
}

func newRequest(ctx context.Context, input string, q chan<- message) {
	r := make(chan int)
	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q could see message")
		return
	case q <- message{
		responseChan: r,
		parameter:    input,
		// We are placing a context in a struct.  This is ok since it
		// is only stored as a passed message and we want q to know
		// when it can discard this message
		ctx: ctx,
	}:
	}

	select {
	case out := <-r:
		fmt.Printf("The len of %s is %d\n", input, out)
	case <-ctx.Done():
		fmt.Println("Context ended before q could process message")
	}
}

func main() {
	/*
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		out := make(chan int)

		c := gen(ctx)
		for i := 0; i < numDigesters; i++ {
			go func() {
				err := sq(ctx, c, out)
				if err != nil {
					fmt.Println(err)
				}
			}()
		}

		for n := range out {
			fmt.Println(n)
			if n == 25 {
				break
			}
		}
	*/
	q := make(chan message)
	go processMessages(q)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	inputs := []string{"hi", "hello"}
	wg.Add(len(inputs))
	for idx, in := range inputs {
		go func(in string, idx int) {
			newRequest(ctx, in, q)
			wg.Done()
			fmt.Printf("goroutine %d returned\n", idx)
		}(in, idx)
	}

	for i := 0; i < 100000; i++ {
	}
	cancel()

	wg.Wait()
}
