package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Request() {
	reader := strings.NewReader("Hi Server")
	if req, err := http.NewRequest("GET", "http://localhost:8080", reader); err != nil {
		fmt.Println(err)
	} else {
		req.Header.Add("client_req", "request_header_msg")

		client := &http.Client{
			Timeout: 100 * time.Millisecond,
		}
		if resp, err := client.Do(req); err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("Header: \n")
			for k, v := range resp.Header {
				fmt.Printf("%s = %v\n", k, v)

			}

			fmt.Printf("body: \n")
			io.Copy(os.Stdout, resp.Body)
		}
	}
}

func main() {
	Request()
}
