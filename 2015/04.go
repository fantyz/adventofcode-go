package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"strconv"
)

var input = `iwrupvqb`

func main() {
	i := -1
Loop:
	for {
		i++
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, strconv.Itoa(i))
		hstr := fmt.Sprintf("%x", h.Sum(nil))
		for j:= 0; j<6; j++ {
			if hstr[j] != '0' {
				continue Loop
			}
		}
		fmt.Println("Number:", i, "MD5:", hstr)
		break
	}
}
