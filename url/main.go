package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&price=100000"

func main() {
	println(myUrl)

	result, _ := url.Parse(myUrl)
	//fmt.Println(result.Scheme)
	//fmt.Println(result.Host)
	//fmt.Println(result.Path)
	//fmt.Println(result.RawQuery)
	//fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params : %T \n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		println("Params is: ", val)
	}
}
