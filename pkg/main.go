package main

import (
	"QzoneRecorder/pkg/impls/database"
	"QzoneRecorder/pkg/impls/qzone"
	_ "QzoneRecorder/pkg/impls/qzone"
	qzmodel "QzoneRecorder/pkg/models/qzone"
	"QzoneRecorder/pkg/utils"

	_ "QzoneRecorder/pkg/impls/database"
	dbmodel "QzoneRecorder/pkg/models/database"

	"fmt"

	"github.com/spf13/viper"
)

func main() {
	err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
	// 登录QQ空间
	qzmodel.QzMgr = qzone.NewQzoneManager()

	logged_in := false
	if viper.GetString("qzone.cookies") != "xxx" {
		err = qzmodel.QzMgr.LoginViaCookies(viper.GetString("qzone.cookies"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("登录失败，尝试使用二维码登录")
		} else {
			logged_in = true
			fmt.Println("使用cookies登录成功")
		}
	}
	if !logged_in {
		// 二维码登录
		cookies, err := qzmodel.QzMgr.LoginViaQRCode(func(path string) {
			fmt.Println("请扫描二维码", path)
		})
		if err != nil {
			fmt.Println("二维码登录失败")
			panic(err)
		}

		// 将cookies写入config.yaml
		viper.Set("qzone.cookies", cookies)
		err = viper.WriteConfig()
		if err != nil {
			panic(err)
		}

		fmt.Println("二维码登录成功, cookies已保存到config.yaml")
	}

	_, err = qzmodel.QzMgr.FetchFeedsList(1)
	if err != nil {
		panic(err)
	}

	// 数据库
	dbmodel.DBMgr = database.NewMySQLAdapter()

	err = dbmodel.DBMgr.Connect()
	if err != nil {
		panic(err)
	}

	// // 初始化数据库
	// err = dbmodel.DBMgr.Initialize()
	// if err != nil {
	// 	panic(err)
	// }

}
