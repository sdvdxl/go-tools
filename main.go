package main
import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
"bytes"
	"bufio"
	"io"
)

func main() {

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


