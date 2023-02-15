package pubsub

import (
	"bytes"
	"fmt"
	"net/http"
)

func Pub(key string, message []byte) {
	requestURL := fmt.Sprintf("http://localhost:6140/pub/%s", key)

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(message))
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
