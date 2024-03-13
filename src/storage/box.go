package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Box struct {
	ProjID uint
	oss.Bucket
}

func NewBox() {
	//client, err := oss.New(util.CONFIG.OSSEndPoint, util.CONFIG.OSSAccessKeyID, util.CONFIG.OSSSecretKey)
	//if err != nil {
	//	log.Error(err.Error())
	//}
}
