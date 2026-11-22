package main

import (
	"fmt"
	"lesson3/http_server"
)

func main() {
	fmt.Println("Запуска сервера")
	
	if err := http_server.StartHTTPServer(); err != nil {
		fmt.Println("сервер выдал ошибку")
	} else {
		fmt.Println("Success")
	}
}
