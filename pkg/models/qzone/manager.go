package qzone

import (
	entities "QzoneRecorder/pkg/models/qzone/entities"
)

var QzMgr QzoneAdapter

// Qzone接口适配器
type QzoneAdapter interface {
	LoginViaQRCode(qr_got_callback func(path string)) (string, error)
	LoginViaCookies(cookies string) error
	GetVisitorAmount() (int, int, error)
	FetchFeedsList() ([]entities.Emotion, error)
}
