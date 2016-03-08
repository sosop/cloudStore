package store

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStore struct {
	fileName string
}

func (f *FileStore) UploadFile(storePath, localFile string) (string, error) {
	srcFile, err := os.Open(localFile)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	srcName := srcFile.Name()
	hashName := md5.Sum([]byte(srcName))
	ext := filepath.Ext(srcName)
	targetName := fmt.Sprintf("%x", hashName) + ext

	targetFullPath := filepath.Join(storePath, targetName)
	targetFile, err := os.Create(targetFullPath)
	if err != nil {
		return "", nil
	}
	defer targetFile.Close()

	return targetFullPath, nil
}

func (f *FileStore) UploadData(storePath string, reader io.Reader, sizeOfFile int64) (string, error) {
	return "", nil
}

func (f FileStore) Download(path, key string) (io.ReadCloser, error) {
	return nil, nil
}
