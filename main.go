package main

import (
	"fmt"
	router "gonix/src"
	"net/http"
)

func main() {

	http.HandleFunc("/", router.WebRouter)

	fmt.Println("started webServer at port", router.LISTEN_TO)

	http.ListenAndServe(router.LISTEN_TO, nil)

}
