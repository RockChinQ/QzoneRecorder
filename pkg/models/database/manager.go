package database

var db_mgr *DatabaseAdapter

type DatabaseAdapter interface {
	Connect() error
	Initialize() error
}
