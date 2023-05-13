package database

import (
	"QzoneRecorder/pkg/models/database"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

// mysql适配器实现manager.DatabaseAdapter接口
type MySQLAdapter struct {
	database.DatabaseAdapter
	db *sql.DB
}

func NewMySQLAdapter() *MySQLAdapter {
	return &MySQLAdapter{}
}

func (adapter *MySQLAdapter) Connect() error {
	// fmt.Println(viper.GetString("database.user") + ":" + viper.GetString("database.password") + "@tcp(" + viper.GetString("database.host") + ":" + viper.GetString("database.port") + ")/" + viper.GetString("database.database") + "?charset=utf8mb4")
	conn, err := sql.Open(
		"mysql",
		viper.GetString("database.user")+":"+viper.GetString("database.password")+"@tcp("+viper.GetString("database.host")+":"+viper.GetString("database.port")+")/"+viper.GetString("database.database")+"?charset=utf8mb4",
	)
	if err != nil {
		return err
	}
	conn.SetMaxOpenConns(10)
	adapter.db = conn

	return nil
}

var sqls = []string{
	`CREATE TABLE IF NOT EXISTS Person (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		qq BIGINT UNSIGNED NOT NULL,
		nick VARCHAR(255) NOT NULL,
		PRIMARY KEY (id, qq)
	)`,
	`CREATE TABLE IF NOT EXISTS Emotion (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		eid VARCHAR(255) NOT NULL,
		qq BIGINT UNSIGNED NOT NULL,
		person_id BIGINT UNSIGNED NOT NULL,
		text TEXT NOT NULL,
		time_stamp BIGINT UNSIGNED NOT NULL,
		FOREIGN KEY (person_id) REFERENCES Person(id),
		PRIMARY KEY (id, eid)
	)`,
	`CREATE TABLE IF NOT EXISTS Image (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		emotion_id BIGINT UNSIGNED NOT NULL,
		url VARCHAR(1024) NOT NULL,
		FOREIGN KEY (emotion_id) REFERENCES Emotion(id),
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS Comment (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		emotion_id BIGINT UNSIGNED NOT NULL,
		person_qq BIGINT UNSIGNED NOT NULL,
		text TEXT NOT NULL,
		time_stamp BIGINT UNSIGNED NOT NULL,
		parent_id BIGINT UNSIGNED NOT NULL,
		FOREIGN KEY (emotion_id) REFERENCES Emotion(id),
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS Traffic (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
		emotion_id BIGINT UNSIGNED NOT NULL,
		likes int UNSIGNED NOT NULL,
		visits int UNSIGNED NOT NULL,
		forwards int UNSIGNED NOT NULL,
		update_timestamp BIGINT UNSIGNED NOT NULL,
		FOREIGN KEY (emotion_id) REFERENCES Emotion(id),
		PRIMARY KEY (id)
	)`,
}

func (adapter *MySQLAdapter) Initialize() error {
	tx, err := adapter.db.Begin()
	if err != nil {
		return err
	}

	// 初始化数据库
	for _, sql := range sqls {
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
