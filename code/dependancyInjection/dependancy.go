package code

import (
	"bytes"
	"fmt"
)

func Printer(pass *bytes.Buffer, s string) {
	fmt.Fprintf(pass, "%v, Printer responding", s)

}
