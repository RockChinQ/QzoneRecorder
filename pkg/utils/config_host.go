package utils

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/viper"
)

func CopyFile(from string, to string) error {
	// 复制文件
	// from: 源文件路径
	// to: 目标文件路径
	// 返回值: 错误信息
	// 若目标文件已存在则不复制
	if FileExists(to) {
		return nil
	}
	// 打开源文件
	from_file, err := os.Open(from)
	if err != nil {
		return err
	}
	// 创建目标文件
	to_file, err := os.Create(to)
	if err != nil {
		return err
	}
	// 复制文件
	_, err = io.Copy(to_file, from_file)
	if err != nil {
		return err
	}
	return nil
}

func FileExists(path string) bool {
	// 检查文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}

func LoadConfig() error {
	// 检查是否有config.yaml
	// 若没有则复制config-template.yaml创建，如何抛出异常
	if !FileExists("config.yaml") {
		err := CopyFile("config-template.yaml", "config.yaml")
		if err != nil {
			return err
		}
		return errors.New("无配置文件，请修改config.yaml中的配置")
	}
	// 读取config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
