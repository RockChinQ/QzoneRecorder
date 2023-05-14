package main

import (
	"QzoneRecorder/pkg/impls/qzone"
	"io/ioutil"
)

func main() {
	// 读取single_html.html的内容
	html_bytes, err := ioutil.ReadFile("single_html.html")

	if err != nil {
		panic(err)
	}

	_, err = qzone.ParseEmotionFromHTML(string(html_bytes))
	if err != nil {
		panic(err)
	}
}
