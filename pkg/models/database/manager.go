package database

var DBMgr DatabaseAdapter

// 数据库接口适配器
type DatabaseAdapter interface {
	Connect() error
	Initialize() error
}
