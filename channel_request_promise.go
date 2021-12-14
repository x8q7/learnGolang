package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestFuture, http request promise.
func RequestFuture(url string) <-chan []byte {
	c := make(chan []byte, 1)
	go func() {
		var body []byte
		defer func() {
			c <- body
		}()

		res, err := http.Get(url)
		if err != nil {
			return
		}
		defer res.Body.Close()

		body, _ = ioutil.ReadAll(res.Body)
	}()

	return c
}

func main() {
	future := RequestFuture("https://www.baidu.com/")
	body := <-future
	fmt.Printf("reponse length: %s", string(body))
}
