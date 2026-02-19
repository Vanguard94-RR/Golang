package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, _ := exec.Command("ls", "-ltrsh").Output()
	fmt.Println(string(out))

}
