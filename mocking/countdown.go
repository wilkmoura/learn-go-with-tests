package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

// stopped at https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#write-the-test-first-2
