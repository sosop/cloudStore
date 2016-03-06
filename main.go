package main

import (
	"cloudStore/store"
	"fmt"
	"os"
)

func main() {
	var s store.Store
	accessKey := ""
	secretKey := ""
	s = store.NewQiniu(accessKey, secretKey, 0)
	// s.UploadFile("test", "/Users/mac/cutting/00.jpg")
	file, err := os.Open("/Users/mac/cutting/00.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	info, _ := file.Stat()
	fmt.Println(info.Size())
	err = s.UploadData("test", file, info.Size())
	if err != nil {
		panic(err)
	}
}
