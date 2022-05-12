package main

import (
	"git/router"
)
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
