package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a download url")
		return
	}

	dUrl := os.Args[1]

	resp, err := http.Get(dUrl)
	if err != nil {
		panic(err)
	}

	fileName := path.Base(dUrl)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	io.Copy(f, resp.Body)
}
