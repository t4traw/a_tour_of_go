package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the golang")
	fmt.Println("The time is", time.Now())
	fmt.Println("And id you try to open a file:")
	fmt.Println(os.Open("filename"))
	fmt.Println("Or aceess the network:")
	fmt.Println(net.Dial("tcp", "google.com"))
}
