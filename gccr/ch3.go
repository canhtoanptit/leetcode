package gccr

import (
	"fmt"
	"time"
)

func F() {
	ch := make(chan struct{})
	stringStream := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- struct{}{}
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-stringStream:
			fmt.Println("stringStream")
		}
	}()
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-ch:
			fmt.Println(s)
			return
		case stringStream <- s:
			fmt.Println("stringStream <- s")
		}
	}
}

func DoWork() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()

		return completed
	}

	strs := make(chan string)
	go func() {
		strs <- "a"
		strs <- "b"
		strs <- "c"
		close(strs)
	}()
loop:
	for {
		select {
		case <-doWork(strs):
			break loop
		}
	}
	fmt.Println("Done.")
}

func NewRunStream() {
	newRanStream := func(done <-chan interface{}, nums ...int) <-chan int {
		rStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(rStream)
			for _, n := range nums {
				select {
				case <-done:
					//return
					fmt.Println("ignore done")
				case rStream <- n:
				}
			}
		}()
		return rStream
	}

	done := make(chan interface{})
	rStream := newRanStream(done, 1, 2, 3)
	for r := range rStream {
		fmt.Println(r)
		if r == 2 {
			close(done)
		}
	}
}
