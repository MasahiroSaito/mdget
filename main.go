package main

import (
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("args: <url>")
		return
	}

	url := args[0]

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer resp.Body.Close()
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
