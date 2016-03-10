package store

import (
	"io"
	"net/http"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

type Store interface {
	UploadFile(bucketName, localFile string) (string, error)
	UploadData(bucketName string, reader io.Reader, sizeOfFile int64) (string, error)
	Download(domain, key string) (io.ReadCloser, error)
	DownloadByUrl(url string) (io.ReadCloser, error)
	MakePrivateUrl(url string) string
}

type Qiniu struct {
	client  *kodo.Client
	context context.Context
}

func NewQiniu(accessKey, secretKey string, zone int) *Qiniu {
	kodo.SetMac(accessKey, secretKey)
	cxt := context.Background()
	cli := kodo.New(zone, nil)
	return &Qiniu{cli, cxt}
}

func (qn *Qiniu) UploadFile(bucketName, localFile string) (string, error) {
	bucket := qn.client.Bucket(bucketName)
	var ret kodo.PutRet
	err := bucket.PutFileWithoutKey(qn.context, &ret, localFile, nil)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func (qn *Qiniu) UploadData(bucketName string, reader io.Reader, sizeOfFile int64) (string, error) {
	bucket := qn.client.Bucket(bucketName)
	var ret kodo.PutRet
	err := bucket.PutWithoutKey(qn.context, &ret, reader, sizeOfFile, nil)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

func (qn *Qiniu) Download(domain, key string) (io.ReadCloser, error) {
	baseUrl := kodo.MakeBaseUrl(domain, key)
	return qn.DownloadByUrl(baseUrl)
}

func (qn *Qiniu) MakePrivateUrl(url string) string {
	return qn.client.MakePrivateUrl(url, nil)
}

func (qn *Qiniu) DownloadByUrl(url string) (io.ReadCloser, error) {
	privateUrl := qn.client.MakePrivateUrl(url, nil)
	resp, err := http.Get(privateUrl)
	return resp.Body, err
}
