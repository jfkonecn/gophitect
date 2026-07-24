package main

import "fmt"

type Something struct {
	num int
}

func refTest(x *Something) {
	x.num++
	fmt.Printf("refTest %v\n", x.num)
}

func valTest(x Something) {
	x.num++
	fmt.Printf("valTest %v\n", x.num)
}

func main() {
	temp := Something{num: 1}
	fmt.Printf("%v\n", temp.num)
	refTest(&temp)
	fmt.Printf("%v\n", temp.num)
	valTest(temp)
	fmt.Printf("%v\n", temp.num)
}
