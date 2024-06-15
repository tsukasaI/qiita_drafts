package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ls("/")

	cat("/sample.txt")
}

func ls(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dir {
		fmt.Println(d.Name())
	}
}

func cat(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
