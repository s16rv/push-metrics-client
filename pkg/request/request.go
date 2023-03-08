package request

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	ContentTypeTextPlain = "text/plain"
)

func GetBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostText(url string, body string) error {
	resp, err := http.Post(url, ContentTypeTextPlain, strings.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error code %d, failed to post", resp.StatusCode)
	}

	return nil
}
