package database

var db_mgr *DatabaseAdapter

// 数据库接口适配器
type DatabaseAdapter interface {
	Connect() error
	Initialize() error
}
