package main

import (
	"fmt"
	"os/exec"
)

func connect() {

	out, _ := exec.Command("ls", "-ltrsh").Output()
	fmt.Println(string(out))

}

func ls() {
	out, _ := exec.Command("ls", "-ltrsh").Output()
	fmt.Println(string(out))
}


func main() {
	connect()
	ls()	
}
