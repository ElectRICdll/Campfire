package storage

import "fmt"

type Box struct {
	Name     string
	ProjID   uint
	RootPath string
}

func NewBox() {
	//client, err := oss.New(util.CONFIG.OSSEndPoint, util.CONFIG.OSSAccessKeyID, util.CONFIG.OSSSecretKey)
	//if err != nil {
	//	log.Error(err.Error())
	//}
}

func (b Box) BoxName() string {
	return fmt.Sprintf("%s-%x", b.Name, b.ProjID)
}

func (b Box) Push() {

}

func (b Box) Pull() {

}

func (b Box) Clone() {

}

func (b Box) Merge() {

}
