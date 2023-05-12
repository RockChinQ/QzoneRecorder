package main

import (
	"QzoneRecorder/pkg/impls/qzone"
	"fmt"
	"testing"
)

func TestQzoneLoginWithProvidedCookies(t *testing.T) {
	// 测试使用提供的cookies登录
	// mgr := qzone.NewQzoneManager()
}

func TestQzoneLoginViaQRCode(t *testing.T) {
	// 测试使用二维码登录
	mgr := qzone.NewQzoneManager()
	cookie_str, err := mgr.LoginViaQRCode(func(path string) {
		t.Logf("请扫描二维码: %s", path)
		fmt.Println("请扫描二维码: ", path)
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(cookie_str)

}
