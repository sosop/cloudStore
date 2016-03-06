package store

import (
	"fmt"
	"io"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

type Store interface {
	UploadFile(bucketName, localFile string) error
	UploadData(bucketName string, reader io.Reader, sizeOfFile int64) error
	Download()
}

type Qiniu struct {
	client  *kodo.Client
	context context.Context
}

func NewQiniu(accessKey, secretKey string, zone int) Qiniu {
	kodo.SetMac(accessKey, secretKey)
	cxt := context.Background()
	cli := kodo.New(zone, nil)
	return Qiniu{cli, cxt}
}

func (qn Qiniu) UploadFile(bucketName, localFile string) error {
	bucket := qn.client.Bucket(bucketName)
	var ret kodo.PutRet
	err := bucket.PutFileWithoutKey(qn.context, &ret, localFile, nil)
	fmt.Println(ret)
	return err
}

func (qn Qiniu) UploadData(bucketName string, reader io.Reader, sizeOfFile int64) error {
	bucket := qn.client.Bucket(bucketName)
	var ret kodo.PutRet
	err := bucket.PutWithoutKey(qn.context, &ret, reader, sizeOfFile, nil)
	fmt.Println(ret)
	return err
}

func (qn Qiniu) Download() {

}
