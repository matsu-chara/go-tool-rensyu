package main

import (
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	url := "http://google.co.jp"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	handleError(err)

	byteArray, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	println(string(byteArray))

	out, err := exec.Command("ls", "-la").Output()
	handleError(err)
	println(string(out))
}

func handleError(err error) {
	if err != nil {
		color.Red("Error %s", err)
		os.Exit(1)
	}
}
