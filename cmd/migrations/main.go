package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd := exec.Command("tern", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully!")
}
