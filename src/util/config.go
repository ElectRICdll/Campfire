package util

import "time"

var CONFIG = struct {
	SQLConn string `yaml:"sql_conn"`

	AuthDuration time.Time `yaml:"auth_duration"`
	SecretKey    []byte    `yaml:"secret_key"`

	MessageRecordCount uint
}{
	SQLConn:            "root:420204@tcp(120.24.78.233:3443)/Campfire?charset=utf8mb4&parseTime=True&loc=Local",
	AuthDuration:       time.Now().Add(time.Hour * 24 * 7),
	SecretKey:          []byte("Clover"),
	MessageRecordCount: 50,
}
