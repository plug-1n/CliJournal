package main

import (
	"fmt"

	"github.com/plug-1n/cli-journal/internal/handlers"
)

func main() {
	fmt.Println("Start process")
	handlers.GetAllData()
}
