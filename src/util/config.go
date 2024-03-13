package util

import "time"

var CONFIG = struct {
	SQLConn string `yaml:"sql_conn"`

	OSSEndPoint    string `yaml:"oss_end_point"`
	OSSAccessKeyID string `yaml:"oss_access_key_id"`
	OSSSecretKey   string `yaml:"oss_secret_key"`

	AuthDuration time.Time `yaml:"auth_duration"`
	SecretKey    []byte    `yaml:"secret_key"`

	MessageRecordCount uint
}{
	SQLConn:            "root:420204@tcp(120.24.78.233:3443)/Campfire?charset=utf8mb4&parseTime=True&loc=Local",
	AuthDuration:       time.Now().Add(time.Hour * 24 * 7),
	SecretKey:          []byte("Clover"),
	MessageRecordCount: 50,
}
