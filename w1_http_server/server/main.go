package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// r 带的header 写入response中
	fmt.Println("request header")
	for k, v := range r.Header {
		fmt.Printf("%s=%v\n", k, v)
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	w.Header().Add("os.Version", os.Getenv("Version"))
	fmt.Println("request body")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	fmt.Printf("remote address: %v\n", r.RemoteAddr)

	fmt.Fprint(w, "server msg")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}
