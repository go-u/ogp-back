package ogps

import (
	"bufio"
	"bytes"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func CopyBodyNonDestructive(res *http.Response) io.ReadCloser {
	buf, _ := ioutil.ReadAll(res.Body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	body_copy := ioutil.NopCloser(bytes.NewBuffer(buf))
	return body_copy

}
func GetDecoder(body io.ReadCloser, header http.Header) (*encoding.Decoder, error) {
	br := bufio.NewReader(body)
	data, err := br.Peek(1024)
	if err != nil {
		return nil, err
	}
	enc, _, _ := charset.DetermineEncoding(data, header.Get("content-type"))
	return enc.NewDecoder(), nil
}

func isValidFQDN(s string) (bool, error) {
	if strings.HasPrefix(s, "http") {
		return false, errors.New("no prefix allowed")
	}
	if strings.HasPrefix(s, "/") {
		return false, errors.New("no suffix allowed")
	}
	return true, nil
}

func isValidOgpImageUrl(s string) (bool, error) {
	if !strings.HasPrefix(s, "https") {
		return false, errors.New("image url is not https ")
	}
	return true, nil
}
