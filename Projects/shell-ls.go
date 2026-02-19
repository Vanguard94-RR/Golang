package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, _ := exec.Command("", "-ltrsh").Output()
	fmt.Println(string(out))

}
