package main

import (
	"context"
	"fmt"
	"time"
)

// Timing runs fn with the given time limit. If a call runs for longer than its time limit or panic,
// it will return context.DeadlineExceeded error or panic error.
func Timing(dt time.Duration, fn func(context.Context)) (err error) {
	ct, cancel := context.WithTimeout(context.Background(), dt)
	defer cancel()

	ch := make(chan struct{})
	go func() {
		// recover the fn call
		//time.Sleep(49 * time.Millisecond)
		time.Sleep(50 * time.Millisecond)
		//time.Sleep(100 * time.Millisecond)
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("Timing panic: %#v", e)
			}
			close(ch)
		}()
		fn(ct)
	}()
	select {
	case <-ct.Done():
		err = ct.Err()
	case <-ch:
	}
	return
}

func main() {

	f := func(context.Context) {
		fmt.Printf("exe!\n")
	}

	if err := Timing(50*time.Millisecond, f); err != nil {
		fmt.Printf("exe timeout: %v\n", err)
	}
}
