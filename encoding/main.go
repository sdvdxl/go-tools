package main

import (
	"fmt"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"bytes"
	"bufio"
)

func main() {

	sr := strings.NewReader("你啊后")
	tr := transform.NewReader(sr, simplifiedchinese.GB18030.NewEncoder())


	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	io.Copy(writer, tr)
	fmt.Println(b.String())


}