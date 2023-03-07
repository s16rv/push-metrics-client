package request

import (
	"io"
	"net/http"
)

func GetBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, nil
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
