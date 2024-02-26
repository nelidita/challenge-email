package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVars() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error open file .env:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) != 2 {
			log.Printf("Error in line format: %s\n", scanner.Text())
			continue
		}

		key, value := parts[0], parts[1]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error read file .env:", err)
	}
}
