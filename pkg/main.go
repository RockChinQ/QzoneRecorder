package main

import (
	"QzoneRecorder/pkg/utils"
)

func main() {
	err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
}
