package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"D:/Go react/internal/store/pgstore/migrations",
		"--config", "D:/Go react/internal/store/pgstore/migrations/tern.conf")
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Migration completed successfully.")
}
