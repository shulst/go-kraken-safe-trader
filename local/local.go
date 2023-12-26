package main

import "github.com/shulst/go-kraken-safe-trader/router"

func main() {
	r := router.Router()
	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
