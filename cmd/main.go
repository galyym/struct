package main

import (
	"fmt"
	"github.com/galyym/struct/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatalln("File not found")
	}

	fmt.Println("Hello world!", file)

	fileById := st.GetById(file.ID)
	fmt.Println(fileById, "file exists")

	fmt.Println(st.files)
}
