package main

import (
	"fmt"
	"os"

	"github.com/misafidiniaina/soduku-cli/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
