package mainSubPackage2

func init() {
	println("main -> mainSubPackage2 -> init()")
}

func Show() string {
	return "main -> main() -> mainSubPackage2 -> Show()"
}
