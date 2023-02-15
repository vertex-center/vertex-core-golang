package pubsub

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Pub(key string, message string) {
	data := url.Values{}
	requestURL := fmt.Sprintf("http://localhost:6140/pub/%s", key)

	req, err := http.NewRequest("POST", requestURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
