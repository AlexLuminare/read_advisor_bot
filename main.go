package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("TG_TOKEN")

	fmt.Printf("token: %s", token)
}
