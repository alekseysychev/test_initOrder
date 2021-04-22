package main

import _ "github.com/alekseysychev/test_initOrder/subPackage"

func init() {
	println("main init")
}

func main() {
	println("main main")
}
