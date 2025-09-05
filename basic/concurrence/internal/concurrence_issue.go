package internal

import (
	"fmt"
	"time"
)

func ConcurrenceIssue() {
	sum := 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				sum++
			}
		}()
	}

	time.Sleep(time.Second * 3)
	fmt.Println(sum)
}

func ConcurrenceIssueWithSlice() {
	slice := make([]int, 0, 10000)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				slice = append(slice, 1)
			}
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println(len(slice))

}

func ConcurrenceIssueWithString() {
	title := "hello world"
	go func() {
		for {
			fmt.Println(title)
			for _, _ = range title {
			}
		}
	}()

	for {
		go func() {
			title = ""
		}()
		go func() {
			title = "hello world"
		}()
	}
}
