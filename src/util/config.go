package util

import "time"

var CONFIG = struct {
	Port    string `yaml:"port"`
	SQLConn string `yaml:"sql_conn"`

	OSSEndPoint    string `yaml:"oss_end_point"`
	OSSAccessKeyID string `yaml:"oss_access_key_id"`
	OSSSecretKey   string `yaml:"oss_secret_key"`

	NativeStorageRootPath string `yaml:"native_storage_root_path"`
	AvatarCacheRootPath   string `yaml:"avatar_cache_root_path"`

	AuthDuration time.Time `yaml:"auth_duration"`
	SecretKey    []byte    `yaml:"secret_key"`

	MessageRecordCount uint

	InvitationKeepDuration time.Duration
}{
	Port:               "9375",
	SQLConn:            "root:420204@tcp(120.24.78.233:3443)/Campfire?charset=utf8mb4&parseTime=True&loc=Local",
	AuthDuration:       time.Now().Add(time.Hour * 24 * 7),
	SecretKey:          []byte("Clover"),
	MessageRecordCount: 50,

	NativeStorageRootPath: "../repo",
	AvatarCacheRootPath:   "../assets/avatar",
}
