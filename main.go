package main

import (
	"cloudStore/store"
	"os"
)

func main() {
	var s store.Store
	accessKey := "O3TdDpbhK1iLGT7lEIw8SD7k41D1I2Z_YItPYizz"
	secretKey := "sjI9C6b7EqO9NEJb9g4ow0z1dozkhqL0r5-K2kUx"
	domain := "7xrkxc.com1.z0.glb.clouddn.com"
	s = store.NewQiniu(accessKey, secretKey, 0)
	// s.UploadFile("test", "/Users/mac/cutting/00.jpg")
	file, err := os.Open("/Users/mac/Downloads/test.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	info, _ := file.Stat()
	data, err := s.UploadData("test", file, info.Size())
	if err != nil {
		panic(err)
	}
	rc, err := s.Download(domain, data)
	if err != nil {
		panic(err)
	}
	defer rc.Close()
}
