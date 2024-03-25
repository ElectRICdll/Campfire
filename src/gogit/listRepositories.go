package gitserver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ListRepositories 列出所有仓库
func ListRepositories() ([]string, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败：%s", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT name FROM repositories")
	if err != nil {
		return nil, fmt.Errorf("查询数据库失败：%s", err)
	}
	defer rows.Close()

	var repositories []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("扫描查询结果失败：%s", err)
		}
		repositories = append(repositories, name)
	}

	return repositories, nil
}
