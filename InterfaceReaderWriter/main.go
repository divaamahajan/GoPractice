package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
func main () {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lw := logWriter{}

	// io.Copy(writer,reader)
	io.Copy(lw,resp.Body)
	// io.Copy(os.Stdout,resp.Body)
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

}

func (lg logWriter) Write (bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("#Bytes=",len(bs))
	return len(bs), nil

}
