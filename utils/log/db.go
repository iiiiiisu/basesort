package log

import (
	"database/sql"
	"fmt"
)

type DatabaseConsole struct {
	db   *sql.DB
	stmt *sql.Stmt
	Logs []SortLog
}

func (c *DatabaseConsole) Init(driverName string, user string, pwd string, host string, port string, name string) error {
	dataSource := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8`, user, pwd, host, port, name)
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(2)
	c.stmt, err = db.Prepare(`INSERT INTO records (sort_name, start_at, length, seed, use_time, result) VALUES (?, ?, ?, ?, ?, ?)`)
	return err
}
func (l *DatabaseConsole) Add(log SortLog) {
	l.Logs = append(l.Logs, log)
	if len(l.Logs) > 10 {
		l.Output()
	}
}

func (l *DatabaseConsole) Output() {
	if len(l.Logs) == 0 {
		return
	}
	/*
		To Do
	*/
	l.Logs = []SortLog{}
}

func (c *DatabaseConsole) Close() {
	c.db.Close()
	c.stmt.Close()
}
