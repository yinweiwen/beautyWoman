package qiniu

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	_ "github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const QiniuBucket = "wallhaven1"

type QiniuUpload struct {
	accessKey string
	secretKey string

	watchdir string
	depth    int
	backdir  string
}

func NewQiniuUpload() *QiniuUpload {
	return &QiniuUpload{
		accessKey: "z2qOOYBNuA1xBCS0RzJVn8jU41Nw2ZbXoUvQjMut",
		secretKey: "p-B0S_Xld8jLQLBRdBJvpK9-nG1dHTfyyjc8cZjL",
		watchdir:  "download",
		depth:     2,
		backdir:   "backup",
	}
}

func (q *QiniuUpload) Do() {
	if len(q.backdir) > 0 {
		os.MkdirAll(q.backdir, 0777)
	}

	for true {
		q.ScanDirRoot()
		time.Sleep(30 * time.Second)
	}
}

func (q *QiniuUpload) ScanDirRoot() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	log.Println("[qiniu] begin scan root")
	q.ScanDir(q.watchdir, 1)
}

func (q *QiniuUpload) ScanDir(p string, depth int) {
	if depth > 2 {
		return
	}
	files, _ := ioutil.ReadDir(p)
	for _, fi := range files {
		pt := path.Join(p, fi.Name())
		if fi.IsDir() {
			q.ScanDir(pt, depth+1)
		} else {
			if strings.Contains(fi.Name(), ".jpg") {
				if err := q.Upload(pt); err == nil {
					if len(q.backdir) > 0 {
						nf := path.Join(q.backdir, fi.Name())
						os.Rename(pt, nf)
					}
				}
				// 防止上报过热
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (q *QiniuUpload) Upload(localFile string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("[qiniu] %v", r))
		}
	}()
	bucket := QiniuBucket
	localFile = strings.ReplaceAll(localFile, "\\", "/")
	key := path.Base(localFile)
	thumbnail := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:Thumbnail/%s", bucket, key)))

	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		PersistentOps: fmt.Sprintf("imageView2/1/w/200/h/200/q/75|saveas/%s", thumbnail), // 生成缩略图
	}
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		log.Println("[qiniu] upload failed,", err)
		return
	}
	log.Println("[qiniu] upload success: ", localFile)
	return
}
