package sink

import qiniu2 "craw/craw/sink/qiniu"

func Sinking() {
	qiniu := qiniu2.NewQiniuUpload()
	qiniu.Do()
}
