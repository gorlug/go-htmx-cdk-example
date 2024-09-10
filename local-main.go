package main

import (
	"cdk-example/app/api"
	"net/http"
)

func main() {
	api.CreateApi("app")
	http.ListenAndServe(":8081", nil)
}
