package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading environment")
	}

	value := os.Getenv("DB_HOST")
	fmt.Println(value)
}
