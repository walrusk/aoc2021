package util

import (
	"log"
	"os"
)

func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}
