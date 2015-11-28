package main

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
    "./tsurapoyo"
    "./get"
)

func main() {
    url := "http://google.co.jp"
	resp, err := get.GetResponseBody(url)
	handleError(err)
	println(resp)

	out, err := exec.Command("ls", "-la").Output()
	handleError(err)
	println(string(out))

    poyo := &tsurapoyo.Tsurapoyo{"つらい", "とても"}
    println(poyo.HontoNoKimochi())
}


func handleError(err error) {
	if err != nil {
		color.Red("Error %s", err)
		os.Exit(1)
	}
}