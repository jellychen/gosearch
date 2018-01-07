package obtainer

import (
	"bufio"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

func DetectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, ""); ok {
			return name
		}
	}
	return "utf-8"
}

func DecodeHTMLBody_Utf8(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = DetectContentCharset(body)
	}
	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}
	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}
	return body, nil
}
