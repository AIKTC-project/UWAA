package main

import (
	"fmt"
	"net/http"

	"blog_api/internal/pkg/get"
	"blog_api/internal/pkg/global"
)

func main() {
	router := http.NewServeMux()
	go router.HandleFunc("/api/sayhello", get.SayHello)

	fmt.Println("Running blog API on " + global.Port)
	http.ListenAndServe(global.Port, router)
}
