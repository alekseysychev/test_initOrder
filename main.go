package main

import (
	_ "github.com/alekseysychev/test_initOrder/mainSubPackage1"
	pkg2 "github.com/alekseysychev/test_initOrder/mainSubPackage2"
)

func init() {
	println("main -> init()")
}

func main() {
	println("main -> main()")
	println(pkg2.Show())
}
