package main

import "fmt"
import "math/rand"
import "net"
import "os"
import "time"

func main() {
	fmt.Println("Hello world")
	fmt.Println("Welcome to the Golang")
	fmt.Println("The time is ", time.Now())
	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("filename"))
	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "google.com"))
	fmt.Println("My favorite number is", rand.Intn(10))
}
