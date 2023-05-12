package main

import (
	"QzoneRecorder/pkg/impls/qzone"
)

func main() {
	mgr := qzone.NewQzoneManager()

	cookie_str, err := mgr.LoginViaQRCode(func(path string) {
		println("请扫描二维码: ", path)
	})
	if err != nil {
		panic(err)
	}

	println(cookie_str)
}
