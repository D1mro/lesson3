package main

import (
	"fmt"
	feature1 "lesson3/feature"
	"lesson3/feature2"
	"lesson3/feature_postres/simple_connection"
)

func main() {
	fmt.Println("hello")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}