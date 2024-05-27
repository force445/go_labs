package routine

import "fmt"

func Routine(s string, count int, ch chan<- string) {
	for i := 0; i < count; i++ {
		fmt.Println(s)
	}
	ch <- "sending into channel"
}
