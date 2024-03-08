package util

import "time"

var CONFIG = struct {
	SQLConn string `yaml:"sql_conn"`

	AuthDuration time.Time `yaml:"auth_duration"`
	SecretKey    []byte    `yaml:"secret_key"`
}{
	SQLConn:      "root:420204@tcp(127.24.78.233:3443)/Campsite.db?charset=utf8mb4&parseTime=True&loc=Local",
	AuthDuration: time.Now().Add(time.Hour * 24 * 7),
	SecretKey:    []byte("Clover"),
}
