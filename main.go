package main

import (
	"fmt"
	"initial/app"
)

func main() {
	application := app.Generate()
	fmt.Println(application)
}