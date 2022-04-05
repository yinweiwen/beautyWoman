package qiniu

import "testing"

func TestUpload(t *testing.T) {
	q := &QiniuUpload{
		accessKey: "z2qOOYBNuA1xBCS0RzJVn8jU41Nw2ZbXoUvQjMut",
		secretKey: "p-B0S_Xld8jLQLBRdBJvpK9-nG1dHTfyyjc8cZjL",
	}
	q.Upload("H:\\coding\\beautyWoman\\src\\download\\wallpaper_2022_4_2\\full_1k_wallhaven_1ked21.jpg")
}
