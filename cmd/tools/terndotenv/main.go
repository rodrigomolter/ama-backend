package main

import (
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("NODE_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
