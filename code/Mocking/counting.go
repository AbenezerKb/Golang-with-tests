package code

import (
	"fmt"
	"io"
	"time"
)

func countdown(buffer io.Writer) {

	//num := 3

	for i := 3; i >= 1; i-- {
		fmt.Fprintln(buffer, i)
		time.Sleep(time.Second)
		fmt.Println("second: ", i)
	}
	fmt.Fprint(buffer, "Go")

	// )
}
