package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sdvdxl/go-tools/datetime"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	timestamp := datetime.Timestamp(time.Now())
	tm, err := json.Marshal(timestamp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tm) + " ===" + fmt.Sprint(time.Now().UnixNano()/1000000))
}

func a() {
	resp, err := http.Get(`http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardn`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sr := strings.NewReader(string(bodyBytes))
	tr := transform.NewReader(sr, simplifiedchinese.GB18030.NewDecoder())

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	io.Copy(writer, tr)
	fmt.Println(b.String())
}
