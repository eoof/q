package main

import (
	"flag"
	"fmt"

	"github.com/eoof/q"
	"github.com/eoof/q/example/demopkg"
)

func main() {
	flag.Parse() //ql and qo added
	data1 := struct {
		A int
		B string
	}{A: 1, B: "foo"}
	q.Q(data1)
	fmt.Println("main, printing result of demopkg.Demo1", demopkg.Demo1())
	demo2()
}

func demo2() {
	q.Q("demo2")
}
