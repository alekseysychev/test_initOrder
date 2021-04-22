package mainSubPackage1

import _ "github.com/alekseysychev/test_initOrder/mainSubPackage1/subPackage"

func init() {
	println("main -> mainSubPackage1 -> init()")
}
