package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	通过url获取 utf-8 内容
 */
func Fetch(url string) ([]byte, error) {
	resp,err := http.Get(url)
	if  err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	// 解码
	bodyReader := bufio.NewReader(resp.Body)
	encoder := determineEncoding(bodyReader)
	newReader := transform.NewReader(bodyReader, encoder.NewDecoder())

	return ioutil.ReadAll(newReader)
}

/*
	获取编码
 */
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch encoder err: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}