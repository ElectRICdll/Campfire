package util

import "time"

var CONFIG = struct {
	AuthDuration time.Time `yaml:"auth_duration"`
	SecretKey    []byte    `yaml:"secret_key"`
}{
	AuthDuration: time.Now().Add(time.Hour * 24 * 7),
	SecretKey:    []byte("Clover"),
}
