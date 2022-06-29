package main

import (
	"builder/convertor"
	"fmt"
	"log"
)

func main() {
	cvrt, err := convertor.GetBuilder(2, 10)
	if err != nil {
		log.Fatal(err)
	}
	cvrt.Convert("10001")
	fmt.Println(cvrt.GetResult()) // 17

	cvrt, err = convertor.GetBuilder(8, 10)
	if err != nil {
		log.Fatal(err)
	}
	cvrt.Convert("771")
	fmt.Println(cvrt.GetResult()) // 505
}
