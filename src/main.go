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
	resp, err := getResponseBody(url)
	handleError(err)
	println(resp)

	out, err := exec.Command("ls", "-la").Output()
	handleError(err)
	println(string(out))
}

func getResponseBody(url string) (response string, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	return string(byteArray), err
}

func handleError(err error) {
	if err != nil {
		color.Red("Error %s", err)
		os.Exit(1)
	}
}
